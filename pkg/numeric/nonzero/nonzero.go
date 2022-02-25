package nonzero

import (
	"constraints"
	"github.com/Wei-N-Ning/gotypes/pkg/option"
)

// NonZero a newtype that ensures the value encapsulated is not zero;
// the value encapsulated can be off integer or float
// for the approx type constraint, see:
// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md
type NonZero[T constraints.Float | constraints.Integer] struct {
	x T
}

func (val NonZero[T]) Unwrap() T {
	return val.x
}

// Add returns None if the result is 0
func Add[T constraints.Float | constraints.Integer](lhs NonZero[T], rhs NonZero[T]) option.Option[NonZero[T]] {
	return TryNewValue(lhs.x + rhs.x)
}

// UnsafeNewValue crashes if the given value is zero
func UnsafeNewValue[T constraints.Float | constraints.Integer](x T) NonZero[T] {
	if x == 0 {
		panic(x)
	}
	return NonZero[T]{x: x}
}

// TryNewValue returns None if the given value is Zero
func TryNewValue[T constraints.Float | constraints.Integer](x T) option.Option[NonZero[T]] {
	if x == 0 {
		return option.None[NonZero[T]]()
	}
	return option.Some(NonZero[T]{x: x})
}
