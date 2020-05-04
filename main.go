package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	input := Converter(numbers...)

	firstStage := NewPipeline(identity)(input)

	mediumStage := FanOut(firstStage, RoundRobin, NewPipeline(double), NewPipeline(square))

	merged := FanIn(mediumStage...)

	finalStage := NewPipeline(divideBy(2))(merged)

	result := Sink(finalStage)

	fmt.Println(result)
}
