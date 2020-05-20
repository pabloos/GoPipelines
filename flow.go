package pipelines

type (
	// Flow represents the in/out coming inside Pipelines
	Flow   chan int //we need a generic channel
	sender func(Flow, Flow, functor) error
)

// TODO insert cancellation logic here
func send(outChan Flow, inChan Flow, mod functor) error {
	// TODO MAIN: extract the receiver, and decouple from the modifier call
	for n := range inChan {
		n, err := mod(n)

		if err != nil {
			return err
		}

		outChan <- n
	}

	return nil
}

func closeFlow(flow Flow) {
	close(flow)
}

func sendAndClose(sender sender, closer func(Flow)) sender {
	return func(outChan Flow, inChan Flow, mod functor) error {
		defer closer(outChan)

		return sender(outChan, inChan, mod)
	}
}
