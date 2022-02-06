package iterator

import (
	. "go-types-nw/lib/algo/option"
)

func Find[T any](iter Iterator[T], f func(x T) bool) Option[T] {
	for {
		elem := iter.Next()
		if elem.IsSome() {
			if f(elem.Unwrap()) {
				return elem
			}
		} else {
			break
		}
	}
	return None[T]()
}
