package main

import "fmt"

func main() {
	input := []int{1, 3, 2}

	startStage := start(input...)

	firstPipeline := NewTube(identity, square, cube)

	firstPhase := firstPipeline(startStage)

	channelArray := split(firstPhase, 2, roundRobin)

	secondPipeline := NewTube(double)

	secondPhase := secondPipeline(channelArray[0])

	thirdPipeline := NewTube(square)

	thirdPhase := thirdPipeline(channelArray[1])

	finalPipelineStart := merge(secondPhase, thirdPhase)

	forthPipeline := NewTube(identity)

	forthPhase := forthPipeline(finalPipelineStart)

	result := final(forthPhase)

	fmt.Println(result)
}
