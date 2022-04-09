package taskqueue

import "context"

type taskOutput[T any] struct {
	value T
	err   error
}

type TaskQueue[T, P any] struct {
	size   int
	fi     chan T
	fo     chan chan taskOutput[P]
	worker func(context.Context, T) (P, error)
}

func (tq *TaskQueue[T, P]) Enqueue(input T) {
	tq.fi <- input
}

func (tq *TaskQueue[T, P]) Dequeue() (P, error) {
	o := <-<-tq.fo
	return o.value, o.err
}

func NewTaskQueue[T, P any](ctx context.Context, size int, f func(context.Context, T) (P, error)) TaskQueue[T, P] {
	tq := TaskQueue[T, P]{
		size:   size,
		fi:     make(chan T, size),
		fo:     make(chan chan taskOutput[P], size),
		worker: f,
	}
	go func() {
		defer close(tq.fi)
		defer close(tq.fo)

		for {
			select {
			case <-ctx.Done():
				return
			case x := <-tq.fi:
				outCh := make(chan taskOutput[P])
				go func(ctx context.Context, input T) {
					defer close(outCh)
					output, err := f(ctx, input)
					outCh <- taskOutput[P]{value: output, err: err}
				}(ctx, x)
				tq.fo <- outCh
			}
		}
	}()
	return tq
}
