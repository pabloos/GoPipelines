package pipelines

type (
	// Flow represents the in/out coming inside Pipelines
	Flow   chan Element //we need a generic channel
	sender func(Flow, Flow, functor)
)

func send(outChan Flow, inChan Flow, mod functor) {
	// TODO MAIN: extract the receiver, and decouple from the modifier call
	for element := range inChan {
		var err error
		// TODO CANCELLATION: handle the error and do not send the result
		element.value, err = mod(element.value)

		if err != nil {
			cancelCh <- cancelSignal{}
		}

		select {
		case <-cancelCh:
			return
		case outChan <- element:
		}
	}
}

func closeFlow(flow Flow) {
	close(flow)
}

func sendAndClose(sender sender, closer func(Flow)) sender {
	return func(outChan Flow, inChan Flow, mod functor) {
		defer closer(outChan)

		sender(outChan, inChan, mod)
	}
}
