package nonzero

import "constraints"
import "go-types-nw/lib/algo/option"

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

func TryNewValue[T constraints.Float | constraints.Integer](x T) option.Option[NonZero[T]] {
	if x == 0 {
		return option.None[NonZero[T]]()
	}
	return option.Some(NonZero[T]{x: x})
}
