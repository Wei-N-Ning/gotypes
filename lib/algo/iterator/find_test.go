package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindFromEmpty(t *testing.T) {
	found := Find(Empty[int](), func(x int) bool { return false })
	assert.False(t, found.IsSome())
}

func TestFind(t *testing.T) {
	found := Find(Range(1, 100), func(x int) bool { return x%50 == 0 })
	assert.True(t, found.IsSome())
}
