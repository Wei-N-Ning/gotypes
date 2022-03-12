package iterator

import (
	"github.com/Wei-N-Ning/gotypes/pkg/option"
	"sync"
)

// see:
// https://stackoverflow.com/questions/19992334/how-to-listen-to-n-channels-dynamic-select-statement

// best answer:
// https://cyolo.io/blog/how-we-enabled-dynamic-channel-selection-at-scale-in-go/

func Select[T any](buffer int64, iters ...Iterator[T]) Iterator[T] {
	r, w := TailAppender[T](buffer)
	go func() {
		wg := sync.WaitGroup{}
		defer func() {
			w <- option.None[T]()
			close(w)
		}()
		for _, iter := range iters {
			go func(iterator Iterator[T]) {
				defer wg.Done()
				iterator.ForEach(func(x T) { w <- option.Some[T](x) })
			}(iter)
			wg.Add(1)
		}
		wg.Wait()
	}()
	return r
}

// Merge
