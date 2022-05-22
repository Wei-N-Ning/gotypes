package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllWithEmptyIterator(t *testing.T) {
	assert.True(t, All(Empty[int](), func(int) bool { return false }))
}

func TestAll(t *testing.T) {
	assert.True(t, All(Range(1, 100), func(x int) bool { return x > 0 }))
}

func TestAny(t *testing.T) {
	assert.True(t, Any(Range(-3, 100), func(x int) bool { return x > 0 }))
}
