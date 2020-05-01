package main

type (
	functor func(int) int
	pipe    chan int //we need a generic channel
	source  func(...int) pipe
	end     func(pipe) []int
	stage   func(pipe) pipe
	stages  []stage

	//Pipeline contains the start and the end of himself and all his stages
	Pipeline struct {
		source
		stages
		end
	}
	// Tube contains the stages passed as functors.
	// It equals to a simple stage, as it receives a pipe and returns other pipe.
	// it needs to be assembled with a channel emiter and a channel receiver.
	Tube stage
)

//Exec actions the pipeline for a certain input
func (pip *Pipeline) Exec(input ...int) []int {
	lastStageIndex := len(pip.stages) - 1

	start := pip.source(input...)

	result := pip.stages[0](start)

	for i := 1; i < lastStageIndex; i++ {
		result = pip.stages[i](result)
	}

	if lastStageIndex == 0 { // if it's a single stage we don't want to remake the last stage (as it's also the first again)
		return pip.end(result)
	}

	return pip.end(pip.stages[lastStageIndex](result))
}

//NewPipeline returns a Pipeline
func NewPipeline(source source, end end, functors ...functor) *Pipeline {
	return &Pipeline{
		source,
		genStages(functors...),
		end,
	}
}

// NewTube creates a tube
func NewTube(functors ...functor) Tube {
	return func(input pipe) pipe {
		stages := genStages(functors...)

		numStages := len(stages)

		acc := stages[0](input)

		if numStages == 1 { // if it's a stage-only pipeline
			return acc
		}

		for i := 1; i < numStages; i++ { //make a reduce with a slice of [1:]
			acc = stages[i](acc)
		}

		return acc
	}
}

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
