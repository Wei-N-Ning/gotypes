package iterator

import (
	. "go-types-nw/lib/algo/option"
)

func takeImpl[T any](iter Iterator[T], num int) <-chan Option[T] {
	ch := make(chan Option[T])
	go func() {
		defer close(ch)
		for i := 0; i < num; i++ {
			elem := iter.Next()
			ch <- elem
			if !elem.IsSome() {
				return
			}
		}
	}()
	return ch
}

// Take is an iterator-filter that yields the first <num> elements;
// The iterator can terminate before it reaches <num>.
// It can be useful when dealing with an infinite series (Fibonacci series, the sequence of prime numbers, etc.)
func (iter Iterator[T]) Take(num int) Iterator[T] {
	return Iterator[T]{ch: takeImpl[T](iter, num), inner: iter}
}

// TakeWhile
