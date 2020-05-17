package pipelines

import "sync"

// Sink transforms the input channel values to an array
func Sink(inputs Flow) []int {
	out := make([]int, 0)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		for input := range inputs {
			out = append(out, input)
		}
		wg.Done()
	}()

	wg.Wait()

	return out
}
