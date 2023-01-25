package option

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFmapOverGenericOptionSomes(t *testing.T) {
	xs := []Option[int]{
		Some(1),
		Some(2),
	}
	for _, x := range xs {
		y := Fmap(x, func(x int) int { return x + 100 })
		fmt.Printf("%v %v\n", y.IsSome(), y.Unwrap())
	}
}

func TestFmapOverGenericOptionNones(t *testing.T) {
	xs := []Option[int]{
		None[int](),
		Some(1),
		None[int](),
	}
	for _, x := range xs {
		y := Fmap(x, func(x int) int { return x + 100 })
		fmt.Printf("%v\n", y.IsSome())
	}
}

func TestUnwrapOr(t *testing.T) {
	require.Equal(t, 0, None[int]().UnwrapOr(0))
	require.Equal(t, 1, Some(1).UnwrapOr(0))
}
