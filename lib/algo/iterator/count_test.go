package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountIf(t *testing.T) {
	x := FromSlice([]int{1, 2, 3, 4, 5}).CountIf(func(x int) bool { return x%2 == 1 })
	assert.Equal(t, 3, x)
}
