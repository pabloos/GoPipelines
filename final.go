package main

import "sync"

//end this receives all the data and presents it the way you want
func final(inputs pipe) []int {
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
