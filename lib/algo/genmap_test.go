package algo

import (
	"fmt"
	"testing"
)

func chanFromSlice[T any](xs []T) <-chan T {
	ch := make(chan T)
	go func() {
		for _, x := range xs {
			ch <- x
		}
		close(ch)
	}()
	return ch
}

func TestGenericMapExpectError(t *testing.T) {
	for x := range chanFromSlice([]int{1, 2, 3}) {
		fmt.Println(x)
	}
}
