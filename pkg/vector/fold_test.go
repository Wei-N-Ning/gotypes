package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorFold(t *testing.T) {
	add := func(a, b int) int {
		return a + b
	}
	t.Run("fold empty vector, expect init", func(t *testing.T) {
		v := FromValues[int]()
		assert.Equal(t, 10, Fold(v, 10, add))
	})
	t.Run("expect return value", func(t *testing.T) {
		v := FromValues(1, 2, 3)
		assert.Equal(t, 16, Fold(v, 10, add))
	})
}
