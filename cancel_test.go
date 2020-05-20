package pipelines

import (
	"reflect"
	"testing"
)

func TestSimplePipelineCancel(t *testing.T) {
	numbers := []int{1, 2, 3}

	input := Converter(numbers...)

	firstStage := NewPipeline(cancel, multBy(2))(input)

	secondStage := NewPipeline(identity)(firstStage)

	result := Sink(secondStage)

	if !reflect.DeepEqual(result, []int{}) {
		t.Errorf("result was: %v", result)
	}
}
