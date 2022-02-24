package iterator

import (
	. "github.com/Wei-N-Ning/gotypes/lib/algo/option"
)

// Concat creates a new iterator from an arbitrary number of iterators of type Iterator[T],
// concatenating their elements
func Concat[T any](iterators ...Iterator[T]) Iterator[T] {
	reader, writer := TailAppender[T](0)
	go func() {
		defer close(writer)
		for _, iter := range iterators {
			iter.ForEach(func(x T) {
				writer <- Some(x)
			})
		}
	}()
	return reader
}
