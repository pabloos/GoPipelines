package main

// Tube contains the stages passed as functors.
// It equals to a simple stage, as it receives a pipe and returns other pipe.
// it needs to be assembled with a channel emiter and a channel receiver.
type Tube stage

// NewTube creates a tube
func NewTube(functors ...functor) Tube {
	return func(input pipe) pipe {
		stages := genStages(functors...)

		numStages := len(stages)

		acc := stages[0](input)

		if numStages == 1 { // if it's a stage-only pipeline
			return acc
		}

		for i := 1; i < numStages; i++ {
			acc = stages[i](acc)
		}

		return acc
	}
}
