package pipelines

import (
	"sort"
	"sync"
)

// Sink transforms the input channel values to an array
func Sink(inputs Flow) []int {
	out := make([]int, 0)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		for input := range inputs {
			out = append(out, input.value)
		}
	}()

	wg.Wait()

	return out
}

// SinkWithOrder transforms the input channel values to an array with the order specified
func SinkWithOrder(inputs Flow, order Order) []int {
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

	sort.SliceStable(out, order(out))

	finalResults := getResults(out)

	return finalResults
}

func getResults(elements []Element) []int {
	newArr := make([]int, len(elements))

	for i, element := range elements {
		newArr[i] = element.value
	}

	return newArr
}
