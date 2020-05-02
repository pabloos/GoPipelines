package main

type (
	pipe   chan int //we need a generic channel
	stage  func(pipe) pipe
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
	return func(input pipe) pipe {
		output := make(pipe)

		go send(output, input, funct)

		return output
	}
}

func send(outChan pipe, inChan pipe, mod functor) {
	for n := range inChan {
		outChan <- mod(n)
	}

	close(outChan)
}

func createBufStage(bufLen int) func(functor) stage {
	return func(funct functor) stage {
		return func(input pipe) pipe {
			output := make(pipe, bufLen)

			go send(output, input, funct)

			return output
		}
	}
}
