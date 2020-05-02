package main

// Scheduler type represents a function that emits values through a an array of pipes
type Scheduler func(pipe, []pipe)

// split implements a round-roubin scheduled fan-out
func fanOut(ch pipe, scheduler Scheduler, tubes ...Tube) (output []pipe) {
	cs := make([]pipe, 0)

	for i := 0; i < len(tubes); i++ {
		cs = append(cs, make(pipe))

		output = append(output, tubes[i](cs[i]))
	}

	go scheduler(ch, cs)

	return output
}

// RoundRobin implements an scheduler tht continuously switch between the outputs: Round -> Robin -> Round -> Robin -> Round -> Robin ...
func RoundRobin(ch pipe, cs []pipe) {
	defer func(cs []pipe) { // at the end, close all incoming channels
		for _, c := range cs {
			close(c)
		}
	}(cs)

	for {
		for _, c := range cs {
			select {
			case val, ok := <-ch:
				if !ok {
					return
				}

				c <- val
			}
		}
	}
}
