package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortVector(t *testing.T) {
	t.Run("empty vector, expect no effect", func(t *testing.T) {
		v := FromValues[int]()
		Sort(&v)
		assert.True(t, v.Empty())
		v1 := FromValues(1)
		Sort(&v1)
		assert.Equal(t, []int{1}, v1.ToSlice())
	})
	t.Run("expect element order", func(t *testing.T) {
		v := FromValues(3, 1, 4, 1, 5, 9, 2, 6)
		Sort(&v)
		assert.Equal(t, []int{1, 1, 2, 3, 4, 5, 6, 9}, v.ToSlice())
	})
}
