package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOnce(t *testing.T) {
	x := Once("asd").Count()
	assert.Equal(t, 1, x)
}

func TestOnceWith(t *testing.T) {
	x := OnceWith(func() string { return "asd" }).Count()
	assert.Equal(t, 1, x)
}
