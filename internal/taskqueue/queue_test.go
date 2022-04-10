package taskqueue

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTaskQueueFirstInFirstOut(t *testing.T) {
	simpleTask := func(_ context.Context, x int) (int, error) { return x, nil }

	t.Run("ensure input-output order", func(t *testing.T) {
		tq := NewTaskQueue(context.Background(), 128, simpleTask)
		xs := []int{1, 2, 3, 4, 5, 6, 7}
		for _, x := range xs {
			tq.Enqueue(x)
		}
		var ys []int
		for range xs {
			out, err := tq.Dequeue()
			assert.NoError(t, err)
			ys = append(ys, out)
		}
		assert.Equal(t, xs, ys)
	})
}

func TestTaskQueueTimeout(t *testing.T) {
	var EOL = fmt.Errorf("END OF LINE")

	slowTask := func(ctx context.Context, x int) (int, error) {
		ticker := time.Tick(10 * time.Second)
		for {
			select {
			case <-ctx.Done():
				return 0, EOL
			case <-ticker:
				return x, nil
			}
		}
	}

	t.Run("timeout via context", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
		defer cancel()
		tq := NewTaskQueue(ctx, 128, slowTask)
		xs := []int{1, 2, 3, 4, 5, 6, 7}
		for _, x := range xs {
			tq.Enqueue(x)
		}
		_, err := tq.Dequeue()
		assert.Equal(t, context.DeadlineExceeded, ctx.Err())
		assert.Equal(t, EOL, err)
	})

	t.Run("parent cancellation", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		tq := NewTaskQueue(ctx, 128, slowTask)
		xs := []int{1, 2, 3, 4, 5, 6, 7}
		for _, x := range xs {
			tq.Enqueue(x)
		}
		ticker := time.Tick(100 * time.Millisecond)
		for range ticker {
			cancel()
			cancel() // calling cancel() twice appears harmless
			break
		}
		_, err := tq.Dequeue()
		assert.Equal(t, context.Canceled, ctx.Err())
		assert.Equal(t, EOL, err)
	})
}

func fibTask(ctx context.Context, n uint) (uint, error) {
	if n < 2 {
		return 1, nil
	} else {
		first, _ := fibTask(ctx, n-1)
		second, _ := fibTask(ctx, n-2)
		return first + second, nil
	}
}

func TestConcurrentTaskExecution(t *testing.T) {
	t.Run("expect tasks utilizing CPUs in parallel", func(t *testing.T) {
		tq := NewTaskQueue(context.Background(), 128, fibTask)
		xs := []uint{35, 36, 35, 36, 35, 36, 35, 36, 35, 36, 35, 36}
		for _, x := range xs {
			tq.Enqueue(x)
		}
		var ys []uint
		for range xs {
			out, err := tq.Dequeue()
			assert.NoError(t, err)
			ys = append(ys, out)
		}
		fmt.Println(ys)
	})
}

func TestBlockingTask(t *testing.T) {
	t.Run("one task can block Dequeue() but not the other task executions", func(t *testing.T) {
		var EOL = fmt.Errorf("END OF LINE")

		slowTask := func(ctx context.Context, x int) (int, error) {
			input := x
			for {
				select {
				case <-ctx.Done():
					return 0, EOL
				default:
					if x <= 0 {
						return input, nil
					} else {
						x--
					}
					time.Sleep(10 * time.Millisecond)
				}
			}
		}

		xs := []int{80, 20, 30, 40, 50, 60, 70, 71, 72, 73, 74}
		tq := NewTaskQueue(context.Background(), 128, slowTask)
		for _, x := range xs {
			tq.Enqueue(x)
		}
		tq.Seal()
		var ys []int
		for output := range tq.OutputChannel() {
			assert.NoError(t, output.Err)
			ys = append(ys, output.Value)
		}
		assert.Equal(t, xs, ys)
	})
}
