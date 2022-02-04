package iterator

import (
	. "go-types-nw/lib/algo/option"
)

// Reduce to apply f to each pair of elements concurrently and feed the result back to the input
// o o o o o o o o o o o ....
// ^^^ ^^^ ^^^ ^^^ ^^^
//  o   o   o   o   o
//  ^^^^^   ^^^^^
//    o       o
//    ^^^^^^^^^
//        o
func Reduce[T any](iter Iterator[T], init T, f func(T, T) T) T {
	reader, writer := TailAppender[T](1024)
	size := 0
	// the first pass: to fill the tail-appender and figure out the size
	for {
		first := iter.Next()
		if !first.IsSome() {
			break
		}
		second := iter.Next()
		if !second.IsSome() {
			init = f(init, first.Unwrap())
			break
		}
		go func() {
			writer <- Some(f(first.Unwrap(), second.Unwrap()))
		}()
		size += 1
	}
	// the second pass and onward: use the size to drive the reduction
	for {
		// terminating condition
		if size == 0 {
			break
		}
		aggregator := make(chan T, 1024)
		for i := 0; i < size/2; i++ {
			var first T = reader.Next().Unwrap()
			var second T = reader.Next().Unwrap()
			go func() {
				aggregator <- f(first, second)
			}()
		}
		// handle the tail element
		if size%2 == 1 {
			init = f(init, reader.Next().Unwrap())
		}
		size = size / 2
		for i := 0; i < size; i++ {
			writer <- Some(<-aggregator)
		}
		close(aggregator)
	}
	return init
}
