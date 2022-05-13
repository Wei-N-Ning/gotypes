package iterator

import "github.com/Wei-N-Ning/gotypes/pkg/option"

func FromChannel[T any](inCh <-chan T) Iterator[T] {
	ch := make(chan option.Option[T], cap(inCh))
	go func() {
		defer func() {
			ch <- option.None[T]()
			close(ch)
		}()

		for x := range inCh {
			ch <- option.Some(x)
		}
	}()
	return Iterator[T]{ch: ch}
}
