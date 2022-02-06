package iterator

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
