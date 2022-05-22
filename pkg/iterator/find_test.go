package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFromEmpty(t *testing.T) {
	_, found := Find(Empty[int](), func(x int) bool { return false })
	assert.False(t, found.IsSome())
}

func TestFind(t *testing.T) {
	idx, found := Find(Range(1, 100), func(x int) bool { return x%50 == 0 })
	assert.True(t, found.IsSome())
	// the elements: 1, 2, 3, 4 .... 50, ...
	// the indices : 0, 1, 2, 3 .... 49, ...
	assert.Equal(t, 49, idx)
}
