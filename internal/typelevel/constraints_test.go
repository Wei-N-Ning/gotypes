package typelevel

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type BigInteger uint64

type QuantizableI interface {
	int | int32 | uint32 | int64 | uint64
}

type QuantizableII interface {
	int | int32 | uint32 | int64 | BigInteger
}

func quantizationI[T QuantizableI](value T) string {
	return fmt.Sprintf("%d", value)
}

func quantizationII[T QuantizableII](value T) string {
	return fmt.Sprintf("%d", value)
}

func TestQuantizationI(t *testing.T) {
	assert.Equal(t, "10", quantizationI(10)) // int
	assert.Equal(t, "10", quantizationI(int32(10)))
	assert.Equal(t, "10", quantizationI(uint32(10)))
	assert.Equal(t, "10", quantizationI(int64(10)))
	assert.Equal(t, "10", quantizationI(uint64(10)))
}

func TestQuantizationII(t *testing.T) {
	assert.Equal(t, "10", quantizationII(10)) // int
	assert.Equal(t, "10", quantizationII(int32(10)))
	assert.Equal(t, "10", quantizationII(uint32(10)))
	assert.Equal(t, "10", quantizationII(int64(10)))

	// compile error! uint64 does not implement QuantizableII
	//assert.Equal(t, "10", quantizationII(uint64(10)))
	assert.Equal(t, "10", quantizationII(BigInteger(10)))
}
