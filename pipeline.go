package pipelines

import "context"

// Pipeline returns a Stage and a CancelFunc to cancel his context
func Pipeline(ctx context.Context, functors ...functor) Stage {
	errChs := make([]errorChannel, len(functors))

	for i := range errChs {
		errChs[i] = make(errorChannel, 1)
	}

	stages := genStages(ctx, errChs, functors...)

	numStages := len(stages)

	return func(input Flow) Flow {
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
