package iterator

import (
	. "go-types-nw/lib/algo/option"
)

func parMapImpl[T, R any](iter Iterator[T], f func(x T) R) <-chan Option[R] {
	ch := make(chan Option[R], 1024)
	outIter := Map(iter, func(x T) Iterator[R] {
		return OnceWith(func() R { return f(x) })
	})
	go func() {
		defer close(ch)
		outIter.ForEach(func(elem Iterator[R]) {
			ch <- elem.Next()
		})
		ch <- None[R]()
	}()
	return ch
}

// ParMap respects the original order, but this causes significant overhead.
// If order is not important, use ParMapUnord instead.
// See parmap_test.go for a rough comparison between these two versions.
func ParMap[T, R any](iter Iterator[T], f func(x T) R) Iterator[R] {
	return Iterator[R]{ch: parMapImpl(iter, f), inner: iter}
}

func parMapUnorderedImpl[T, R any](iter Iterator[T], f func(x T) R) <-chan Option[R] {
	ch := make(chan Option[R], 1024)
	aggregator := make(chan R)
	go func() {
		defer func() {
			ch <- None[R]()
			close(ch)
		}()
		num := 0
		iter.ForEach(func(x T) {
			go func() {
				aggregator <- f(x)
			}()
			num += 1
		})
		for i := 0; i < num; i++ {
			ch <- Some[R](<-aggregator)
		}
	}()
	return ch
}

// ParMapUnord disregard the order but can achieve better performance.
func ParMapUnord[T, R any](iter Iterator[T], f func(x T) R) Iterator[R] {
	return Iterator[R]{ch: parMapUnorderedImpl(iter, f), inner: iter}
}

func parMapReduceImpl[T, R any](iter Iterator[T], init R, mapper func(x T) R, reducer func(R, R) R) R {
	ch := make(chan Option[R], 1024)
	aggregator := make(chan R)
	go func() {
		defer func() {
			ch <- None[R]()
			close(ch)
		}()
		num := 0
		iter.ForEach(func(x T) {
			go func() {
				aggregator <- mapper(x)
			}()
			num += 1
		})
		for i := 0; i < num; i++ {
			ch <- Some[R](<-aggregator)
		}
	}()
	return ch
}
