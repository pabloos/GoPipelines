package main

type (
	flow   chan int //we need a generic channel
	stage  func(flow) flow
	stages []stage
)

func genStages(functors ...functor) stages {
	stages := make(stages, 0)

	for _, functor := range functors {
		stages = append(stages, getStage(functor))
	}

	return stages
}

func getStage(funct functor) stage {
	return func(input flow) flow {
		output := make(flow)

		go send(output, input, funct)

		return output
	}
}

func send(outChan flow, inChan flow, mod functor) {
	for n := range inChan {
		outChan <- mod(n)
	}

	close(outChan)
}

func createBufStage(bufLen int) func(functor) stage {
	return func(funct functor) stage {
		return func(input flow) flow {
			output := make(flow, bufLen)

			go send(output, input, funct)

			return output
		}
	}
}
