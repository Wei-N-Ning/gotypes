package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChainIterators(t *testing.T) {
	xs := Chain(FromSlice([]int{1, 2, 3}), FromSlice([]int{}), FromSlice([]int{10, 20, 30})).Slice()
	assert.Equal(t, []int{1, 2, 3, 10, 20, 30}, xs)
}
