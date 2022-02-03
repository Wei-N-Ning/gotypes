package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChunkSliceExpectSize(t *testing.T) {
	xs := ChunkSlice(Repeat("a"), 4).Take(2).Slice()
	assert.Equal(t, [][]string{{"a", "a", "a", "a"}, {"a", "a", "a", "a"}}, xs)
}
