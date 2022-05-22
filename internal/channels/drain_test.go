package channels

import (
	"testing"

	"github.com/Wei-N-Ning/gotypes/pkg/iterator"
	"github.com/stretchr/testify/assert"
)

// drain a channel:
// read all the values from a channel till 1) the read is blocked 2) the channel is closed
// 1) the read is blocked; return None, true
// 2) the channel is closed; return None, false - the boolean is to simulate the behaviour of the normal read operator

func drain[T any](inCh <-chan T) (int, bool) {
	numElems := 0
	for {
		select {
		case _, isChanOpen := <-inCh:
			if !isChanOpen {
				return numElems, false
			}
			numElems++
		default:
			return numElems, true
		}
	}
}

func fromRange(from int, to int, buffer int) <-chan int {
	outCh := make(chan int, buffer)
	go func() {
		defer close(outCh)
		iterator.Range(from, to).ForEach(func(x int) { outCh <- x })
	}()
	return outCh
}

func TestDrainOpenChannel(t *testing.T) {
	t.Run("drain an unbuffered open channel, expect return when read is blocked", func(t *testing.T) {
		for _, n := range []int{10, 100, 1000, 2000, 3000} {
			from := 1
			to := n
			outCh := make(chan int, to-from+1)
			barrier := make(chan struct{})
			go func() {
				iterator.Range(from, to).ForEach(func(x int) { outCh <- x })
				barrier <- struct{}{}
			}()
			<-barrier
			numElems, isChanOpen := drain[int](outCh)
			assert.Equal(t, to-from, numElems)
			assert.True(t, isChanOpen)
		}
	})
}
