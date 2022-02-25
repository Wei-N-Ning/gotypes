package vector

// vector encapsulates a slice and provides the algorithms (push, pop, for_each, map...)
// that are missing from the slices
// it offers the method to reserve and shrink the memory it occupies, improving the
// performance

// note the Go type reference suggests: type Vector[T any] []T
// this looks short and handy but has a few flaws:
// - pop/try-pop is O(n) instead of O(1) (since I have to resort to slice manipulation)
// - calling the "normal" functions such as len/append/cap etc. requires a type parameter
// - the growth factor is not explicitly set

const vectorGrowthFactor int = 2

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

func (vec *Vector[T]) Reserve(additional int) {
	if additional > 0 {
		vec.reallocate(vec.capacity + additional)
	}
}

func (vec *Vector[T]) reallocate(newCapacity int) {
	xs := make([]T, newCapacity)
	copy(xs, vec.xs)
	vec.xs = xs
	vec.capacity = newCapacity
}

func (vec *Vector[T]) ShrinkToFit() {
	vec.reallocate(vec.size)
}

func (vec *Vector[T]) ToSlice() []T {
	return vec.xs[:vec.size]
}