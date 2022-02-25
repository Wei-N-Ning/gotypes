package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChainIterators(t *testing.T) {
	xs := FromSlice([]int{1, 2, 3}).Chain(Empty[int]()).Chain(FromSlice([]int{10, 20, 30})).Slice()
	assert.Equal(t, []int{1, 2, 3, 10, 20, 30}, xs)
}
