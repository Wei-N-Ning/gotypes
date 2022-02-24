package iterator

import (
	. "github.com/Wei-N-Ning/gotypes/lib/algo/option"
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

func (iter Iterator[T]) Tail() Iterator[T] {
	iter.Next()
	return iter
}

// AdvanceBy advances the iterator position by <num> which is equivalent to
// repeatedly calling Next() <num> times
func (iter Iterator[T]) AdvanceBy(num int) Iterator[T] {
	for i := 0; i < num; i++ {
		if !iter.Next().IsSome() {
			break
		}
	}
	return iter
}

func (iter Iterator[T]) Last() Option[T] {
	opt := None[T]()
	iter.ForEach(func(x T) {
		opt.Replace(x)
	})
	return opt
}
