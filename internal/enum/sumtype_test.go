package enum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumtypeConstraint(t *testing.T) {
	assert.Equal(t, 1, ErrorCode(InvalidPassword{}))

	// won't compile: error does not implement ProdError
	//ErrorCode(fmt.Errorf("won't work"))
}
