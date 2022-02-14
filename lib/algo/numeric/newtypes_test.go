package numeric

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewType(t *testing.T) {
	fee := NewUncheck(1.2)
	assert.Equal(t, 1.2, fee.Unwrap())

	x := GasFee(1)
	var y GasFee = x + 1

	compute := func(GasFee) int {
		return 0
	}

	z := compute(1)
	w := compute(x)
	w += z

	assert.Equal(t, 2, y.Unwrap())
	assert.Equal(t, 0, w)
}
