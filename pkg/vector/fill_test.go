package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestVector_FillWith(t *testing.T) {
	t.Run("empty vector, expect no effect", func(t *testing.T) {
		v := WithCapacity[int](4)
		v.FillWith(100, func(x int) int { return x + 1 })
		assert.Equal(t, []int{}, v.ToSlice())
	})
	t.Run("singleton vector", func(t *testing.T) {
		v := FromValues(1)
		v.FillWith(100, func(x int) int { return x + 1 })
		assert.Equal(t, []int{101}, v.ToSlice())
	})
	t.Run("expect all elements with new value", func(t *testing.T) {
		v := FromValues(1, 2, 3, 4, 5, 6)
		v.FillWith(100, func(x int) int { return x + 1 })
		assert.Equal(t, []int{101, 101, 101, 101, 101, 101}, v.ToSlice())
	})
}

func TestVector_FillRangeWith(t *testing.T) {
	t.Run("empty vector, expect no effect", func(t *testing.T) {
		v := WithCapacity[int](4)
		v.FillRangeWith(100, 0, 10, func(x int) int { return x + 1 })
		assert.Equal(t, []int{}, v.ToSlice())
	})
	t.Run("singleton vector", func(t *testing.T) {
		v := FromValues(1)
		v.FillRangeWith(100, 0, 10, func(x int) int { return x + 1 })
		assert.Equal(t, []int{101}, v.ToSlice())
	})
	t.Run("expect range adjusted", func(t *testing.T) {
		v := FromValues(1, 2, 3, 4, 5, 6)
		v.FillRangeWith(100, 0, 1000, func(x int) int { return x + 1 })
		assert.Equal(t, []int{101, 101, 101, 101, 101, 101}, v.ToSlice())
	})
	t.Run("expect elements partially filled", func(t *testing.T) {
		v := FromValues(1, 2, 3, 4, 5, 6)
		v.FillRangeWith(100, 2, 4, func(x int) int { return x + 1 })
		assert.Equal(t, []int{1, 2, 101, 101, 5, 6}, v.ToSlice())
	})
}
