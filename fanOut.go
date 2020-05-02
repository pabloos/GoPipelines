package main

// fanOut distribute through
func fanOut(input pipe, scheduler Scheduler, tubes ...Tube) (output []pipe) {
	cs := make([]pipe, 0)

	for i := 0; i < len(tubes); i++ {
		cs = append(cs, make(pipe))

		output = append(output, tubes[i](cs[i]))
	}

	go schedule(scheduler)(input, cs)

	return output
}
