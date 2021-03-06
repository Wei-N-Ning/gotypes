package iterator

import (
	. "github.com/Wei-N-Ning/gotypes/pkg/option"
)

func TailAppender[T any](buffer int64) (Iterator[T], chan<- Option[T]) {
	ch := make(chan Option[T], buffer)
	var reader <-chan Option[T] = ch
	return Iterator[T]{ch: reader, inner: ch}, ch
}
