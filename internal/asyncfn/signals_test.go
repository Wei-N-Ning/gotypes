package asyncfn

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

func TestMixedSignalsExpectOperatingSystemSignal(t *testing.T) {
	ctx := context.Background()
	osSigCh := stubSignals(ReloadConfigSignal, TerminationSignal)
	timeoutSigCh := timeoutSignal(ctx, 2*time.Second)

	for {
		if x, yielded := trySelectOne(osSigCh, timeoutSigCh); yielded {
			require.Equal(t, ReloadConfigSignal, x)
			break
		}
	}
}

func TestMixedSignalsExpectTimeoutSignal(t *testing.T) {
	ctx := context.Background()
	osSigCh := unixSignals(ctx, unix.SIGUSR1, unix.SIGTERM)
	timeoutSigCh := timeoutSignal(ctx, 40*time.Millisecond)

	for {
		if x, yielded := trySelectOne(osSigCh, timeoutSigCh); yielded {
			require.Equal(t, TimeoutSignal, x)
			break
		}
	}
}

func TestMixedSignalsParentContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	osSigCh := unixSignals(ctx, unix.SIGUSR1, unix.SIGTERM)
	timeoutSigCh := timeoutSignal(ctx, 400000*time.Second)

	go func() {
		timer := time.NewTimer(40 * time.Millisecond)
		<-timer.C
		cancel()
	}()

	for {
		select {
		case _, isOpen := <-osSigCh:
			if !isOpen {
				osSigCh = nil
			}
		case _, isOpen := <-timeoutSigCh:
			if !isOpen {
				timeoutSigCh = nil
			}
		default:
		}
		if osSigCh == nil && timeoutSigCh == nil {
			break
		}
	}
}

func TestMixedSignalsParentContextCancellationSelectOne(t *testing.T) {
	t.Run("parent context cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		osSigCh := unixSignals(ctx, unix.SIGUSR1, unix.SIGTERM)
		timeoutSigCh := timeoutSignal(ctx, 400000*time.Second)

		go func() {
			timer := time.NewTimer(40 * time.Millisecond)
			<-timer.C
			cancel()
		}()

		_, yielded := selectOne(ctx, osSigCh, timeoutSigCh)
		require.False(t, yielded)
	})
	t.Run("all the channels are closed", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		osSigCh := unixSignals(ctx, unix.SIGUSR1, unix.SIGTERM)
		timeoutSigCh := timeoutSignal(ctx, 400000*time.Second)

		go func() {
			timer := time.NewTimer(40 * time.Millisecond)
			<-timer.C
			cancel()
		}()

		_, yielded := selectOne(context.Background(), osSigCh, timeoutSigCh)
		require.False(t, yielded)
	})
}
