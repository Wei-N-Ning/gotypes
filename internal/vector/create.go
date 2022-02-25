package vector

func WithCapacity[T any](cap int) Vector[T] {
	vec := Vector[T]{xs: make([]T, cap), capacity: cap, size: 0}
	if cap > 0 {
		vec.reallocate(cap)
	}
	return vec
}

func FromSlice[T any](xs []T) Vector[T] {
	vec := WithCapacity[T](len(xs))
	copy(vec.xs, xs)
	vec.size = len(xs)
	return vec
}

func FromValues[T any](xs ...T) Vector[T] {
	return FromSlice(xs)
}
