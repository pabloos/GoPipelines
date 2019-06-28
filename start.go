package main

//the one who emits the data
func start(numbers ...int) pipe {
	outputs := make(chan int, len(numbers))

	go func() {
		for _, number := range numbers {
			outputs <- number
		}
		close(outputs)
	}()

	return outputs
}