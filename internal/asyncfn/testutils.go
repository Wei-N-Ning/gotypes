package asyncfn

func stubSignals(sigs ...Signal) <-chan Signal {
	ch := make(chan Signal, len(sigs))
	go func() {
		defer close(ch)
		for _, sig := range sigs {
			ch <- sig
		}
	}()
	return ch
}
