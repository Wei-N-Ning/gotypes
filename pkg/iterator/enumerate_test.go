package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithIndexExpectIndices(t *testing.T) {
	xs := Map(
		WithIndex(Repeat("a").Take(4)),
		func(p Pair[int64, string]) int64 {
			return p.First
		}).Slice()
	assert.Equal(t, []int64{0, 1, 2, 3}, xs)
}
