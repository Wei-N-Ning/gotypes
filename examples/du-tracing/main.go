package main

import (
	"fmt"
	"github.com/Wei-N-Ning/gotypes/pkg/iterator"
	"github.com/Wei-N-Ning/gotypes/pkg/iterator/fs"
	"io"
	"os"
	"runtime/trace"
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

func getSizeSlow(item fs.Item) int64 {
	if item.DirEntry.IsDir() {
		return 0
	} else {
		r, err := os.Open(item.Path)
		if err != nil {
			return 0
		}
		bs := make([]byte, 128)
		totalRead := 0
		for {
			numRead, err := r.Read(bs)
			if err == io.EOF {
				break
			}
			totalRead += numRead
		}
		return int64(totalRead)
	}
}

func main() {
	args := os.Args[1:]
	dirPath := "."
	if len(args) > 0 {
		dirPath = args[0]
	}

	var x int64

	func() {
		err := trace.Start(os.Stderr)
		if err != nil {
			panic(err)
		}
		defer trace.Stop()

		// ----------- round 1: fast IO (zero IO) --------------

		// serial-for-each
		//fs.DirIter(dirPath).ForEach(func(item fs.Item) {
		//	x += getSize(item)
		//})

		// serial-map-reduce
		//x = iterator.MapReduce(fs.DirIter(dirPath), 0, getSize, addTwo)

		// Parallel-map
		//x = iterator.ParMap(fs.DirIter(dirPath), func(item fs.Item) int64 {
		//	return getSize(item)
		//}).Reduce(0, addTwo)

		// Parallel-unordered-map
		//x = iterator.ParMapUnord(fs.DirIter(dirPath), func(item fs.Item) int64 {
		//	return getSize(item)
		//}).Reduce(0, addTwo)

		// Parallel-map-reduce
		//x = iterator.ParMapReduce(fs.DirIter(dirPath), 0, getSize, addTwo)

		// ----------- round 2: slow IO --------------

		// serial-for-each
		//fs.DirIter(dirPath).ForEach(func(item fs.Item) {
		//	x += getSizeSlow(item)
		//})

		// serial-map-reduce
		//x = iterator.MapReduce(fs.DirIter(dirPath), 0, getSizeSlow, addTwo)

		// Parallel-map
		//x = iterator.ParMap(fs.DirIter(dirPath), func(item fs.Item) int64 {
		//	return getSizeSlow(item)
		//}).Reduce(0, addTwo)

		// Parallel-unordered-map
		//x = iterator.ParMapUnord(fs.DirIter(dirPath), func(item fs.Item) int64 {
		//	return getSizeSlow(item)
		//}).Reduce(0, addTwo)

		// Parallel-map-reduce
		x = iterator.ParMapReduce(fs.DirIter(dirPath), 0, getSizeSlow, addTwo)

	}()

	fmt.Println(x/(1024*1024), "M")

}
