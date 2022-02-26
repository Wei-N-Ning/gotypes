package vector

import "golang.org/x/exp/constraints"

func min[T constraints.Ordered](lhs T, rhs T) T {
	if lhs < rhs {
		return lhs
	} else {
		return rhs
	}
}
