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

// the implementation detail of Take
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

func repeatImpl[T any](x T) <-chan Option[T] {
	ch := make(chan Option[T])
	go func() {
		defer close(ch)
		for {

			ch <- Some[T](x)
		}
	}()
	return ch
}

func Repeat[T any](x T) Iterator[T] {
	return Iterator[T]{ch: repeatImpl(x)}
}

// ForEach applies a function to each element and discard the result.
func (iter Iterator[T]) ForEach(f func(T)) {
	for {
		elem := iter.Next()
		if elem.IsSome() {
			f(elem.Unwrap())
		} else {
			return
		}
	}
}

// the implementation detail of Filter
func filterImpl[T any](iter Iterator[T], f func(T) bool) <-chan Option[T] {
	ch := make(chan Option[T])
	go func() {
		defer close(ch)
		for {
			elem := iter.Next()
			if !elem.IsSome() {
				ch <- elem
				return
			}
			if f(elem.Unwrap()) {
				ch <- elem
			}
		}
	}()
	return ch
}

// Filter applies a function to each element and gets a boolean value.
// If the value is true, the element is included in the resulting new iterator, otherwise it is discarded.
func (iter Iterator[T]) Filter(f func(T) bool) Iterator[T] {
	return Iterator[T]{ch: filterImpl(iter, f), inner: iter}
}

// the implementation detail of Map
func mapImpl[T any, R any](iter Iterator[T], f func(T) R) <-chan Option[R] {
	ch := make(chan Option[R])
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

func (iter Iterator[T]) Count() int {
	num := 0
	for {
		elem := iter.Next()
		if elem.IsSome() {
			num += 1
		} else {
			break
		}
	}
	return num
}

// Fold is haskell's foldLeft
// It takes each element out of the iterator, apply a computation `f func(_acc R, _elem T) R`
// then update the init value;
// When there is no more element to process, it returns the init value as the final result.
func Fold[T any, R any](iter Iterator[T], init R, f func(_acc R, _elem T) R) R {
	for {
		elem := iter.Next()
		if elem.IsSome() {
			init = f(init, elem.Unwrap())
		} else {
			break
		}
	}
	return init
}
