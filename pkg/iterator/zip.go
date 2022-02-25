package iterator

import (
	. "github.com/Wei-N-Ning/gotypes/pkg/option"
)

func zipImpl[T, P any](lhs Iterator[T], rhs Iterator[P]) <-chan Option[Pair[T, P]] {
	ch := make(chan Option[Pair[T, P]])
	go func() {
		defer func() {
			ch <- None[Pair[T, P]]()
			close(ch)
		}()
		for {
			lElem := lhs.Next()
			rElem := rhs.Next()
			if lElem.IsSome() && rElem.IsSome() {
				ch <- Some(NewPair(lElem.Unwrap(), rElem.Unwrap()))
			} else {
				return
			}
		}
	}()
	return ch
}

// Zip takes two iterators (with element type T and P), returning a new iterator that yields
// Pair[T, P].
// The number of pairs is determined by the shortest iterator, i.e.
// Zip(lhs, rhs).Count() == min(lhs.Count(), rhs.Count())
func Zip[T, P any](lhs Iterator[T], rhs Iterator[P]) Iterator[Pair[T, P]] {
	return Iterator[Pair[T, P]]{ch: zipImpl(lhs, rhs), inner: struct {
		lhs Iterator[T]
		rhs Iterator[P]
	}{lhs: lhs, rhs: rhs}}
}
