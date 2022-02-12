package nonzero

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNonZeroExpectOptional(t *testing.T) {
	assert.False(t, TryNewValue(0).IsSome())
	assert.True(t, TryNewValue(1).IsSome())
	assert.Equal(t, -1, UnsafeNewValue(-1).Unwrap())
}

func TestEnforceTypeCompliance(t *testing.T) {
	// compile error:
	// string does not implement constraints.Float|constraints.Integer
	//val := TryNewValue("a")
	//assert.False(t, val.IsSome())

	assert.False(t, TryNewValue(0.0).IsSome())
	assert.True(t, TryNewValue(0.0000001).IsSome())
	assert.True(t, TryNewValue(-0.0000001).IsSome())
}

func TestAddExpectOptional(t *testing.T) {
	assert.False(t, Add(UnsafeNewValue(-1), UnsafeNewValue(1)).IsSome())
}
