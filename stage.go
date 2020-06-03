package pipelines

import "context"

type (
	// Stage represents the different phases of a Pipeline, which it's a Stage itself
	Stage func(Flow) Flow

	stages []Stage
)

func genStages(ctx context.Context, errChs []errorChannel, functors ...functor) stages {
	stages := make(stages, 0)

	for i, functor := range functors {
		stages = append(stages, getStage(ctx, errChs[i], functor))
	}

	return stages
}

func getStage(ctx context.Context, errCh errorChannel, funct functor) Stage {
	sender := sendAndClose(send, closeFlow)

	return func(input Flow) Flow {
		output := make(Flow)

		go sender(ctx, errCh, output, input, funct)

		return output
	}
}
