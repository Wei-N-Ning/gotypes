package numeric

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNonZeroUnwrapValue(t *testing.T) {
	val := UnsafeNewValue(1)
	assert.Equal(t, 1, val.Unwrap())

}
