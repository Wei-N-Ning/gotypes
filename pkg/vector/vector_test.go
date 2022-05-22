package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSlice(t *testing.T) {
	t.Run("an empty vector results in an empty slice", func(t *testing.T) {
		vec := WithCapacity[string](64)
		xs := vec.ToSlice() // xs[0:0] -> []T{}
		assert.Equal(t, []string{}, xs)
	})
	t.Run("expect elements from the resulting slice", func(t *testing.T) {
		vec := WithCapacity[string](64)
		vec.Push("a")
		vec.Push("b")
		vec.Push("c")
		assert.Equal(t, []string{"a", "b", "c"}, vec.ToSlice())
	})

}

func TestResetSize(t *testing.T) {
	t.Run("ensure empty", func(t *testing.T) {
		v := FromValues(1, 2, 3, 4)
		assert.False(t, v.Empty())
		assert.GreaterOrEqual(t, v.Capacity(), 4)
		v.ResetSize()
		assert.True(t, v.Empty())
		assert.Equal(t, []int{}, v.ToSlice())
		assert.GreaterOrEqual(t, v.Capacity(), 4)
	})
}

func TestClear(t *testing.T) {
	t.Run("ensure elements are gone", func(t *testing.T) {
		v := FromValues(1, 2, 3, 4)
		v.Clear()
		assert.Equal(t, []int{}, v.ToSlice())
	})
}
