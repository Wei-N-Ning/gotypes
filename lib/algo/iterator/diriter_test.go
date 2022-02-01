package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirIterExpectFileCount(t *testing.T) {
	num := DirIter("/tmp").Count()
	assert.Greater(t, num, 1)
}
