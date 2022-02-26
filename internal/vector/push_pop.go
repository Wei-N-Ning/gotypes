package vector

import (
	"github.com/Wei-N-Ning/gotypes/pkg/option"
)

func (vec *Vector[T]) Push(x T) {
	if (vec.size + 1) > vec.capacity {
		vec.reallocate((vec.size + 1) * vectorGrowthFactor)
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
