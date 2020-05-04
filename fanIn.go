package main

import (
	"sync"
)

// FanIn merges n flows to 1
func FanIn(flows ...flow) flow {
	var wg sync.WaitGroup

	out := make(flow)

	send := func(c flow) {
		for n := range c {
			out <- n
		}
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
