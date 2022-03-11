package iterator

import "github.com/Wei-N-Ning/gotypes/pkg/option"

// see:
// https://stackoverflow.com/questions/19992334/how-to-listen-to-n-channels-dynamic-select-statement

func Race[T any](buffer int64, iters ...Iterator[T]) Iterator[T] {
	r, w := TailAppender[T](buffer)
	go func() {
		defer func() {
			w <- option.None[T]()
			close(w)
		}()
	}()
	return r
}
