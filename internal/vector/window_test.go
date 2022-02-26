package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVectorSlidingWindow(t *testing.T) {
	t.Run("empty vector, expect error", func(t *testing.T) {
		v := WithCapacity[int](0)
		_, err := Window(v, 12)
		assert.Error(t, err)
	})
	t.Run("invalid window size, expect error", func(t *testing.T) {
		v := FromValues(1, 2)
		_, err := Window(v, 0)
		assert.Error(t, err)
	})
	t.Run("Window size is larger than vec size, expect error", func(t *testing.T) {
		v := FromValues(1, 2)
		_, err := Window(v, 12)
		assert.Error(t, err)
		_, err = Window(v, 2)
		assert.Error(t, err)
	})
	t.Run("expect window elements", func(t *testing.T) {
		v := FromValues(1, 2, 3, 4, 5, 6, 7)
		var vs Vector[*Vector[int]]
		vs, err := Window(v, 3)
		assert.NoError(t, err)
		out := WithCapacity[int](1024)
		vs.ForEach(func(es *Vector[int]) {
			out.Extend(*es)
		})
		assert.Equal(t, []int{1, 2, 3, 2, 3, 4, 3, 4, 5, 4, 5, 6, 5, 6, 7}, out.ToSlice())
	})
}
