package vector

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestEmptyVectorExpectNoEffect(t *testing.T) {
	vec := WithCapacity[string](64)
	effect := 0
	vec.ForEach(func(s string) {
		effect += len(s)
	})
	assert.Equal(t, 0, effect)
}

func TestExpectEffect(t *testing.T) {
	vec := FromValues("_asd", "asd", "bsd", "_bsd")
	effect := 0
	vec.ForEach(func(s string) {
		if !strings.HasPrefix(s, "_") {
			effect += len(s)
		}
	})
	assert.Equal(t, 6, effect)
}
