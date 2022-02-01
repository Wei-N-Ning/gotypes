package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRangeToSlice(t *testing.T) {
	xs := Range(0, 5).Slice()
	assert.Equal(t, []int{0, 1, 2, 3, 4}, xs)
}
