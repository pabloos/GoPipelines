package main

import "fmt"

func main() {
	input := []int{1, 3, 2, 3}

	startStage := start(input...)

	firstPipeline := NewTube(identity)
	secondPipeline := NewTube(double)
	thirdPipeline := NewTube(square)

	firstPhase := firstPipeline(startStage)

	channelArray := fanOut(firstPhase, Schedule(RoundRobin), secondPipeline, thirdPipeline)

	finalPipelineStart := merge(channelArray...)

	forthPipeline := NewTube(double)

	forthPhase := forthPipeline(finalPipelineStart)

	result := final(forthPhase)

	fmt.Println(result)
}
