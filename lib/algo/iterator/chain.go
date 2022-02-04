package iterator

import (
	. "go-types-nw/lib/algo/option"
)

func chainImpl[T any](iterators ...Iterator[T]) <-chan Option[T] {
	ch := make(chan Option[T])
	go func() {
		defer func() {
			ch <- None[T]()
			close(ch)
		}()
		for _, iter := range iterators {
			iter.ForEach(func(x T) { ch <- Some(x) })
		}
	}()
	return ch
}

func Chain[T any](iterators ...Iterator[T]) Iterator[T] {
	return Iterator[T]{ch: chainImpl(iterators...), inner: iterators}
}
