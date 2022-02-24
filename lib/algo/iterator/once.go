package iterator

import (
	. "github.com/Wei-N-Ning/gotypes/lib/algo/option"
)

func onceImpl[T any](x T) <-chan Option[T] {
	ch := make(chan Option[T])
	go func() {
		defer close(ch)
		ch <- Some(x)
		ch <- None[T]()
	}()
	return ch
}

func Once[T any](x T) Iterator[T] {
	return Iterator[T]{ch: onceImpl(x)}
}

func onceWithImpl[T any](f func() T) <-chan Option[T] {
	ch := make(chan Option[T])
	go func() {
		defer close(ch)
		ch <- Some(f())
		ch <- None[T]()
	}()
	return ch
}

func OnceWith[T any](f func() T) Iterator[T] {
	return Iterator[T]{ch: onceWithImpl(f)}
}
