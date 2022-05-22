package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHead(t *testing.T) {
	t.Run("expect None", func(t *testing.T) {
		xs := FromValues[int]()
		assert.False(t, xs.Head().IsSome())
	})
	t.Run("expect Some", func(t *testing.T) {
		xs := FromValues(1)
		assert.True(t, xs.Head().IsSome())
	})
}

func TestTail(t *testing.T) {
	t.Run("expect empty vector", func(t *testing.T) {
		xs := FromValues[int]()
		tl := xs.Tail()
		assert.True(t, tl.Empty())
	})
	t.Run("expect empty vector from singleton vector", func(t *testing.T) {
		xs := FromValues[int](1)
		tl := xs.Tail()
		assert.True(t, tl.Empty())
	})
	t.Run("expect tail elements", func(t *testing.T) {
		xs := FromValues(1, 2, 3, 4, 5, 6)
		tl := xs.Tail()
		assert.Equal(t, []int{2, 3, 4, 5, 6}, tl.ToSlice())
	})
}
