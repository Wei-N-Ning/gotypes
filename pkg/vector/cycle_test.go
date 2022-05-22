package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCycle(t *testing.T) {
	t.Run("expect elements", func(t *testing.T) {
		v := FromValues(1, 2)
		vs := Cycle(v, 2)
		assert.Equal(t, []int{1, 2, 1, 2}, vs.ToSlice())
	})
}
