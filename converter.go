package pipelines

import "sync/atomic"

// TODO set a wrapper object that encapsulates the deliever order
// TODO add the cancellation channel from here

// Converter passes the data from array to a flow
func Converter(numbers ...int) Flow {
	var orderNum uint64

	outputs := make(Flow, len(numbers))

	go func() {
		for _, number := range numbers {
			outputs <- Element{
				value:    number,
				orderNum: orderNum,
			}

			atomic.AddUint64(&orderNum, 1)
		}
		close(outputs)
	}()

	return outputs
}
