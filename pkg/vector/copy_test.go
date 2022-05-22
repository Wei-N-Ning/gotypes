package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyExpectSizeLimit(t *testing.T) {
	t.Run("when dst has more elements", func(t *testing.T) {
		src := FromValues(1, 2, 3, 4, 5, 6)
		src.reallocate(400)
		dst := FromValues(7, 8, 9, 0, 1, 2, 3, 4, 5)
		dst.reallocate(500)
		copied := Copy(&dst, src)
		assert.Equal(t, src.Size(), copied)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 3, 4, 5}, dst.ToSlice())
	})
	t.Run("when src has more elements", func(t *testing.T) {
		src := FromValues(7, 8, 9, 0, 1, 2, 3, 4, 5)
		src.reallocate(500)
		dst := FromValues(1, 2, 3, 4, 5, 6)
		dst.reallocate(400)
		copied := Copy(&dst, src)
		assert.Equal(t, dst.Size(), copied)
		assert.Equal(t, []int{7, 8, 9, 0, 1, 2}, dst.ToSlice())
	})
}

func TestCopyAtOffset(t *testing.T) {
	t.Run("expect copying starts at index 2", func(t *testing.T) {
		src := FromValues(0, 0, 0, 1, 1, 1)
		src.reallocate(400)
		dst := FromValues(7, 8, 9, 0, 1, 2, 3, 4, 5)
		dst.reallocate(500)
		copied := CopyAt(&dst, src, 2)
		assert.Equal(t, 6, copied)
		assert.Equal(t, []int{7, 8, 0, 0, 0, 1, 1, 1, 5}, dst.ToSlice())
	})
	t.Run("start position is out of bound, not copying", func(t *testing.T) {
		src := FromValues(0, 0, 0, 1, 1, 1)
		src.reallocate(400)
		dst := FromValues(7, 8, 9, 0, 1, 2, 3, 4, 5)
		dst.reallocate(500)
		copied := CopyAt(&dst, src, 19)
		assert.Equal(t, 0, copied)
		assert.Equal(t, []int{7, 8, 9, 0, 1, 2, 3, 4, 5}, dst.ToSlice())
	})
}
