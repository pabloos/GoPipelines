package pipelines

import (
	"context"
	"reflect"
	"testing"
)

func TestWOCanceWOSinkl(t *testing.T) {
	ctx, cancelCtx := context.WithCancel(context.Background())

	defer cancelCtx()

	numbers := []int{1, 2, 3}

	input := Converter(numbers...)

	firstStage := Pipeline(ctx, multBy(2), cancel)(input)

	result := Sink(firstStage)
	wanted := []int{}

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("result was: %v, but wanted: ", result)
	}
}

func TestSimplePipelineCancel(t *testing.T) {
	ctx, cancelCtx := context.WithCancel(context.Background())

	defer cancelCtx()

	numbers := []int{1, 2, 3}

	input := Converter(numbers...)

	firstStage := Pipeline(ctx, cancel, multBy(2))(input)

	secondStage := Pipeline(ctx, identity)(firstStage)

	result := Sink(secondStage)
	wanted := []int{}

	if !reflect.DeepEqual(result, wanted) {
		t.Errorf("result was: %v", result)
	}
}
