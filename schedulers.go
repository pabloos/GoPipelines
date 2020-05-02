package main

// Scheduler type represents a function that emits values through a an array of pipes
type Scheduler func(pipe, []pipe)

// Schedule prepares a scheduler to operate:
// - it injects the close channel action by a decorator pattern
func Schedule(scheduler Scheduler) Scheduler {
	return func(ch pipe, cs []pipe) {
		defer closeChannels(cs)

		scheduler(ch, cs)
	}
}

func closeChannels(channels []pipe) {
	for _, c := range channels {
		close(c)
	}
}

// RoundRobin implements an scheduler tht continuously switch between the outputs: Round -> Robin -> Round -> Robin -> Round -> Robin ...
func RoundRobin(ch pipe, cs []pipe) {
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
