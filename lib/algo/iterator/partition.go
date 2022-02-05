package iterator

import (
	. "go-types-nw/lib/algo/option"
)

func partitionedImpl[T any](iter Iterator[T], f func(x T) bool) (<-chan Option[T], <-chan Option[T]) {
	trueCh := make(chan Option[T], 1024)
	falseCh := make(chan Option[T], 1024)
	go func() {
		defer func() {
			trueCh <- None[T]()
			falseCh <- None[T]()
			close(trueCh)
			close(falseCh)
		}()
		iter.ForEach(func(x T) {
			if f(x) {
				trueCh <- Some(x)
			} else {
				falseCh <- Some(x)
			}
		})
	}()
	return trueCh, falseCh
}

// Partitioned has a flaw:
// if the L channel is blocked on write, the R channel cannot make progress either
// (can think of it as a dead lock);
// they must be buffered channels
func Partitioned[T any](iter Iterator[T], f func(x T) bool) (Iterator[T], Iterator[T]) {
	l, r := partitionedImpl(iter, f)
	return Iterator[T]{ch: l, inner: iter}, Iterator[T]{ch: r, inner: iter}
}
