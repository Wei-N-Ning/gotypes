package numeric

import "constraints"

type NonZero[T constraints.Float | constraints.Integer] struct {
	x T
}

func (val NonZero[T]) Unwrap() T {
	return val.x
}

func UnsafeNewValue[T constraints.Float | constraints.Integer](x T) NonZero[T] {
	if x == 0 {
		panic(x)
	}
	return NonZero[T]{x: x}
}
