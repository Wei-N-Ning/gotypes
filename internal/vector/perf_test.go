package vector

import (
	"fmt"
	"testing"
)

func TestCompareAppendPerformance(t *testing.T) {
	vec := WithCapacity[int](2048)
	appendToVec := func() {
		for idx := 0; idx < 1024; idx++ {
			vec.Push(idx)
		}
	}
	var slice []int
	appendToSlice := func() {
		for idx := 0; idx < 1024; idx++ {
			slice = append(slice, idx)
		}
	}

	fmt.Println("append to vector", timeThis(appendToVec), "us")
	fmt.Println("append to slice ", timeThis(appendToSlice), "us")
}
