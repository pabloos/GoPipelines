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

	// cancellable := cancelWrp()
	// sender := cancellable(send)

	// TODO insert cancellation login here
	sendSync := func(c Flow) {
		//TODO MAIN: extract the receiver
		for n := range c {
			out <- n
		}

		defer wg.Done()
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
