package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	input := Converter(numbers...)

	firstStage := NewTube(identity)(input)

	mediumStage := FanOut(firstStage, RoundRobin, NewTube(double), NewTube(square))

	merged := FanIn(mediumStage...)

	finalStage := NewTube(divideBy(2))(merged)

	result := Sink(finalStage)

	fmt.Println(result)
}
