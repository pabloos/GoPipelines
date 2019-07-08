package main

import (
	"fmt"
	"sync"
)

//the function who send data to the first stage
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
	The stages. Each one represents the phases of the workchain
	As you can see, all of them has the same input and output type: channels
*/
func firstStage(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n + 2 //we transform the data and send to the channel
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

//end receives all the data and presents it the way you want
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
	//we set up the pipeline (this requires to nest every stage into the next one) -> not easy to read
	for _, n := range end(thirdStage(secondStage(firstStage(source(2, 3, 2, 34))))) {
		fmt.Println(n)
	}
}
