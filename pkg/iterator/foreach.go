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

// ForEachWithIndex adds the element index for the caller's convenience;
func (iter Iterator[T]) ForEachWithIndex(f func(int, T)) {
	idx := 0
	for {
		elem := iter.Next()
		if elem.IsSome() {
			f(idx, elem.Unwrap())
			idx += 1
		} else {
			return
		}
	}
}
