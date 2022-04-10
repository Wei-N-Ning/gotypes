package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeatElement(t *testing.T) {
	t.Run("repeat 0 time, expect an empty vector", func(t *testing.T) {
		v := Repeat("a", 0)
		assert.True(t, v.Empty())
	})
	t.Run("expect elements", func(t *testing.T) {
		v := Repeat("a", 4)
		assert.Equal(t, []string{"a", "a", "a", "a"}, v.ToSlice())
	})
}
