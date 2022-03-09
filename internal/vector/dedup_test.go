package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDedup(t *testing.T) {
	t.Run("empty vector", func(t *testing.T) {
		vs := FromValues[int]()
		dvs := Dedup(vs)
		assert.Equal(t, []int{}, dvs.ToSlice())
	})
	t.Run("expect elements removed", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3}, Dedup(FromValues(1, 2, 2, 3)).ToSlice())
		assert.Equal(t, []int{1}, Dedup(FromValues(1, 1, 1, 1)).ToSlice())
	})
	t.Run("expect no effect", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 4, 2, 3}, Dedup(FromValues(1, 2, 4, 2, 3)).ToSlice())
	})
}

func TestDedupIndex(t *testing.T) {
	t.Run("empty vector", func(t *testing.T) {
		vs := FromValues[int]()
		DedupInplace(&vs)
		assert.Equal(t, []int{}, vs.ToSlice())
	})
	t.Run("expect elements removed", func(t *testing.T) {
		vs := FromValues[int](1, 1, 2, 2, 3, 4, 4, 5)
		DedupInplace(&vs)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, vs.ToSlice())
	})
	t.Run("expect elements reduced to singleton", func(t *testing.T) {
		vs := FromValues[int](1, 1, 1, 1, 1, 1, 1, 1)
		DedupInplace(&vs)
		assert.Equal(t, []int{1}, vs.ToSlice())
	})
}
