package iterator

import (
	"fmt"
	"testing"
)
import . "go-types-nw/lib/algo/option"

func TestFibonacciNumberIterator(t *testing.T) {
	type Item struct {
		first  int
		second int
	}
	init := Some[Item](Item{first: 0, second: 1})
	f := func(x Item) Option[Item] {
		first := x.second
		second := x.first + x.second
		return Some[Item](Item{first: first, second: second})
	}
	iter := Successor(init, f).Take(10)
	iter.ForEach(func(x Item) {
		fmt.Printf("fib: %d ", x.first)
	})
	fmt.Println()
}
