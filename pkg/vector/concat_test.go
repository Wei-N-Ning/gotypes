package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcatVectors(t *testing.T) {
	t.Run("concat one vector, expect no effect", func(t *testing.T) {
		v1 := FromValues(1, 2)
		v2 := Concat[int](v1)
		assert.Equal(t, v1.ToSlice(), v2.ToSlice())
	})
	t.Run("concat 3 vectors, expect new elements", func(t *testing.T) {
		xs := Concat[int](FromValues(0, 0), FromValues(-1, -1), FromValues(10, 10))
		assert.Equal(t, xs.ToSlice(), []int{0, 0, -1, -1, 10, 10})
	})
}
