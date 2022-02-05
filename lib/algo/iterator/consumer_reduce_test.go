package iterator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func addTwo(x int, y int) int {
	return x + y
}

func TestReduceFromEmptyExpectInitValue(t *testing.T) {
	x := Empty[int]().Reduce(101, addTwo)
	assert.Equal(t, 101, x)
}

func TestReduceSingletonSliceExpectValue(t *testing.T) {
	x := FromSlice([]int{1}).Reduce(101, addTwo)
	assert.Equal(t, 102, x)
}

func TestReducePairExpectValue(t *testing.T) {
	x := FromSlice([]int{1, -1}).Reduce(101, addTwo)
	assert.Equal(t, 101, x)
}

func TestReduceTripleExpectValue(t *testing.T) {
	x := FromSlice([]int{1, -1, 10}).Reduce(0, addTwo)
	assert.Equal(t, 10, x)
}

func TestReduceFromSliceExpectValue(t *testing.T) {
	x := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).Reduce(0, addTwo)
	assert.Equal(t, (1+10)*10/2, x)
}

func TestReduceMiniPerf(t *testing.T) {
	const DIFFICULTY = 34
	ser := timeThis(func() {
		Fold(RangeInclusive(1, 8), 0, func(x int, y int) int {
			return fib(0*(x+y) + DIFFICULTY)
		})
	})
	fmt.Println("serial", ser)
	con := timeThis(func() {
		RangeInclusive(1, 8).Reduce(0, func(x int, y int) int {
			return fib(0*(x+y) + DIFFICULTY)
		})
	})
	fmt.Println("concur", con)
}
