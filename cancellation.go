package pipelines

// var pipFactory = cancelWrp()

// func cancelWrp() func(sender) sender {
// 	return func(sender sender) sender {
// 		return func(out Flow, in Flow, mod functor) error { // go sender
// 			err := sender(out, in, mod)

// 			if err != nil {
// 				return err
// 			}

// 			return nil
// 		}
// 	}
// }

type cancelSignal struct{}
type cancelChannel chan cancelSignal

var cancelCh = make(cancelChannel)

// func init() {
// 	cancelCh = make(cancelChannel)
// }
