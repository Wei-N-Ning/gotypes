package iterator

import (
	. "go-types-nw/lib/algo/option"
)

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

func RepeatN[T any](x T, num int) Iterator[T] {
	return Repeat(x).Take(num)
}
