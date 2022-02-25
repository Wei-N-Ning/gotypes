package fs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirIterExpectFileCount(t *testing.T) {
	num := DirIter("/var/tmp").Count()
	assert.Greater(t, num, 1)
}
