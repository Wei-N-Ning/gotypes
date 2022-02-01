package iterator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
	"time"
)

func fib(x int) int {
	if x < 1 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func timeThis(f func()) int64 {
	start := time.Now()
	f()
	return time.Since(start).Nanoseconds()
}

func TestParMapExpectLaziness(t *testing.T) {
	const WORKLOAD = 999
	// all the computation here would take forever, but it is not evaluated at all
	iter := ParMapUnord(Repeat(WORKLOAD).Take(999), func(x int) string {
		fib(x)
		return fmt.Sprintf("%02d", x)
	})
	iter = Once("a")
	assert.Equal(t, 1, iter.Count())
}

func TestParMapCompareRuntime(t *testing.T) {
	const WORKLOAD = 30

	single := func() {
		fib(WORKLOAD)
	}
	singleTime := timeThis(single)

	serial := func() {
		Map(Repeat(WORKLOAD).Take(runtime.NumCPU()), func(x int) string {
			fib(x)
			return fmt.Sprintf("%02d", x)
		}).Count()
	}
	serialTime := timeThis(serial)

	parallel := func() {
		ParMap(Repeat(WORKLOAD).Take(runtime.NumCPU()), func(x int) string {
			fib(x)
			return fmt.Sprintf("%02d", x)
		}).Count()
	}
	parallelTime := timeThis(parallel)

	parallelUnordered := func() {
		ParMapUnord(Repeat(WORKLOAD).Take(runtime.NumCPU()), func(x int) string {
			fib(x)
			return fmt.Sprintf("%02d", x)
		}).Count()
	}
	parallelUnorderedTime := timeThis(parallelUnordered)

	fmt.Println("single", singleTime)
	fmt.Println("serial", serialTime)
	fmt.Println("order ", parallelTime)
	fmt.Println("unord ", parallelUnorderedTime)
}
