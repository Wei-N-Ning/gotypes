package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIteratorToString(t *testing.T) {
	s := RangeInclusive(1, 5).String(",")
	assert.Equal(t, "1,2,3,4,5", s)
}
