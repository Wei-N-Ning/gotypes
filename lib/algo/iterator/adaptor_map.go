package iterator

import (
	. "go-types-nw/lib/algo/option"
)

func mapImpl[T any, R any](iter Iterator[T], f func(T) R) <-chan Option[R] {
	ch := make(chan Option[R], 1024)
	go func() {
		defer close(ch)
		for {
			elem := iter.Next()
			if !elem.IsSome() {
				ch <- None[R]()
				return
			} else {
				ch <- Some[R](f(elem.Unwrap()))
			}
		}
	}()
	return ch
}

// Map applies a function to each element and produces a new value;
// The results are yield in the new iterator.
func Map[T any, R any](iter Iterator[T], f func(T) R) Iterator[R] {
	return Iterator[R]{ch: mapImpl(iter, f), inner: iter}

}

func MapReduce[T, R any](iter Iterator[T], init R, mapper func(x T) R, reducer func(R, R) R) R {
	rw := make(chan R, 1024)
	numTasks := 0
	iter.ForEach(func(x T) {
		numTasks += 1
		rw <- mapper(x)
	})
	for {
		if numTasks == 0 {
			break
		}
		for i := 0; i < numTasks/2; i++ {
			rw <- reducer(<-rw, <-rw)
		}
		if numTasks%2 == 1 {
			init = reducer(init, <-rw)
		}
		numTasks = numTasks / 2
	}
	return init
}
