package main

// FanOut distribute through
func FanOut(input flow, scheduler Scheduler, tubes ...stage) (output []flow) {
	cs := make([]flow, 0)

	for i := 0; i < len(tubes); i++ {
		cs = append(cs, make(flow))

		output = append(output, tubes[i](cs[i]))
	}

	go schedule(scheduler)(input, cs)

	return output
}
