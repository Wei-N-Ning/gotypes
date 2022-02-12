package vector

import (
	"go-types-nw/lib/algo/option"
)

// vector encapsulates a slice and provides the algorithms (push, pop, for_each, map...)
// that are missing from the slices
// it offers the method to reserve and shrink the memory it occupies, improving the
// performance

type Vector[T any] struct {
	xs       []T
	capacity int
	size     int
}

func WithCapacity[T any](cap int) Vector[T] {
	vec := Vector[T]{xs: make([]T, cap), capacity: cap, size: 0}
	if cap > 0 {
		vec.grow(cap)
	}
	return vec
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

func (vec *Vector[T]) Push(x T) {
	newSize := vec.size + 1
	if newSize >= vec.capacity {
		vec.grow(newSize * 2)
	}
	vec.size += 1
	vec.xs[vec.size-1] = x
}

func (vec *Vector[T]) TryPop() option.Option[T] {
	if vec.size == 0 {
		return option.None[T]()
	}
	x := vec.xs[vec.size-1]
	vec.size -= 1
	return option.Some(x)
}

func (vec *Vector[T]) grow(newCapacity int) {
	xs := make([]T, newCapacity)
	copy(xs, vec.xs)
	vec.xs = xs
	vec.capacity = newCapacity
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
