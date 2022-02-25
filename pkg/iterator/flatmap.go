package iterator

import (
	. "github.com/Wei-N-Ning/gotypes/pkg/option"
)

func flatMapImpl[T, R any](iter Iterator[T], mapper func(x T) Iterator[R]) <-chan Option[R] {
	ch := make(chan Option[R])
	aggregator := make(chan Iterator[R])
	go func() {
		defer close(aggregator)
		iter.ForEach(func(x T) {
			aggregator <- mapper(x)
		})
	}()
	go func() {
		defer func() {
			ch <- None[R]()
			close(ch)
		}()
		for it := range aggregator {
			it.ForEach(func(x R) {
				ch <- Some(x)
			})
		}
	}()
	return ch
}

func FlatMap[T, R any](iter Iterator[T], mapper func(x T) Iterator[R]) Iterator[R] {
	return Iterator[R]{ch: flatMapImpl(iter, mapper), inner: iter}
}
