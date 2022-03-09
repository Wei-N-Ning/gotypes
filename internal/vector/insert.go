package vector

// Insert inserts an element at position, shifting all the elements after it to the right
func (vec *Vector[T]) Insert(idx int, x T) {
	idx = max(min(vec.Size(), idx), 0)
	if idx == vec.Size() {
		vec.Push(x)
		return
	}
	vec.ShiftRangeLeft(uint(idx), 1)
	vec.xs[idx] = x
}
