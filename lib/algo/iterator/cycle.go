package iterator

import (
	. "github.com/Wei-N-Ning/gotypes/lib/algo/option"
)

func cycleImpl[T any](iter Iterator[T]) <-chan Option[T] {
	ch := make(chan Option[T])
	buffer := make(chan T, 1024)
	go func() {
		defer func() {
			ch <- None[T]()
			close(ch)
		}()
		numElems := 0
		iter.ForEach(func(x T) {
			ch <- Some(x)
			buffer <- x
			numElems += 1
		})
		if numElems == 0 {
			return
		}
		for {
			x := <-buffer
			ch <- Some(x)
			buffer <- x
		}
	}()
	return ch
}

// Cycle creates an infinite sequence using the given iterator;
// When it yields the last element from the original iterator, it repeats the sequence from the first element again.
func (iter Iterator[T]) Cycle() Iterator[T] {
	return Iterator[T]{ch: cycleImpl(iter), inner: iter}
}
