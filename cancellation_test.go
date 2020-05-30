package pipelines

import (
	"reflect"
	"testing"
)

func TestWOCanceWOSinkl(t *testing.T) {
	numbers := []int{1, 2, 3}

	input := Converter(numbers...)

	firstStage := Pipeline(multBy(2))(input)

	result := <-firstStage

	if !reflect.DeepEqual(result.value, 2) {
		t.Errorf("result was: %v", result.value)
	}

	result = <-firstStage

	if !reflect.DeepEqual(result.value, 4) {
		t.Errorf("result was: %v", result.value)
	}
}

func TestSimplePipelineCancel(t *testing.T) {
	numbers := []int{1, 2, 3}

	input := Converter(numbers...)

	firstStage := Pipeline(multBy(2), cancel, multBy(2))(input)

	secondStage := Pipeline(identity)(firstStage)

	result := Sink(secondStage)

	// TODO: FIX THIS
	// ! UNDETERMINISTC
	// * some times this produces a negative wg count on the sink phase
	if !reflect.DeepEqual(result, []int{}) && !reflect.DeepEqual(result, []int{4, 8}) {
		t.Errorf("result was: %v", result)
	}
}
