package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateVectorWithCapacity(t *testing.T) {
	t.Run("an empty vector can have non-zero capacity", func(t *testing.T) {
		vec := WithCapacity[int](34)
		assert.Equal(t, 0, vec.Size())
		assert.Equal(t, 34, vec.Capacity())
		assert.True(t, vec.Empty())
	})

}

func TestCreateFromValues(t *testing.T) {
	t.Run("singleton", func(t *testing.T) {
		vec := FromValues("a")
		assert.Equal(t, 1, vec.Size())
	})
	t.Run("from variadic arguments", func(t *testing.T) {
		vec := FromValues(1, 2, 3)
		assert.Equal(t, 3, vec.Size())
	})
}
