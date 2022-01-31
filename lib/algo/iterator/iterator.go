package iterator

import . "go-types-nw/lib/algo/option"

type Iterator[T any] struct {
	ch <-chan Option[T]
}

func (iter *Iterator[T]) Next() Option[T] {
	return <-iter.ch
}

func successorChan[T any](init Option[T], f func(T) Option[T]) <-chan Option[T] {
	ch := make(chan Option[T])
	go func() {
		defer func() {
			close(ch)
		}()
		for {
			if init.IsSome() {
				ch <- init // block
				init = f(init.Unwrap())
			} else {
				return
			}
		}
	}()
	return ch
}

func Successor[T any](init Option[T], f func(T) Option[T]) Iterator[T] {
	return Iterator[T]{ch: successorChan(init, f)}
}

func takeChan[T any](iter *Iterator[T], num int) <-chan Option[T] {
	ch := make(chan Option[T])
	go func() {
		defer func() {
			close(ch)
		}()
		for i := 0; i < num; i++ {
			ch <- iter.Next() // block
		}
	}()
	return ch
}

func (iter Iterator[T]) Take(num int) Iterator[T] {
	return Iterator[T]{ch: takeChan[T](&iter, num)}
}

func (iter *Iterator[T]) ForEach(f func(T)) {
	for {
		elem := iter.Next()
		if elem.IsSome() {
			f(elem.Unwrap())
		} else {
			return
		}
	}
}
