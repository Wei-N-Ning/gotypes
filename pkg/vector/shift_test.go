package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShiftLeft(t *testing.T) {
	t.Run("empty vector, expect new size with default element value", func(t *testing.T) {
		v := WithCapacity[int](4)
		assert.Equal(t, []int{}, v.ToSlice())
		v.ShiftLeft(3)
		assert.Equal(t, []int{0, 0, 0}, v.ToSlice())
	})
	t.Run("singleton vector, expect elements", func(t *testing.T) {
		v := FromValues(1)
		assert.Equal(t, []int{1}, v.ToSlice())
		v.ShiftLeft(3)
		assert.Equal(t, []int{0, 0, 0, 1}, v.ToSlice())
	})
	t.Run("expect elements", func(t *testing.T) {
		v := FromValues(1, 2, 3)
		assert.Equal(t, []int{1, 2, 3}, v.ToSlice())
		v.ShiftLeft(5)
		assert.Equal(t, []int{0, 0, 0, 0, 0, 1, 2, 3}, v.ToSlice())
	})
}

func TestShiftRangeLeft(t *testing.T) {
	t.Run("empty vector, expect new size with default element value", func(t *testing.T) {
		v := WithCapacity[int](4)
		assert.Equal(t, []int{}, v.ToSlice())
		v.ShiftRangeLeft(2, 3)
		assert.Equal(t, []int{0, 0, 0}, v.ToSlice())
	})
	t.Run("singleton vector, expect elements", func(t *testing.T) {
		v := FromValues(1)
		assert.Equal(t, []int{1}, v.ToSlice())
		v.ShiftRangeLeft(0, 3)
		assert.Equal(t, []int{0, 0, 0, 1}, v.ToSlice())
	})
	t.Run("expect elements", func(t *testing.T) {
		v := FromValues(1, 2, 3)
		assert.Equal(t, []int{1, 2, 3}, v.ToSlice())
		v.ShiftRangeLeft(2, 5)
		assert.Equal(t, []int{1, 2, 0, 0, 0, 0, 0, 3}, v.ToSlice())
	})
}

func TestShiftRangeRight(t *testing.T) {
	t.Run("empty vector, expect no effect", func(t *testing.T) {
		v := WithCapacity[int](4)
		assert.Equal(t, []int{}, v.ToSlice())
		v.ShiftRangeRight(2, 3)
		assert.Equal(t, []int{}, v.ToSlice())
	})
	t.Run("singleton vector, expect no effect", func(t *testing.T) {
		v := FromValues(1)
		assert.Equal(t, []int{1}, v.ToSlice())
		v.ShiftRangeRight(0, 3)
		assert.Equal(t, []int{0}, v.ToSlice())
	})
	t.Run("position > distance", func(t *testing.T) {
		v := FromValues(1, 2, 3, 4, 5, 6, 7, 8, 9)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, v.ToSlice())
		v.ShiftRangeRight(5, 4)
		assert.Equal(t, []int{1, 6, 7, 8, 9, 0, 0, 0, 0}, v.ToSlice())
	})
	t.Run("position == distance", func(t *testing.T) {
		v := FromValues(1, 2, 3, 4, 5, 6, 7, 8, 9)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, v.ToSlice())
		v.ShiftRangeRight(5, 5)
		assert.Equal(t, []int{6, 7, 8, 9, 0, 0, 0, 0, 0}, v.ToSlice())
	})
	t.Run("position < distance", func(t *testing.T) {
		v := FromValues(1, 2, 3, 4, 5, 6, 7, 8, 9)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, v.ToSlice())
		v.ShiftRangeRight(5, 1000)
		assert.Equal(t, []int{6, 7, 8, 9, 0, 0, 0, 0, 0}, v.ToSlice())
	})
}
