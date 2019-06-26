package main

import (
	"fmt"
	"sync"
)

//the one who emits the data
func source(numbers ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

/*
	There are the stages. Each one represents the phases of the workchain
	As you can see ,all ones has the same input and output type: channels
*/

func firstStage(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n
		}
		close(out)
	}()
	return out
}

func secondStage(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func thirdStage(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n + n
		}
		close(out)
	}()
	return out
}

//end this receives all the data and presents it the way you want
func end(in <-chan int) []int {
	out := make([]int, 0)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		for n := range in {
			out = append(out, n)
		}
		wg.Done()
	}()

	wg.Wait()

	return out
}

func main() {
	// Set up the pipeline and consume the output.
	for _, n := range end(thirdStage(secondStage(firstStage(source(2, 3, 2, 34))))) {
		fmt.Println(n)
	}
}
