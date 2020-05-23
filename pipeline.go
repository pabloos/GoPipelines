package pipelines

type Pip Stage

// Pipeline returns a stage
func Pipeline(functors ...functor) Pip {
	stages := genStages(functors...)

	numStages := len(stages)

	return func(input Flow) Flow {
		acc := stages[0](input)

		if numStages == 1 { // if it's a stage-only flowline
			return acc
		}

		for i := 1; i < numStages; i++ {
			acc = stages[i](acc)
		}

		return acc
	}
}
