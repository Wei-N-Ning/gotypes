package iterator

import (
	"testing"
)

func TestFromMap(t *testing.T) {
	var iter Iterator[Pair[string, int]] = FromMap(map[string]int{
		"map size":  1,
		"cre count": 10,
		"code":      0xe1d1,
	})
	iter.Slice()
}
