package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapExpectSizeCapacity(t *testing.T) {
	t.Run("returned vector has the same length and capacity as the original one", func(t *testing.T) {
		v1 := FromValues("e1", "e1m1", "map:e1m1")
		v2 := Map(v1, func(x string) int { return len(x) })
		assert.Equal(t, v1.Capacity(), v2.Capacity())
		assert.Equal(t, v1.Size(), v2.Size())

		assert.Equal(t, 8, v2.TryPop().Unwrap())
		assert.Equal(t, 4, v2.TryPop().Unwrap())
		assert.Equal(t, 2, v2.TryPop().Unwrap())
		assert.False(t, v2.TryPop().IsSome())
	})
}
