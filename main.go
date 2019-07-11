package main

import (
	"fmt"
	"sync"
)

type (
	pipe    <-chan int
	stage   func(in pipe, t functor) pipe
	functor func(int) int
)

func sender(outChan chan int, inChan pipe, t functor) {
	for n := range inChan {
		outChan <- t(n)
	}
	close(outChan)
}

//the one who emits the data
func source(numbers ...int) pipe {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

//end this receives all the data and presents it the way you want
func end(in pipe) []int {
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

func firstStage(in pipe, t functor) pipe {
	out := make(chan int)
	go sender(out, in, t)
	return out
}

func secondStage(in pipe, t functor) pipe {
	out := make(chan int)
	go sender(out, in, t)
	return out
}

func thirdStage(in pipe, t functor) pipe {
	out := make(chan int)
	go sender(out, in, t)
	return out
}

func main() {
	for _, n := range end(thirdStage(secondStage(firstStage(source(1, 2), add2), square), add2)) {
		fmt.Println(n)
	}
}
