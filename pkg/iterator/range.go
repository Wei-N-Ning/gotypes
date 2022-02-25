package iterator

import (
	. "github.com/Wei-N-Ning/gotypes/pkg/option"
)

func rangeImpl(from int, to int) <-chan Option[int] {
	ch := make(chan Option[int])
	go func() {
		defer close(ch)
		for i := from; i < to; i++ {
			ch <- Some[int](i)
		}
		ch <- None[int]()
	}()
	return ch
}

// Range is exclusive
// from=1, to=100 results in a sequence of (1, 2, ... 99), 99 elements
func Range(from int, to int) Iterator[int] {
	return Iterator[int]{ch: rangeImpl(from, to)}
}

func RangeInclusive(first int, last int) Iterator[int] {
	return Iterator[int]{ch: rangeImpl(first, last+1)}
}
