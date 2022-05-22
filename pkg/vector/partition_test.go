package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorPartition(t *testing.T) {
	f := func(x int) bool {
		return x < 5
	}
	t.Run("empty vector, expect no effect", func(t *testing.T) {
		v := FromValues[int]()
		p := v.Partition(f)
		assert.Equal(t, 0, p)
		assert.True(t, v.Empty())
	})
	t.Run("already partitioned, expect head and tail elements", func(t *testing.T) {
		v := FromValues[int](3, 4, 5, 6, 7, 8)
		p := v.Partition(f)
		assert.Equal(t, 2, p)
		assert.Equal(t, []int{3, 4, 5, 6, 7, 8}, v.ToSlice())
	})
	t.Run("expect partition cursor", func(t *testing.T) {
		v := FromValues[int](10, 1, -3, 4, 9, 13, 0)
		// there are 4 elements passing f:
		// 1, -3, 4, 0
		// hence the partition cursor is the element at index 4
		p := v.Partition(f)
		assert.Equal(t, 4, p)
	})
	t.Run("all elements failing f, expect cursor = 0", func(t *testing.T) {
		v := FromValues[int](10, 11, 12, 13, 14)
		p := v.Partition(f)
		assert.Equal(t, 0, p)
	})
}
