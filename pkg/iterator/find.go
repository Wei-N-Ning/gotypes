package iterator

import (
	. "github.com/Wei-N-Ning/gotypes/pkg/option"
)

// Find returns the element index and the element if it is found;
// otherwise returns (-1, None);
// Caller should check if the element "IsNone" before consuming the index.
func Find[T any](iter Iterator[T], f func(x T) bool) (int, Option[T]) {
	idx := 0
	for {
		elem := iter.Next()
		if elem.IsSome() {
			if f(elem.Unwrap()) {
				return idx, elem
			}
		} else {
			break
		}
		idx += 1
	}
	return -1, None[T]()
}
