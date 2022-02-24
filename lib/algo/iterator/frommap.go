package iterator

import (
	. "github.com/Wei-N-Ning/gotypes/lib/algo/option"
)

func fromMapImpl[K comparable, V any](m map[K]V) <-chan Option[Pair[K, V]] {
	ch := make(chan Option[Pair[K, V]])
	go func() {
		defer func() {
			ch <- None[Pair[K, V]]()
			close(ch)
		}()
		for key, value := range m {
			ch <- Some(NewPair(key, value))
		}
	}()
	return ch
}

func FromMap[K comparable, V any](m map[K]V) Iterator[Pair[K, V]] {
	return Iterator[Pair[K, V]]{ch: fromMapImpl(m)}
}
