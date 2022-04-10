package vector

// ShiftLeft moves all the elements to the left, e.g.
// given [1, 2, 3], distance = 3
// result in [0, 0, 0, 1, 2, 3]
// * the original positions are filled by the default value of T
func (vec *Vector[T]) ShiftLeft(distance uint) {
	vec.ShiftRangeLeft(0, distance)
}

// ShiftRangeLeft moves all the elements within range [position, size) to the left
func (vec *Vector[T]) ShiftRangeLeft(position uint, distance uint) {
	vec.Reserve(int(distance))
	for idx := vec.Size() - 1; idx >= int(position); idx-- {
		vec.xs[idx+int(distance)] = vec.xs[idx]
	}
	var defaultValue T
	vec.FillRange(defaultValue, int(position), int(distance))
	vec.size += int(distance)
}

func (vec *Vector[T]) ShiftRangeRight(position int, distance uint) {
	if position < 0 || position >= vec.Size() {
		return
	}
	if position < int(distance) {
		distance = uint(position)
	}
	var defaultValue T
	vec.FillRange(defaultValue, position-int(distance), position)
	for idx := position; idx < vec.Size(); idx++ {
		vec.xs[idx-int(distance)] = vec.xs[idx]
	}
	vec.FillRange(defaultValue, position, vec.Size())
}
