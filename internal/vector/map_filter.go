package vector

import (
	"github.com/Wei-N-Ning/gotypes/pkg/option"
)

func MapFilter[T any, R any](vec *Vector[T], f func(T) option.Option[R]) *Vector[R] {
	newVec := WithCapacity[R](vec.Capacity())
	newIdx := 0
	for idx := 0; idx < vec.size; idx++ {
		opt := f(vec.xs[idx])
		if opt.IsSome() {
			newVec.xs[newIdx] = opt.Unwrap()
			newIdx += 1
		}
	}
	newVec.size = newIdx
	return newVec
}
