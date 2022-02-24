package iterator

import (
	. "github.com/Wei-N-Ning/gotypes/lib/algo/option"
)

func successorImpl[T any](init Option[T], f func(T) Option[T]) <-chan Option[T] {
	ch := make(chan Option[T])
	go func() {
		defer close(ch)
		for {
			ch <- init
			if init.IsSome() {
				init = f(init.Unwrap())
			} else {
				return
			}
		}
	}()
	return ch
}

// Successor is copycat of Rust's successor function.
// It is an iterator-creator (or source).
// It repeatedly applies f to init (an Option[T]), yielding the result, till init becomes None.
// You can think of it as a series defined in the recursive form:
// Pn+1 = F( Pn )
func Successor[T any](init Option[T], f func(T) Option[T]) Iterator[T] {
	return Iterator[T]{ch: successorImpl(init, f)}
}
