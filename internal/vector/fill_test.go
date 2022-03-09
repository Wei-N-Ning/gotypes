package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFill(t *testing.T) {
	t.Run("empty vector, expect no effect", func(t *testing.T) {
		v := WithCapacity[int](4)
		v.Fill(100)
		assert.Equal(t, []int{}, v.ToSlice())
	})
	t.Run("singleton vector", func(t *testing.T) {
		v := FromValues(1)
		v.Fill(100)
		assert.Equal(t, []int{100}, v.ToSlice())
	})
	t.Run("expect all elements with new value", func(t *testing.T) {
		v := FromValues(1, 2, 3, 4, 5, 6)
		v.Fill(100)
		assert.Equal(t, []int{100, 100, 100, 100, 100, 100}, v.ToSlice())
	})
}

func TestFillRange(t *testing.T) {
	t.Run("empty vector, expect no effect", func(t *testing.T) {
		v := WithCapacity[int](4)
		v.FillRange(100, 0, 10)
		assert.Equal(t, []int{}, v.ToSlice())
	})
	t.Run("singleton vector", func(t *testing.T) {
		v := FromValues(1)
		v.FillRange(100, 0, 10)
		assert.Equal(t, []int{100}, v.ToSlice())
	})
	t.Run("expect range adjusted", func(t *testing.T) {
		v := FromValues(1, 2, 3, 4, 5, 6)
		v.FillRange(100, 0, 1000)
		assert.Equal(t, []int{100, 100, 100, 100, 100, 100}, v.ToSlice())
	})
	t.Run("expect elements partially filled", func(t *testing.T) {
		v := FromValues(1, 2, 3, 4, 5, 6)
		v.FillRange(100, 2, 4)
		assert.Equal(t, []int{1, 2, 100, 100, 5, 6}, v.ToSlice())
	})
}
