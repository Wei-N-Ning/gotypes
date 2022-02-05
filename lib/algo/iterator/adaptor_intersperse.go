package iterator

import (
	. "go-types-nw/lib/algo/option"
)

func intersperseImpl[T any](iter Iterator[T], sep T) <-chan Option[T] {
	ch := make(chan Option[T])
	go func() {
		defer func() {
			ch <- None[T]()
			close(ch)
		}()
		head := iter.Next()
		if head.IsSome() {
			ch <- head
			iter.ForEach(func(x T) {
				ch <- Some(sep)
				ch <- Some(x)
			})
		}
	}()
	return ch
}

// Intersperse creates a new Iterator of the same type.
// It yields each element from the original iterator, inserting <sep> between two
// adjacent elements, e.g.
// given 1, 2, 3, 4 ... and sep = 100
// it produces 1, 100, 2, 100, 3, 100, 4 ...
func Intersperse[T any](iter Iterator[T], sep T) Iterator[T] {
	return Iterator[T]{ch: intersperseImpl(iter, sep), inner: iter}
}
