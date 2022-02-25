package iterator

import (
	. "github.com/Wei-N-Ning/gotypes/pkg/option"
)

func chunkSliceImpl[T any](iter Iterator[T], size int) <-chan Option[[]T] {
	ch := make(chan Option[[]T])
	go func() {
		defer func() {
			ch <- None[[]T]()
			close(ch)
		}()
		for {
			var xs []T
			for i := 0; i < size; i++ {
				elem := iter.Next()
				if elem.IsSome() {
					// keep filling the slice
					xs = append(xs, elem.Unwrap())
				} else {
					// yield the last slice then gracefully exit
					ch <- Some(xs)
					return
				}
			}
			// if it gets here, the slice is full, waiting to yield
			ch <- Some(xs)
			// the slice will be reset in the next round
		}
	}()
	return ch
}

// ChunkSlice
// given an iterator that yields:
// 1, 2, 3, 4, 5 ...... n
// given a chunk size of p
// It produces a new iterator that yields a slice of p elements at a time (or less) from the original sequence:
// (1, 2, 3 ... p), (p+1, .... p+p), ...
func ChunkSlice[T any](iter Iterator[T], size int) Iterator[[]T] {
	return Iterator[[]T]{ch: chunkSliceImpl(iter, size), inner: iter}
}
