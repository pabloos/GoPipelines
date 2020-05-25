package pipelines

import (
	"sync"
)

type Result struct {
	value []int
	sync.RWMutex
}

// Sink transforms the input channel values to an array
func Sink(inputs Flow) []int {
	var result Result

	result.value = make([]int, 0)

	var wg sync.WaitGroup

	wg.Add(1) //we only want to wait for just one of the goroutines

	// go multiplexerResultCancel(cancelCh, &wg, &result)
	go collectResults(inputs, &wg, &result)

	wg.Wait()

	return result.value
}

func collectResults(inputs Flow, wg *sync.WaitGroup, result *Result) {
	defer wg.Done()

	for input := range inputs {
		result.Lock()
		result.value = append(result.value, input)
		result.Unlock()
	}
}

// ! not efective -> delete
// func multiplexerResultCancel(cancel cancelChannel, wg *sync.WaitGroup, result *Result) {
// 	defer wg.Done()

// 	for {
// 		select {
// 		case <-cancel:
// 			fmt.Println("here")
// 			result.Lock()
// 			result.value = []int{} //empty results
// 			result.Unlock()

// 			return
// 		}
// 	}
// }
