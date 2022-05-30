package asyncfn

import "context"

// trySelectOne is a non-blocking call.
// it polls N read-only Signal channels and returns the first yielded (Signal, true) to the caller;
// ("true" indicates that there is one channel that yields - but the caller doesn't know which one)
// if no channel yields, return a tuple of (InvalidSignal, false)
// ("false" indicates that no channel has yielded)
func trySelectOne(channels ...<-chan Signal) (Signal, bool) {
	for _, ch := range channels {
		select {
		case value, isChanOpen := <-ch:
			if isChanOpen {
				return value, true
			}
		default:
		}
	}
	return InvalidSignal, false
}

// selectOne is a blocking call.
// it polls N read-only Signal channels and returns the first yielded Signal to the caller;
// it will poll indefinitely until:
// - case 1, one of the channel yields,
// - case 2, the parent context issues cancellation, or
// - case 3, all the channels are closed
// in case 2 and 3, it returns (InvalidSignal, false)
// the second return value serves the same purposes as that in the non-blocking version, trySelectOne
// it signals the caller whether any channel has yielded
func selectOne(ctx context.Context, channels ...<-chan Signal) (Signal, bool) {
	numOpen := len(channels)
	for {
		select {

		case <-ctx.Done():
			// case 2, the parent context issues cancellation
			return InvalidSignal, false

		default:
			for idx := range channels {
				if channels[idx] == nil {
					continue
				}
				select {
				case value, isChanOpen := <-channels[idx]:
					if !isChanOpen {
						channels[idx] = nil
						numOpen--
					} else {
						// case 1, one of the channel yields
						return value, true
					}
				default:
				}
			}
		}

		if numOpen == 0 {
			// case 3, all the channels are closed
			return InvalidSignal, false
		}
	}
}
