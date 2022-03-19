package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	t.Run("empty vector, expect new element", func(t *testing.T) {
		v := WithCapacity[int](4)
		v.Insert(1, 10)
		assert.Equal(t, []int{10}, v.ToSlice())
	})
	t.Run("singleton vector, expect new elements", func(t *testing.T) {
		v := FromValues(1)
		v.Insert(0, 10) // (10), 1
		v.Insert(1, 11) // 10, (11), 1
		v.Insert(4, -1) // 10, 11, 1, (-1)
		assert.Equal(t, []int{10, 11, 1, -1}, v.ToSlice())
	})
	t.Run("expect new elements", func(t *testing.T) {
		v := FromValues(1, 2, 3, 4, 5, 6, 7, 8)
		v.Insert(4, 999)
		assert.Equal(t, []int{1, 2, 3, 4, 999, 5, 6, 7, 8}, v.ToSlice())
	})
}
