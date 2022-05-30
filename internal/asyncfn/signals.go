package asyncfn

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// Signal abstracts away the platform-specific signal type(s) and provide a unified interface.
type Signal struct{ string }

var (
	InvalidSignal      = Signal{}
	ReloadConfigSignal = Signal{"RELOAD"}
	TerminationSignal  = Signal{"TERM"}
	TimeoutSignal      = Signal{"TIMEOUT"}
	RestartSignal      = Signal{"RESTART"}
	CustomSignal       = Signal{"CUSTOM"}
)

// NewDiscardedSignal returns a signal that is supposed to be discarded by the caller
func NewDiscardedSignal(s string) Signal {
	return Signal{"DISCARD_" + s}
}

// Discarded checks whether "self" is a discarded signal
func (s *Signal) Discarded() bool {
	return strings.HasPrefix(s.string, "DISCARD_")
}

// unixSignals responds to the operating system signals, such as
// kill -SIGTERM XXXX [XXXX - PID for your program]
//
// it returns a read-only channel that yields Signal (mapping the low-level platform-specific signal type to a high-level
// Signal type)
//
// example:
//
// to catch the operating system's SIGTERM and SIGHUP signals, do
//
// ch := unixSignals(ctx, syscall.SIGTERM, syscall.SIGHUP)
// for sig := range ch { // ... }
func unixSignals(ctx context.Context, sigs ...os.Signal) <-chan Signal {
	inCh := make(chan os.Signal, 32)
	outCh := make(chan Signal, 32)
	signal.Notify(inCh, sigs...)
	go func() {
		signal.Stop(inCh)
		defer close(inCh)
		defer close(outCh)
		for {
			select {

			case <-ctx.Done():
				return

			case sig, isChanOpen := <-inCh:
				if !isChanOpen {
					return
				}

				switch sig {
				case syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT:
					outCh <- TerminationSignal
				case syscall.SIGHUP, syscall.SIGUSR1:
					outCh <- ReloadConfigSignal
				case syscall.SIGUSR2:
					outCh <- CustomSignal
				default:
					// by default, we ignore any other signals but keep their details so we can log them out
					outCh <- NewDiscardedSignal(sig.String())
				}
			}
		}
	}()
	return outCh
}

// timeoutSignal returns an one-off read-only channel that yields the TimeoutSignal when the system hits the timeout
func timeoutSignal(ctx context.Context, timeout time.Duration) <-chan Signal {
	timer := time.NewTimer(timeout)
	outCh := make(chan Signal, 1)
	go func() {
		defer close(outCh)
		for {
			select {
			case <-ctx.Done():
				return
			case <-timer.C:
				outCh <- TimeoutSignal
				return
			}
		}
	}()
	return outCh
}
