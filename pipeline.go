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
