package main

import "fmt"

func main() {
	input := []int{1, 3, 2, 3}

	startStage := start(input...)

	firstPipeline := NewTube(identity)
	secondPipeline := NewTube(double)
	thirdPipeline := NewTube(square)

	firstPhase := firstPipeline(startStage)

	channelArray := FanOut(firstPhase, RoundRobin, secondPipeline, thirdPipeline)

	finalPipelineStart := FanIn(channelArray...)

	forthPipeline := NewTube(double)

	forthPhase := forthPipeline(finalPipelineStart)

	result := final(forthPhase)

	fmt.Println(result)
}
