package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcatFromEmptyIteratorsExpectEmpty(t *testing.T) {
	iter := Concat(Empty[int](), Empty[int](), Empty[int](), Empty[int]())
	assert.False(t, iter.Next().IsSome())
}

func TestConcatExpectElements(t *testing.T) {
	xs := Concat(FromSlice([]int{1}), FromSlice([]int{10}), Empty[int](), FromSlice([]int{100})).Slice()
	assert.Equal(t, []int{1, 10, 100}, xs)

	ys := Concat(
		FromSlice([]int{1, 2, 3}),
		Empty[int](),
		FromSlice([]int{}),
		Empty[int](),
		FromSlice([]int{10, 20, 30}),
		Empty[int](),
	).Slice()
	assert.Equal(t, []int{1, 2, 3, 10, 20, 30}, ys)
}
