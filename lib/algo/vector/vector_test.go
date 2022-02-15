package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromEmptyVectorToSlice(t *testing.T) {
	vec := WithCapacity[string](64)
	xs := vec.ToSlice() // xs[0:0] -> []T{}
	assert.Equal(t, []string{}, xs)
}

func TestFromVectorToSlice(t *testing.T) {
	vec := WithCapacity[string](64)
	vec.Push("a")
	vec.Push("b")
	vec.Push("c")
	assert.Equal(t, []string{"a", "b", "c"}, vec.ToSlice())
}
