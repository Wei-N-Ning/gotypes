package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCycleFromEmpty(t *testing.T) {
	iter := Empty[int]().Cycle()
	assert.False(t, iter.Next().IsSome())
}

func TestCycleNTimesExpectElements(t *testing.T) {
	xs := FromSlice([]int{1, 10, 101}).Cycle().Take(8).Slice()
	assert.Equal(t, []int{1, 10, 101, 1, 10, 101, 1, 10}, xs)
}
