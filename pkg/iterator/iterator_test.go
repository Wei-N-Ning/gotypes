package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTailExpectValues(t *testing.T) {
	xs := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}).Tail().Take(5).Tail().Slice()
	assert.Equal(t, []int{3, 4, 5, 6}, xs)

	ys := FromSlice([]int{}).Tail()
	assert.False(t, ys.Next().IsSome())
}

func TestLastExpectNone(t *testing.T) {
	x := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}).Last()
	assert.Equal(t, 8, x.Unwrap())

	assert.False(t, FromSlice([]int{}).Last().IsSome())
}

func TestAdvanceByFromEmpty(t *testing.T) {
	var ys []int = Empty[int]().AdvanceBy(10).Slice()
	assert.Equal(t, 0, len(ys))
}

func TestAdvanceBy(t *testing.T) {
	xs := FromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8}).AdvanceBy(4).Slice()
	assert.Equal(t, []int{5, 6, 7, 8}, xs)
}
