package pipelines

// Scheduler type represents a function that emits values through a an array of Flows
type Scheduler func(Flow, []Flow)

// Schedule prepares a scheduler to operate:
// - it injects the close channel action by a decorator pattern
func schedule(scheduler Scheduler) Scheduler {
	return func(ch Flow, cs []Flow) {
		defer closeChannels(cs)

		scheduler(ch, cs)
	}
}

func closeChannels(channels []Flow) {
	for _, c := range channels {
		close(c)
	}
}

// RoundRobin implements an scheduler that continuously switches between the outputs
// TODO insert cancellation logic here
// maybe it's enough with an implicit cancelation (as it's implemented now)
func RoundRobin(ch Flow, cs []Flow) {
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
