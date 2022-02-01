package main

import (
	"fmt"
	"go-types-nw/lib/algo/iterator"
	"os"
)

func main() {
	args := os.Args[1:]
	dirPath := "."
	if len(args) > 0 {
		dirPath = args[0]
	}
	// the ParMapUnordered version performs poorly (likely due to the overhead of channel)
	// the computation is too trivial to benefit from parallelism
	iter := iterator.Map(iterator.DirIter(dirPath), func(item iterator.Item) int64 {
		if info, err := item.DirEntry.Info(); err != nil {
			return 0
		} else {
			return info.Size()
		}
	})
	x := iterator.Fold(iter, int64(0), func(acc int64, elem int64) int64 { return acc + elem })
	fmt.Println(x/(1024*1024), "M")
}
