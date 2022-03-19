package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelect(t *testing.T) {
	as := FromValues(1, 2, 3)
	bs := FromValues(-1, -2, -3)
	cs := FromValues(10, 20, 30)
	ds := FromValues(-10, -20, -30)
	it := Select(16, as, bs, cs, ds)
	assert.Equal(t, 3*4, it.Count())
}
