package iterator

// Reduce to apply f to each pair of elements concurrently and feed the result back to the input
// o o o o o o o o o o o ....
// ^^^ ^^^ ^^^ ^^^ ^^^
//  o   o   o   o   o
//  ^^^^^   ^^^^^
//    o       o
//    ^^^^^^^^^
//        o
func (iter Iterator[T]) Reduce(init T, f func(T, T) T) T {
	rw := make(chan T, 1024)
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
			rw <- f(first.Unwrap(), second.Unwrap())
		}()
		size += 1
	}
	// the second pass and onward: use the size to drive the reduction
	for {
		// terminating condition
		if size == 0 {
			break
		}
		for i := 0; i < size/2; i++ {
			var first T = <-rw
			var second T = <-rw
			go func() {
				rw <- f(first, second)
			}()
		}
		// handle the tail element
		if size%2 == 1 {
			init = f(init, <-rw)
		}
		size = size / 2
	}
	return init
}
