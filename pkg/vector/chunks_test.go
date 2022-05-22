package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorChunks(t *testing.T) {
	t.Run("empty vector, expect error", func(t *testing.T) {
		v := WithCapacity[int](0)
		_, err := Chunks(v, 12)
		assert.Error(t, err)
	})
	t.Run("invalid chunk size, expect error", func(t *testing.T) {
		v := FromValues(1, 2)
		_, err := Chunks(v, 0)
		assert.Error(t, err)
	})
	t.Run("chunk size is larger than vec size", func(t *testing.T) {
		v := FromValues(1, 2)
		var vs Vector[*Vector[int]]
		vs, err := Chunks(v, 12)
		assert.NoError(t, err)
		assert.Equal(t, 1, vs.Size())
		assert.Equal(t, []int{1, 2}, vs.TryPop().Unwrap().ToSlice())
	})
	t.Run("expect tail elements", func(t *testing.T) {
		v := FromValues(1, 2, 3, 4, 5, 6)
		var vs Vector[*Vector[int]]
		vs, err := Chunks(v, 4)
		assert.NoError(t, err)
		assert.Equal(t, 2, vs.Size())
		assert.Equal(t, []int{5, 6}, vs.TryPop().Unwrap().ToSlice())
	})
}
