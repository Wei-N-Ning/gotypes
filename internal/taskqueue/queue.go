package taskqueue

import "context"

type TaskOutput[T any] struct {
	Value T
	Err   error
}

type TaskQueue[T, P any] struct {
	size   int
	fi     chan T
	fo     chan chan TaskOutput[P]
	worker func(context.Context, T) (P, error)
}

func (tq *TaskQueue[T, P]) Enqueue(input T) {
	tq.fi <- input
}

func (tq *TaskQueue[T, P]) Seal() {
	close(tq.fi)
}

func (tq *TaskQueue[T, P]) Dequeue() (P, error) {
	o := <-<-tq.fo
	return o.Value, o.Err
}

func (tq *TaskQueue[T, P]) OutputChannel() <-chan TaskOutput[P] {
	outCh := make(chan TaskOutput[P], tq.size)
	go func() {
		defer close(outCh)
		for taskFuture := range tq.fo {
			outCh <- <-taskFuture
		}
	}()
	return outCh
}

func NewTaskQueue[T, P any](ctx context.Context, size int, f func(context.Context, T) (P, error)) TaskQueue[T, P] {
	tq := TaskQueue[T, P]{
		size:   size,
		fi:     make(chan T, size),
		fo:     make(chan chan TaskOutput[P], size),
		worker: f,
	}
	go func() {
		defer close(tq.fo)

		for {
			select {
			case <-ctx.Done():
				return
			case x, isChanOpen := <-tq.fi:
				if !isChanOpen {
					return
				}
				outCh := make(chan TaskOutput[P])
				go func(ctx context.Context, input T) {
					defer close(outCh)
					output, err := f(ctx, input)
					outCh <- TaskOutput[P]{Value: output, Err: err}
				}(ctx, x)
				tq.fo <- outCh
			}
		}
	}()
	return tq
}
