package option

import (
	"fmt"
	"testing"
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
		None[int](),
	}
	for _, x := range xs {
		y := Fmap(x, func(x int) int { return x + 100 })
		fmt.Printf("%v\n", y.IsSome())
	}
}
