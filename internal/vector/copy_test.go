package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopyExpectSizeLimit(t *testing.T) {
	t.Run("when dst has more elements", func(t *testing.T) {
		src := FromValues(1, 2, 3, 4, 5, 6)
		dst := FromValues(7, 8, 9, 0, 1, 2, 3, 4, 5)
		copied := Copy(&dst, src)
		assert.Equal(t, src.Size(), copied)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 3, 4, 5}, dst.ToSlice())
	})
	t.Run("when src has more elements", func(t *testing.T) {
		src := FromValues(7, 8, 9, 0, 1, 2, 3, 4, 5)
		dst := FromValues(1, 2, 3, 4, 5, 6)
		copied := Copy(&dst, src)
		assert.Equal(t, dst.Size(), copied)
		assert.Equal(t, []int{7, 8, 9, 0, 1, 2}, dst.ToSlice())
	})
}
