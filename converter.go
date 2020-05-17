package pipelines

// Converter passes the data from array to a flow
func Converter(numbers ...int) flow {
	outputs := make(flow, len(numbers))

	go func() {
		for _, number := range numbers {
			outputs <- number
		}
		close(outputs)
	}()

	return outputs
}
