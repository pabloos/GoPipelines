package pipelines

import (
	"context"
	"reflect"
	"testing"
)

// non deterministc aproach
func TestFanInFanOut(t *testing.T) {
	ctx, cancelCtx := context.WithCancel(context.Background())

	defer cancelCtx()

	numbers := []int{1, 2, 3}

	input := Converter(numbers...)

	firstStage := Pipeline(ctx, identity)(input)

	secondStage := FanOut(ctx, firstStage, RoundRobin, Pipeline(ctx, double), Pipeline(ctx, square))

	merged := FanIn(ctx, secondStage...)

	thirdStage := Pipeline(ctx, divideBy(2))(merged)

	result := Sink(thirdStage)

	if !reflect.DeepEqual(result, numbers) &&
		!reflect.DeepEqual(result, []int{1, 3, 2}) &&
		!reflect.DeepEqual(result, []int{2, 1, 3}) &&
		!reflect.DeepEqual(result, []int{2, 3, 1}) {
		t.Errorf("result was: %v", result)
	}
}

func TestNewPipeline(t *testing.T) {
	t.Parallel()

	type args struct {
		functors []functor
	}

	tests := []struct {
		name  string
		args  args
		want  []int
		input []int
	}{
		{
			name: "zero input value",
			args: args{
				functors: []functor{
					double,
				},
			},
			input: []int{},
			want:  []int{},
		},
		{
			name: "one input value",
			args: args{
				functors: []functor{
					double,
				},
			},
			input: []int{
				1,
			},
			want: []int{
				2,
			},
		},
		{
			name: "multiple input values",
			args: args{
				functors: []functor{
					double,
				},
			},
			input: []int{
				1, 2, 3,
			},
			want: []int{
				2, 4, 6,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())

			defer cancel()

			i := Converter(tt.input...)

			pip := Pipeline(ctx, tt.args.functors...)(i)

			got := Sink(pip)

			for i, valueGotten := range got {
				if valueGotten != tt.want[i] {
					t.Errorf("Pipeline() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
