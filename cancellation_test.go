package pipelines

import (
	"reflect"
	"testing"
)

func TestWOCanceWOSinkl(t *testing.T) {
	numbers := []int{1, 2, 3}

	input := Converter(numbers...)

	firstStage := NewPipeline(multBy(2))(input)

	result := <-firstStage

	if !reflect.DeepEqual(result, 2) {
		t.Errorf("result was: %v", result)
	}

	result = <-firstStage

	if !reflect.DeepEqual(result, 4) {
		t.Errorf("result was: %v", result)
	}
}

func TestSimplePipelineCancel(t *testing.T) {
	numbers := []int{1, 2, 3}

	input := Converter(numbers...)

	firstStage := NewPipeline(multBy(2), cancel, multBy(2))(input)

	secondStage := NewPipeline(identity)(firstStage)

	result := Sink(secondStage)

	// TODO: FIX THIS
	// ! UNDETERMINISTC
	// * some times this produces a negative wg count on the sink phase
	if !reflect.DeepEqual(result, []int{}) {
		t.Errorf("result was: %v", result)
	}
}
