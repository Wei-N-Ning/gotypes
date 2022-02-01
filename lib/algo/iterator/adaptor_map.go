package iterator

import (
	. "go-types-nw/lib/algo/option"
)

func mapImpl[T any, R any](iter Iterator[T], f func(T) R) <-chan Option[R] {
	ch := make(chan Option[R])
	go func() {
		defer close(ch)
		for {
			elem := iter.Next()
			if !elem.IsSome() {
				ch <- None[R]()
				return
			} else {
				ch <- Some[R](f(elem.Unwrap()))
			}
		}
	}()
	return ch
}

// Map applies a function to each element and produces a new value;
// The results are yield in the new iterator.
func Map[T any, R any](iter Iterator[T], f func(T) R) Iterator[R] {
	return Iterator[R]{ch: mapImpl(iter, f), inner: iter}

}
