package pipelines

// FanOut distribute through
func FanOut(input Flow, scheduler Scheduler, stages ...Stage) (flows []Flow) {
	cs := make([]Flow, 0)

	for i := 0; i < len(stages); i++ {
		cs = append(cs, make(Flow))

		flows = append(flows, stages[i](cs[i]))
	}

	go schedule(scheduler)(input, cs)

	return flows
}
