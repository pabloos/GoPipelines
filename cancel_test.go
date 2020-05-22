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

	// <-secondStage
	// result := <-secondStage

	// TODO: FIX THIS
	// ! UNDETERMINISTC
	// * some times the first and maybe the second number arrives soon to the sink phase
	// * and the pipeline returns [4] or [4, 8]
	if !reflect.DeepEqual(result, []int{}) &&
		!reflect.DeepEqual(result, []int{4}) &&
		!reflect.DeepEqual(result, []int{4, 8}) {
		t.Errorf("result was: %v", result)
	}
}

// func TestSimpleFanInFanOutCancel(t *testing.T) {
// 	numbers := []int{1, 2, 3}

// 	input := Converter(numbers...)

// 	firstStage := NewPipeline(identity)(input)

// 	secondStage := FanOut(firstStage, RoundRobin, NewPipeline(cancel))

// 	merged := FanIn(secondStage...)

// 	thirdStage := NewPipeline(divideBy(2))(merged)

// 	result := Sink(thirdStage)

// 	t.Log(result)

// 	if !reflect.DeepEqual(result, []int{}) {
// 		t.Errorf("result was: %v", result)
// 	}
// }

// func TestFanInFanOutCancel(t *testing.T) {
// 	numbers := []int{1, 2, 3}

// 	input := Converter(numbers...)

// 	firstStage := NewPipeline(identity)(input)

// 	secondStage := FanOut(firstStage, RoundRobin, NewPipeline(cancel), NewPipeline(square))

// 	merged := FanIn(secondStage...)

// 	thirdStage := NewPipeline(divideBy(2))(merged)

// 	result := Sink(thirdStage)

// 	t.Log(result)

// 	if !reflect.DeepEqual(result, []int{}) {
// 		t.Errorf("result was: %v", result)
// 	}
// }
