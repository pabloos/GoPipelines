package pipelines

import (
	"context"
)

type (
	cancelSignal struct{}
	cancelChan   chan cancelSignal
)

// var pipFactory = cancelWrp()

func cancelWrp() func(sender) sender { //cancelInstance
	_, cancel := context.WithCancel(context.Background())

	return func(sender sender) sender { //withCancel
		return func(out Flow, in Flow, mod functor) error { // go sender
			err := sender(out, in, mod)

			if err != nil {
				cancel()

				return err
			}

			// return err

			return nil
		}
	}
}
