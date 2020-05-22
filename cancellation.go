package pipelines

type cancelSignal struct{}
type cancelChannel chan cancelSignal

var cancelCh cancelChannel

func init() {
	cancelCh = make(cancelChannel)
}
