package pipelines

// Scheduler type represents a function that emits values through a an array of flows
type Scheduler func(flow, []flow)

// Schedule prepares a scheduler to operate:
// - it injects the close channel action by a decorator pattern
func schedule(scheduler Scheduler) Scheduler {
	return func(ch flow, cs []flow) {
		defer closeChannels(cs)

		scheduler(ch, cs)
	}
}

func closeChannels(channels []flow) {
	for _, c := range channels {
		close(c)
	}
}

// RoundRobin implements an scheduler tht continuously switch between the outputs: Round -> Robin -> Round -> Robin -> Round -> Robin ...
func RoundRobin(ch flow, cs []flow) {
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
