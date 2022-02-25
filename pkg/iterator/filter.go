package iterator

import (
	. "github.com/Wei-N-Ning/gotypes/pkg/option"
)

func filterImpl[T any](iter Iterator[T], f func(T) bool) <-chan Option[T] {
	ch := make(chan Option[T])
	go func() {
		defer close(ch)
		for {
			elem := iter.Next()
			if !elem.IsSome() {
				ch <- elem
				return
			}
			if f(elem.Unwrap()) {
				ch <- elem
			}
		}
	}()
	return ch
}

// Filter applies a function to each element and gets a boolean value.
// If the value is true, the element is included in the resulting new iterator, otherwise it is discarded.
func (iter Iterator[T]) Filter(f func(x T) bool) Iterator[T] {
	return Iterator[T]{ch: filterImpl(iter, f), inner: iter}
}
