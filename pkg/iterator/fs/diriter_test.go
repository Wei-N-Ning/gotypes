package fs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirIterExpectFileCount(t *testing.T) {
	num := DirIter("/var/tmp").Count()
	assert.Greater(t, num, 1)
}
