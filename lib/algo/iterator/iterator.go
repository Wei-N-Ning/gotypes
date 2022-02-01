package iterator

import (
	. "go-types-nw/lib/algo/option"
)

// Iterator is essentially a generic read-channel with behaviors.
// The write-end of the channel must signal the termination by sending through a None value.
type Iterator[T any] struct {
	ch    <-chan Option[T]
	inner interface{}
}

func (iter Iterator[T]) Next() Option[T] {
	return <-iter.ch
}
