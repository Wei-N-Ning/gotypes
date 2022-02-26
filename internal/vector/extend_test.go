package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtendExpectNewSizeCap(t *testing.T) {
	vec := FromValues(1, 2, 3)
	vec.Extend(FromValues(5, 6, 7))
	assert.Equal(t, 6, vec.Size())
	assert.Equal(t, 6, vec.Capacity())
	assert.Equal(t, vec.ToSlice(), []int{1, 2, 3, 5, 6, 7})
}
