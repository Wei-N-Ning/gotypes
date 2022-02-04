package iterator

// Count the number of elements
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

// CountIf applies a filter function to each element.
// It counts the number of elements passing the check.
func (iter Iterator[T]) CountIf(fi func(x T) bool) int {
	return iter.Filter(fi).Count()
}
