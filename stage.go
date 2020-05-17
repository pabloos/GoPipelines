package pipelines

type (
	// Stage represents the different phases of a Pipeline, which it's a Stage itself
	Stage func(Flow) Flow

	stages []Stage
)

func genStages(functors ...functor) stages {
	stages := make(stages, 0)

	for _, functor := range functors {
		stages = append(stages, getStage(functor))
	}

	return stages
}

func getStage(funct functor) Stage {
	sender := sendAndClose(send, closeFlow)

	return func(input Flow) Flow {
		output := make(Flow)

		go sender(output, input, funct)

		return output
	}
}

// TODO insert cancellation logic here
func createBufStage(bufLen int) func(functor) Stage {
	sender := sendAndClose(send, closeFlow)

	return func(funct functor) Stage {
		return func(input Flow) Flow {
			output := make(Flow, bufLen)

			go sender(output, input, funct)

			return output
		}
	}
}
