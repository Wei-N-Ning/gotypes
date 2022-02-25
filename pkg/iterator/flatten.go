package iterator

import (
	. "github.com/Wei-N-Ning/gotypes/pkg/option"
)

// given an iterator of iterators in the form of Iterator[Iterator[T]],
// this function flattens their "nested structure" and produces a new iterator that
// contains the concatenated elements from all the iterators
func flattenImpl[T any](iterators Iterator[Iterator[T]]) <-chan Option[T] {
	ch := make(chan Option[T])
	go func() {
		defer func() {
			ch <- None[T]()
			defer close(ch)
		}()

		for {
			elem := iterators.Next()
			if elem.IsSome() {
				var iter Iterator[T] = elem.Unwrap()
				iter.ForEach(func(x T) {
					ch <- Some(x)
				})
			} else {
				return
			}
		}
	}()
	return ch
}

func Flatten[T any](iterators Iterator[Iterator[T]]) Iterator[T] {
	return Iterator[T]{ch: flattenImpl(iterators), inner: iterators}
}
