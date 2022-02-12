package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateVectorWithCapacity(t *testing.T) {
	vec := WithCapacity[int](34)
	assert.Equal(t, 0, vec.Size())
	assert.Equal(t, 34, vec.Capacity())
	assert.True(t, vec.Empty())
}

func TestCreateFromValues(t *testing.T) {
	vec := FromValues(1, 2, 3)
	assert.Equal(t, 3, vec.Size())
}
