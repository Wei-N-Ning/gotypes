package vector

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func Sort[T constraints.Ordered](vec *Vector[T]) {
	if vec.Size() < 2 {
		return
	}
	slices.Sort(vec.xs[:vec.Size()])
}
