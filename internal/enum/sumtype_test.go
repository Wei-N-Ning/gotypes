package enum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumtypeConstraint(t *testing.T) {
	assert.Equal(t, 1, ErrorCode(InvalidPassword{}))
	assert.Equal(t, 2, ErrorCode(NotFound{}))
	assert.Equal(t, 3, ErrorCode(Timeout{}))
	assert.Equal(t, 4, ErrorCode(InternalError{}))

	// won't compile: error does not implement ProdError
	//ErrorCode(fmt.Errorf("won't work"))
}
