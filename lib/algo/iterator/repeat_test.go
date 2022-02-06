package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeatTakeCount(t *testing.T) {
	assert.Equal(t, 4, Repeat(0).Take(4).Count())
	assert.Equal(t, 4, Repeat("a").Take(4).Count())
	assert.Equal(t, 4, Repeat([]string{"a"}).Take(4).Count())
}
