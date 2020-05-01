package main

import (
	"sync"
)

// FanIn merges n pipes to 1
func FanIn(pipes ...pipe) pipe {
	var wg sync.WaitGroup

	out := make(pipe)

	send := func(c pipe) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(pipes))

	for _, c := range pipes { // start a send goroutine for each input channel in pipes.
		go send(c)
	}

	go func() {
		wg.Wait() //wait until all goroutines are done and close output channel
		close(out)
	}()

	return out
}
