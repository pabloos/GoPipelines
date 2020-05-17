package pipelines

import (
	"sync"
)

// TODO implements a decorator that cares about the order of delivery order of inputs
// following the strategy defined by the user ([pre, in, post] - order => as it comes, tree-based, stacked)

// FanIn merges n flows to 1
func FanIn(flows ...Flow) Flow {
	var wg sync.WaitGroup

	out := make(Flow)

	// TODO insert cancellation login here
	send := func(c Flow) {
		send(out, c, identity)

		wg.Done()
	}

	wg.Add(len(flows))

	for _, c := range flows { // start a send goroutine for each input channel in flows.
		go send(c)
	}

	go func() {
		wg.Wait() //wait until all goroutines are done and close output channel
		close(out)
	}()

	return out
}
