package numeric

import "golang.org/x/exp/constraints"

// source: https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md
// generic types can have methods

// this doesn't feel much better than the single-field structure approach (see nonzero)
// but it looks slightly closer to Rust's tuple-struct: struct Name(String)

type Fee[T constraints.Float] [1]T

func (f Fee[T]) Unwrap() T {
	return f[0]
}

func NewUncheck[T constraints.Float](x T) Fee[T] {
	return [1]T{x}
}

type GasFee int

type GasFeeTypeSet interface {
	GasFee

	Unwrap() int
}

func (gas GasFee) Unwrap() int {
	return int(gas)
}

func DoubleGasFee[T GasFeeTypeSet](x T) GasFee {
	return GasFee(int(x) * 2)
}
