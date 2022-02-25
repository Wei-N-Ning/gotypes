package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapReduceEmptyExpectInitValue(t *testing.T) {
	x := MapReduce(Empty[int](), 101, func(x int) int { return x * 100 }, addTwo)
	assert.Equal(t, 101, x)
}

func TestMapReduceSingletonExpectValue(t *testing.T) {
	x := MapReduce(FromSlice([]int{1}), 101, func(x int) int { return x * 100 }, addTwo)
	assert.Equal(t, 201, x)
}

func TestMapReducePairExpectValue(t *testing.T) {
	x := MapReduce(FromSlice([]int{1, -1}), 101, func(x int) int { return x * 100 }, addTwo)
	assert.Equal(t, 101, x)
}

func TestMapReduceTripleExpectValue(t *testing.T) {
	x := MapReduce(FromSlice([]int{1, -1, 100}), 101, func(x int) int { return x * 100 }, addTwo)
	assert.Equal(t, 10101, x)
}

func TestMapReduceRangeExpectValue(t *testing.T) {
	x := MapReduce(Repeat(1).Take(100), 101, func(x int) int { return x * 1 }, addTwo)
	assert.Equal(t, 201, x)
}
