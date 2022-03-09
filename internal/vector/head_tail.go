package vector

import "github.com/Wei-N-Ning/gotypes/pkg/option"

func (vec Vector[T]) Head() option.Option[T] {
	if vec.Empty() {
		return option.None[T]()
	}
	return option.Some(vec.xs[0])
}

func (vec Vector[T]) Tail() Vector[T] {
	if vec.Empty() {
		return FromValues[T]()
	}
	ys := vec.xs[1:vec.size]
	return FromSlice[T](ys)
}
