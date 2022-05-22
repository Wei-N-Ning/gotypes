package vector

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEachEffect(t *testing.T) {
	t.Run("no side effects from an empty vector", func(t *testing.T) {
		vec := WithCapacity[string](64)
		effect := 0
		vec.ForEach(func(s string) {
			effect += len(s)
		})
		assert.Equal(t, 0, effect)
	})
	t.Run("expect side effects from non-empty vectors", func(t *testing.T) {
		vec := FromValues("_asd", "asd", "bsd", "_bsd")
		effect := 0
		vec.ForEach(func(s string) {
			if !strings.HasPrefix(s, "_") {
				effect += len(s)
			}
		})
		assert.Equal(t, 6, effect)
	})

}
