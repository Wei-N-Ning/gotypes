package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIteratorToStringExpectEmptyString(t *testing.T) {
	s := Empty[int]().String("  ")
	assert.Equal(t, "", s)
}

func TestIteratorToString(t *testing.T) {
	s := RangeInclusive(1, 5).String(",")
	assert.Equal(t, "1,2,3,4,5", s)
}
