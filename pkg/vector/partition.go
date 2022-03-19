package vector

// Partition changes the positions of the elements in the given vector:
// those that pass f are moved to the head positions;
// those that fail to the tail positions.
// Partition is NOT stable - it doesn't maintain the element order in each partition.
// Returns the partition cursor: the index of the first element of the second partition,
// (it contains all the elements failing f)
func (vec *Vector[T]) Partition(f func(elem T) bool) int {
	if vec.Size() < 2 {
		return 0
	}
	p := 0
	for idx := 0; idx < vec.Size(); idx++ {
		if f(vec.xs[idx]) {
			if idx != p {
				vec.Swap(idx, p)
			}
			p += 1
		}
	}
	return p
}
