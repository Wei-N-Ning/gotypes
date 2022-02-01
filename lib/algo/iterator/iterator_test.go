package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeatTakeCount(t *testing.T) {
	iter := Repeat(0).Take(4)
	x := iter.Count()
	assert.Equal(t, 4, x)
}
