package iterator

import (
	. "github.com/Wei-N-Ning/gotypes/pkg/option"
)

func fromSliceImpl[T any](xs []T) <-chan Option[T] {
	ch := make(chan Option[T])
	go func() {
		defer close(ch)
		for _, x := range xs {
			ch <- Some(x)
		}
		ch <- None[T]()
	}()
	return ch
}

func FromSlice[T any](xs []T) Iterator[T] {
	return Iterator[T]{ch: fromSliceImpl(xs)}
}
