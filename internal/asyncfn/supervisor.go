package asyncfn

import "context"

// Func represents an effectful function
type Func func(ctx context.Context) error

// Handler represents an effectful signal handler that can tell the parent context to cancel
type Handler func(ctx context.Context, cancelFunc context.CancelFunc) error

var (
	// NOOP is an effectless function that does nothing
	NOOP = func(context.Context) error { return nil }
)

// Supervisor is a structure that encapsulates the signal-receiving mechanism, the signal-handling
// mechanism and a function subject.
//
// In a nutshell, the supervisor can stop the function subject based on the signal it receives
// and restart it if it sees fit.
//
type Supervisor struct {
	signals  []<-chan Signal
	handlers map[Signal]Handler // SIGTERM, SIGINT
	subject  Func               // would be the liberator PollReplication function
	preRun   Func               // (optionally) initialize
	afterRun Func               // (optionally) teardown/cleanup
}

// NewSupervisor is the only factory for the callers to construct an instance of Supervisor
// Callers are supposed to call NewSupervisor(), then populating each field using the WithXXX builder function.
func NewSupervisor(subject Func) *Supervisor {
	return &Supervisor{subject: subject, preRun: NOOP, afterRun: NOOP, handlers: map[Signal]Handler{}}
}

func (sup *Supervisor) WithSignals(signals ...<-chan Signal) *Supervisor {
	sup.signals = signals
	return sup
}

func (sup *Supervisor) WithHandler(signal Signal, hd Handler) *Supervisor {
	sup.handlers[signal] = hd
	return sup
}

func (sup *Supervisor) WithPreRun(fn Func) *Supervisor {
	sup.preRun = fn
	return sup
}

func (sup *Supervisor) WithAfterRun(fn Func) *Supervisor {
	sup.afterRun = fn
	return sup
}

func (sup *Supervisor) Supervise(parentCtx context.Context) error {
	ctx, cancel := context.WithCancel(parentCtx)

	// runs the function subject under supervision

	subjectDoneCh := make(chan error, 1)
	go func(parentCtx context.Context, doneCh chan<- error) {
		defer close(subjectDoneCh)
		if err := sup.preRun(parentCtx); err != nil {
			doneCh <- err
			return
		}
		if err := sup.subject(parentCtx); err != nil {
			doneCh <- err
			return
		}
		if err := sup.afterRun(parentCtx); err != nil {
			doneCh <- err
			return
		}
	}(ctx, subjectDoneCh)

	// runs the top-level, signal loop

	signalLoopDoneCh := make(chan error, 1)
	go func(parentCtx context.Context, doneCh chan<- error) {
		defer close(signalLoopDoneCh)
		for {
			select {
			case <-parentCtx.Done():
				return
			case err := <-subjectDoneCh:
				signalLoopDoneCh <- err
				return
			default:
				sig, yielded := trySelectOne(sup.signals...)
				if !yielded {
					continue
				}
				if handler, found := sup.handlers[sig]; found {
					err := handler(parentCtx, cancel)
					if err != nil {
						signalLoopDoneCh <- err
						return
					}
				}
			}
		}
	}(ctx, signalLoopDoneCh)

	// drain the top-level, signal channel and bubble up the error to the caller

	for err := range signalLoopDoneCh {
		if err != nil {
			return err
		}
	}

	return nil
}

//
//func AsyncCall() <-chan interface{} {
//	return nil
//}
//
//func Do(parent context.Context) {
//	for {
//		fut := AsyncCall()
//		ctx, cancel := context.WithCancel(parent)
//		for {
//			select {
//			case <-ctx.Done():
//			case <-fut:
//			default:
//			}
//		}
//	}
//}
