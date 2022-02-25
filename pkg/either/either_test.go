package either

import (
	"fmt"
	"testing"
)

func TestFmapOverGenericEitherExpectOk(t *testing.T) {
	xs := []Either[int]{
		Ok[int](1),
		Ok[int](2),
	}
	for _, x := range xs {
		y := Fmap(x, func(x int) int { return x + 100 })
		fmt.Printf("%v %v\n", y.IsOk(), y.Unwrap())
	}
}

func TestFmapOverGenericEitherExpectErr(t *testing.T) {
	xs := []Either[int]{
		Ok[int](1),
		Err[int](fmt.Errorf("")),
		Ok[int](2),
	}
	for _, x := range xs {
		y := Fmap(x, func(x int) int { return x + 100 })
		fmt.Printf("%v\n", y.IsOk())
	}
}
