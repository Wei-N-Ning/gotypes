package vector

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"sort"
)

func Sort[T constraints.Ordered](vec *Vector[T]) {
	if vec.Size() < 2 {
		return
	}
	slices.Sort(vec.xs[:vec.Size()])
}

func SortBy[T constraints.Ordered](vec *Vector[T], less func(a, b int) bool) {
	if vec.Size() < 2 {
		return
	}
	slices.Sort(vec.xs[:vec.Size()])
	sort.Slice(vec.xs[:vec.Size()], less)
}
