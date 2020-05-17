package pipelines

// TODO set a wrapper object that encapsulates the deliever order
// TODO add the cancellation channel from here

// Converter passes the data from array to a flow
func Converter(numbers ...int) Flow {
	outputs := make(Flow, len(numbers))

	go func() {
		for _, number := range numbers {
			outputs <- number
		}
		close(outputs)
	}()

	return outputs
}
