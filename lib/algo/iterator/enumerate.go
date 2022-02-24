package iterator

import (
	. "github.com/Wei-N-Ning/gotypes/lib/algo/option"
)

func withIndexImpl[T any](iter Iterator[T]) <-chan Option[Pair[int64, T]] {
	ch := make(chan Option[Pair[int64, T]])
	go func() {
		defer func() {
			ch <- None[Pair[int64, T]]()
			close(ch)
		}()
		var index int64 = 0
		for {
			elem := iter.Next()
			if elem.IsSome() {
				ch <- Some(NewPair(index, elem.Unwrap()))
				index += 1
			} else {
				return
			}
		}
	}()
	return ch
}

func WithIndex[T any](iter Iterator[T]) Iterator[Pair[int64, T]] {
	return Iterator[Pair[int64, T]]{ch: withIndexImpl(iter), inner: iter}
}
