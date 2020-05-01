package main

type (
	functor func(int) int
	source  func(...int) pipe
	end     func(pipe) []int

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

func NewFinalPipeline(source source, end end, functors ...functor) func(...int) []int {
	stages := genStages(functors...)

	return func(input ...int) []int {
		lastStageIndex := len(stages) - 1

		start := source(input...)

		result := stages[0](start)

		for i := 1; i < lastStageIndex; i++ {
			result = stages[i](result)
		}

		if lastStageIndex == 0 { // if it's a single stage we don't want to remake the last stage (as it's also the first again)
			return end(result)
		}

		return end(stages[lastStageIndex](result))
	}
}
