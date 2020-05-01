package main

import "fmt"

func main() {
	input := []int{1, 3, 2}

	startStage := start(input...)

	identityStage := getStage(identity)

	squareStage := getStage(square)

	cubeStage := getStage(cube)

	startPipeline := cubeStage(squareStage(identityStage(startStage)))

	channelArray := split(startPipeline, 2, roundRobin)

	doubleStage := getStage(double)

	secondPipeline := doubleStage(channelArray[0])

	squareStage2 := getStage(square)

	thirdPipeline := squareStage2(channelArray[1])

	finalPipelineStart := merge(secondPipeline, thirdPipeline)

	identity2Stage := getStage(identity)

	finalPip := identity2Stage(finalPipelineStart)

	result := final(finalPip)

	fmt.Println(result)
}
