package main

import (
	"fmt"
	"github.com/Wei-N-Ning/gotypes/pkg/iterator"
	"github.com/Wei-N-Ning/gotypes/pkg/iterator/fs"
	"github.com/Wei-N-Ning/gotypes/pkg/iterator/io"
	"os"
	"strings"
)

// This mini ack-clone is to demonstrate the performance advantage of ParMapUnordered.
// There are a few things worth noting:
// - this demo does not use Regex (cheating)
// - the task of searching for pattern line-by-line from all the source files is IO-bound
// - in production, ParMap (the ordered) version is probably more reasonable
//   (hyperfine shows ParMap version performs equally well)
// - the demo does not use any terminal output control like the real ack does (cheating)

/*

Benchmark 1: ./main ./lib func
  Time (mean ± σ):       2.5 ms ±   0.6 ms    [User: 2.1 ms, System: 3.1 ms]
  Range (min … max):     1.8 ms …   6.2 ms    554 runs

  Warning: Command took less than 5 ms to complete. Results might be inaccurate.
  Warning: Statistical outliers were detected. Consider re-running this benchmark on a quiet PC without any interferences from other programs. It might help to use the '--warmup' or '--prepare' options.

Benchmark 2: ack --type=go func ./lib
  Time (mean ± σ):      35.4 ms ±   1.8 ms    [User: 26.3 ms, System: 4.9 ms]
  Range (min … max):    33.3 ms …  44.0 ms    76 runs

  Warning: Statistical outliers were detected. Consider re-running this benchmark on a quiet PC without any interferences from other programs. It might help to use the '--warmup' or '--prepare' options.

Summary
  './main ./lib func' ran
   13.99 ± 3.38 times faster than 'ack --type=go func ./lib'

*/

func main() {
	args := os.Args[1:]
	dirPath := "."
	pattern := ""
	if len(args) > 1 {
		dirPath = args[0]
		pattern = args[1]
	} else {
		os.Exit(1)
	}
	x := iterator.ParMap(fs.DirIter(dirPath), func(item fs.Item) bool {
		if item.DirEntry.IsDir() && !strings.HasSuffix(item.Path, ".go") {
			return false
		}
		if strings.Contains(item.Path, ".git") {
			return false
		}
		if strings.Contains(item.Path, ".idea") {
			return false
		}
		f, err := os.Open(item.Path)
		if err != nil {
			return false
		}
		io.Lines(f).ForEach(func(line io.Line) {
			if strings.Contains(line.Value, pattern) {
				fmt.Println(line.ToString())
			}
		})
		return true
	})
	x.Count()
}
