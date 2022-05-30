package asyncfn

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var (
	FooError = fmt.Errorf("foo")
)

func TestStatelessSubjects(t *testing.T) {
	t.Run("run and return, expect no error", func(t *testing.T) {
		sup := NewSupervisor(NOOP) // all the function subjects are NOOP - they do nothing
		err := sup.Supervise(context.Background())
		require.NoError(t, err)
	})
	t.Run("expect subject's error bubbled up", func(t *testing.T) {
		// given the Supervisor a function subject that returns an error
		sup := NewSupervisor(func(context.Context) error {
			// this would be the streaming function (PollReplication) or a forever-running server function;
			// return an error to simulate the real world erroring condition
			return FooError
		})
		err := sup.Supervise(context.Background())
		require.Error(t, err)
		require.Equal(t, FooError, err)
	})
	t.Run("expect catching timeout signal", func(t *testing.T) {
		ctx := context.Background()

		// given the Supervisor a function subject that runs forever
		sup := NewSupervisor(func(context.Context) error {
			// would be the function that runs forever
			time.Sleep(36000 * time.Second)
			return nil
		})

		// the system accepts the timeout signal; if there are other signal sources, pass them to this function,
		// e.g.
		// sup.WithSignals(timeoutSignal(...), unixSignals(...))
		sup.WithSignals(timeoutSignal(ctx, 100*time.Millisecond))

		// each handler function is called for one and only one type of signal;
		// the handler is given two values:
		// the system-wide context and its cancellation function;
		// the handler can craft child context(s) using the given system-wide context and/or issue cancellation
		// by calling the "fn" function, which is of type context.CancelFunc;
		// cancellation is propagated system-wide and will result in a shutdown or restart
		sup.WithHandler(TimeoutSignal, func(_ context.Context, fn context.CancelFunc) error {
			fn()
			return nil
		})

		// what happens is:
		// - Supervisor invokes the function subject - the one that sleeps for 36000 seconds - in a goroutine;
		// - At the same time, Supervisor watches for incoming signals
		// - Supervisor gets a timeout signal after roughly 100ms
		// - Calls the handler function associated with the timeout signal
		// - The handler issues a system-wide shutdown
		// - Supervisor shutdowns the function subject and then itself
		err := sup.Supervise(context.Background())
		require.NoError(t, err)
	})
	t.Run("expect catching operating system signal", func(t *testing.T) {
		ctx := context.Background()

		sup := NewSupervisor(func(context.Context) error {
			time.Sleep(36000 * time.Second)
			return nil
		})
		sup.WithSignals(stubSignals(TerminationSignal))
		sup.WithHandler(TerminationSignal, func(ctx context.Context, cancelFunc context.CancelFunc) error {
			cancelFunc()
			return nil
		})

		err := sup.Supervise(ctx)
		require.NoError(t, err)
	})
}

type RestartableSubject struct {
	State int
}

func NewSubject() *RestartableSubject {
	return &RestartableSubject{}
}

func (subj *RestartableSubject) Init(context.Context) error {
	subj.State = 1000
	return nil
}

func (subj *RestartableSubject) Run(context.Context) error {
	subj.State++
	return nil
}

func (subj *RestartableSubject) Cleanup(context.Context) error {
	return nil
}

func (subj *RestartableSubject) Restart(_ context.Context, fn context.CancelFunc) error {
	return nil
}

func TestStatefulSubjects(t *testing.T) {
	t.Run("restart subject", func(t *testing.T) {
		ctx := context.Background()

		subject := NewSubject()

		sup := NewSupervisor(subject.Run)
		sup.WithSignals(stubSignals(RestartSignal))
		sup.WithHandler(RestartSignal, subject.Restart)

		err := sup.Supervise(ctx)
		require.NoError(t, err)
	})
}
