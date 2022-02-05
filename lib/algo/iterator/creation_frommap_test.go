package iterator

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFromMap(t *testing.T) {
	var iter Iterator[Pair[string, int]] = FromMap(map[string]int{
		"map size":  1,
		"cre count": 10,
		"code":      0xe1d1,
	})
	num := iter.CountIf(func(p Pair[string, int]) bool {
		return strings.HasPrefix(p.First, "map")
	})
	assert.Equal(t, 1, num)
}
