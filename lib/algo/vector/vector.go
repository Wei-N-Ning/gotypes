package vector

// vector encapsulates a slice and provides the algorithms (push, pop, for_each, map...)
// that are missing from the slices
// it offers the method to reserve and shrink the memory it occupies, improving the
// performance

type Vector[T any] struct {
	xs       []T
	capacity int
	size     int
}

func (vec *Vector[T]) Capacity() int {
	return vec.capacity
}

func (vec *Vector[T]) Size() int {
	return vec.size
}

func (vec *Vector[T]) Empty() bool {
	return vec.size == 0
}

func (vec *Vector[T]) grow(newCapacity int) {
	xs := make([]T, newCapacity)
	copy(xs, vec.xs)
	vec.xs = xs
	vec.capacity = newCapacity
}
