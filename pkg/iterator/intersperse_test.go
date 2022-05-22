package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersperse(t *testing.T) {
	xs := Intersperse(RangeInclusive(1, 4), 100).Slice()
	assert.Equal(t, []int{1, 100, 2, 100, 3, 100, 4}, xs)
}
