package main

import (
	"fmt"
	"github.com/Wei-N-Ning/gotypes/lib/algo/iterator"
	"github.com/Wei-N-Ning/gotypes/lib/algo/iterator/fs"
	"os"
)

func addTwo(lhs int64, rhs int64) int64 {
	return lhs + rhs
}

func getSize(item fs.Item) int64 {
	if info, err := item.DirEntry.Info(); err != nil {
		return 0
	} else {
		return info.Size()
	}
}

func main() {
	args := os.Args[1:]
	dirPath := "."
	if len(args) > 0 {
		dirPath = args[0]
	}
	// the ParMapUnordered version performs poorly (likely due to the overhead of channel)
	// the computation is too trivial to benefit from parallelism
	x := iterator.ParMapReduce(fs.DirIter(dirPath), 0, getSize, addTwo)
	fmt.Println(x/(1024*1024), "M")
}
