package main

// Converter passes the data from array to a pipe
func Converter(numbers ...int) pipe {
	outputs := make(pipe, len(numbers))

	go func() {
		for _, number := range numbers {
			outputs <- number
		}
		close(outputs)
	}()

	return outputs
}
