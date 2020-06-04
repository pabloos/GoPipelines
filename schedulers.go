package pipelines

import "context"

// Scheduler type represents a function that emits values through a an array of Flows
type Scheduler func(context.Context, Flow, []Flow)

// Schedule prepares a scheduler to operate:
// - it injects the close channel action by a decorator pattern
func schedule(scheduler Scheduler) Scheduler {
	return func(ctx context.Context, ch Flow, cs []Flow) {
		defer closeChannels(cs)

		scheduler(ctx, ch, cs)
	}
}

func closeChannels(channels []Flow) {
	for _, c := range channels {
		close(c)
	}
}

// RoundRobin implements an scheduler that continuously switches between the outputs
func RoundRobin(ctx context.Context, ch Flow, cs []Flow) {
	for {
		for _, c := range cs {
			select {
			case val, ok := <-ch:
				if !ok {
					return
				}

				c <- val
			case <-ctx.Done():
				return
			}
		}
	}
}
