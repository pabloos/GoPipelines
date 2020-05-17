package pipelines

type (
	// Flow represents the in/out coming inside Pipelines
	Flow   chan int //we need a generic channel
	sender func(Flow, Flow, functor)
)

// TODO insert cancellation logic here
func send(outChan Flow, inChan Flow, mod functor) {
	for n := range inChan {
		outChan <- mod(n)
	}
}

func closeFlow(flow Flow) {
	close(flow)
}

func sendAndClose(sender sender, closer func(Flow)) sender {
	return func(outChan Flow, inChan Flow, mod functor) {
		sender(outChan, inChan, mod)

		defer closer(outChan)
	}
}
