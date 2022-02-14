package numeric

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewType(t *testing.T) {
	fee := NewUncheck(1.2)
	assert.Equal(t, 1.2, fee.Unwrap())
}
