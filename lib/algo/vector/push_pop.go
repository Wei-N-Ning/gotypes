package vector

import "go-types-nw/lib/algo/option"

func (vec *Vector[T]) Push(x T) {
	newSize := vec.size + 1
	if newSize >= vec.capacity {
		vec.grow(newSize * vectorGrowthFactor)
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
