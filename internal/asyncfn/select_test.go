package asyncfn

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// makeChannels returns a slice of signal channels that yield after the given timeouts
func makeChannels(timeouts []int) []<-chan Signal {
	var chs []<-chan Signal
	for idx, timeout := range timeouts {
		ch := make(chan Signal, 1)
		go func(idx int64, timeout int) {
			defer close(ch)
			if timeout > 0 {
				time.Sleep(time.Duration(timeout) * time.Millisecond)
			}
			ch <- Signal{strconv.FormatInt(idx, 16)}
		}(int64(idx), timeout)
		chs = append(chs, ch)
	}
	return chs
}

func TestTrySelectOne(t *testing.T) {
	t.Run("should yield", func(t *testing.T) {
		chs := makeChannels([]int{10, 100, 23, 0, 3})
		<-chs[0]
		_, yielded := trySelectOne(chs...)
		require.True(t, yielded)
	})
	t.Run("should not yield", func(t *testing.T) {
		chs := makeChannels([]int{1000, 100, 230, 110, 300})
		_, yielded := trySelectOne(chs...)
		require.False(t, yielded)
	})
	t.Run("empty signal slice, should not yield", func(t *testing.T) {
		chs := makeChannels([]int{})
		_, yielded := trySelectOne(chs...)
		require.False(t, yielded)
	})
}

func TestSelectOnBlocking(t *testing.T) {
	t.Run("should block and yield", func(t *testing.T) {
		chs := makeChannels([]int{10, 100, 23, 0, 3})
		<-chs[0]
		_, yielded := selectOne(context.Background(), chs...)
		require.True(t, yielded)
	})
	t.Run("parent context cancel (timeout), should not yield", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 40*time.Millisecond)
		defer cancel()
		chs := makeChannels([]int{1000, 400, 230, 510, 300})
		_, yielded := selectOne(ctx, chs...)
		require.False(t, yielded)
	})
	t.Run("empty signal slice, should not yield", func(t *testing.T) {
		chs := makeChannels([]int{})
		_, yielded := selectOne(context.Background(), chs...)
		require.False(t, yielded)
	})

}
