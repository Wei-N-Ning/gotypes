package io

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenEmptyReaderExpectEmptyLineIterator(t *testing.T) {
	reader := strings.NewReader("")
	iter := Lines(reader)
	assert.False(t, iter.Next().IsSome())
}

func TestExpectLines(t *testing.T) {
	reader := strings.NewReader(`there
is a silence

where hath been no sound`)
	xs := Lines(reader).Slice()
	assert.Equal(t, []Line{
		Line{Num: 1, Value: "there"},
		Line{Num: 2, Value: "is a silence"},
		Line{Num: 3, Value: ""},
		Line{Num: 4, Value: "where hath been no sound"},
	}, xs)
}
