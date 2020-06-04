package pipelines

import (
	"context"
	"fmt"
)

type (
	// Flow represents the in/out coming inside Pipelines
	Flow         chan Element //we need a generic channel
	errorChannel chan error

	sender func(context.Context, errorChannel, Flow, Flow, functor)
)

// TODO MAIN: extract the receiver, and decouple from the modifier call
func send(ctx context.Context, errChan errorChannel, outChan Flow, inChan Flow, mod functor) {
	// !!!!
	defer close(errChan)

	for element := range inChan {
		var err error

		element.value, err = mod(element.value)

		if err != nil {
			fmt.Println("err check")
			errChan <- err

			return
		}

		select {
		case outChan <- element:
		case <-ctx.Done():
			fmt.Println("Done applyed")
			return
		}
	}
}

func closeFlow(flow Flow) {
	close(flow)
	fmt.Println("channel clossed")
}

func sendAndClose(sender sender, closer func(Flow)) sender {
	return func(ctx context.Context, errChan errorChannel, outChan Flow, inChan Flow, mod functor) {
		defer closer(outChan)

		sender(ctx, errChan, outChan, inChan, mod)
	}
}
