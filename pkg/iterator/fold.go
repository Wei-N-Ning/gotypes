package iterator

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

// Scan
// In the case of Scan() the state (init R) is mutable therefore the caller can model
// logic that has side effect encapsulated in the mutable state
// see the rust version:
// https://doc.rust-lang.org/std/iter/trait.Iterator.html#method.scan
