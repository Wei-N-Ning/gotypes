package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func addTwo(x int, y int) int {
	return x + y
}

func TestReduceFromEmptyExpectInitValue(t *testing.T) {
	x := Reduce(Empty[int](), 101, addTwo)
	assert.Equal(t, 101, x)
}

func TestReduceSingletonSliceExpectValue(t *testing.T) {
	x := Reduce(FromSlice([]int{1}), 101, addTwo)
	assert.Equal(t, 102, x)
}

func TestReducePairExpectValue(t *testing.T) {
	x := Reduce(FromSlice([]int{1, -1}), 101, addTwo)
	assert.Equal(t, 101, x)
}

func TestReduceTripleExpectValue(t *testing.T) {
	x := Reduce(FromSlice([]int{1, -1, 10}), 0, addTwo)
	assert.Equal(t, 10, x)
}

func TestReduceFromSliceExpectValue(t *testing.T) {
	x := Reduce(FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), 0, addTwo)
	assert.Equal(t, (1+10)*10/2, x)
}
