package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunkSliceExpectValues(t *testing.T) {
	xs := ChunkSlice(Repeat("a"), 4).Take(2).Slice()
	assert.Equal(t, [][]string{{"a", "a", "a", "a"}, {"a", "a", "a", "a"}}, xs)
}

func TestChunkSliceExpectTailValue(t *testing.T) {
	xs := ChunkSlice(Repeat("a").Take(6), 4).Slice()
	assert.Equal(t, [][]string{{"a", "a", "a", "a"}, {"a", "a"}}, xs)
}
