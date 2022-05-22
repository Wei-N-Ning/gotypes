package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushExpectSizeGrowth(t *testing.T) {
	vec := WithCapacity[string](0)
	vec.Push("a")
	assert.Equal(t, 1, vec.Size())
	vec.Push("b")
	assert.Equal(t, 2, vec.Size())
}

func TestPushExpectCapacityGrowth(t *testing.T) {
	vec := WithCapacity[string](0)
	vec.Push("a")
	assert.Equal(t, 2, vec.Capacity())
	vec.Push("b")
	assert.Equal(t, 2, vec.Capacity())
	vec.Push("c")
	assert.Equal(t, 6, vec.Capacity())

	v2 := WithCapacity[string](16)
	v2.Push("a")
	v2.Push("b")
	assert.Equal(t, 16, v2.Capacity())
}

func TestPushPop(t *testing.T) {
	vec := WithCapacity[string](0)
	vec.Push("a")
	vec.Push("b")
	assert.Equal(t, "b", vec.TryPop().Unwrap())
	assert.Equal(t, "a", vec.TryPop().Unwrap())
	assert.False(t, vec.TryPop().IsSome())
	assert.Equal(t, 0, vec.Size())
}
