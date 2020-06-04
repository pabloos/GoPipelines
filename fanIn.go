package pipelines

import (
	"context"
	"sync"
)

// FanIn merges n flows to 1
func FanIn(ctx context.Context, flows ...Flow) Flow {
	var wg sync.WaitGroup

	out := make(Flow)

	sendSync := func(c Flow) {
		defer wg.Done()

		//TODO MAIN: extract the receiver
		for n := range c {
			select {
			case out <- n:
			case <-ctx.Done():
				return
			}
		}
	}

	wg.Add(len(flows))

	for _, c := range flows { // start a send goroutine for each input channel in flows.
		go sendSync(c)
	}

	go func() {
		wg.Wait() //wait until all goroutines are done and close output channel
		defer close(out)
	}()

	return out
}
