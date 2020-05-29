package pipelines

import (
	"sync"
)

// Sink transforms the input channel values to an array
func Sink(inputs Flow, order Order) []int {
	out := make([]Element, 0)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		for input := range inputs {
			out = append(out, input)
		}
	}()

	wg.Wait()

	orderedResults := order(out)

	finalResults := getResults(orderedResults)

	return finalResults
}

func getResults(elements []Element) []int {
	newArr := make([]int, len(elements))

	for i, element := range elements {
		newArr[i] = element.value
	}

	return newArr
}
