package main

import (
	"fmt"
	"sync"
)

type stage func(in <-chan int, t transformer) <-chan int

type transformer func(int) int

func sender(outChan chan int, inChan <-chan int, t transformer) {
	for n := range inChan {
		outChan <- t(n)
	}
	close(outChan)
}

func square(number int) int {
	return number * number
}

func add2(number int) int {
	return number + 2
}

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

func firstStage(in <-chan int, t transformer) <-chan int {
	out := make(chan int)
	go sender(out, in, t)
	return out
}

func secondStage(in <-chan int, t transformer) <-chan int {
	out := make(chan int)
	go sender(out, in, t)
	return out
}

func thirdStage(in <-chan int, t transformer) <-chan int {
	out := make(chan int)
	go sender(out, in, t)
	return out
}

func main() {
	for _, n := range end(thirdStage(secondStage(firstStage(source(1, 2), add2), square), add2)) {
		fmt.Println(n)
	}
}
