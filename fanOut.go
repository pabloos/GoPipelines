package pipelines

import "context"

// FanOut distribute through
func FanOut(ctx context.Context, input Flow, scheduler Scheduler, pipelines ...Stage) (flows []Flow) {
	cs := make([]Flow, 0)

	for i := 0; i < len(pipelines); i++ {
		cs = append(cs, make(Flow))

		flows = append(flows, pipelines[i](cs[i]))
	}

	go schedule(scheduler)(ctx, input, cs)

	return flows
}
