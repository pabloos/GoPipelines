package pipelines

import (
	"reflect"
	"testing"
)

// TODO non deterministc aproach
// tests runs in a deterministc way, while fan in and out does it with a non deterministic deliver order
func TestFanInFanOut(t *testing.T) {
	numbers := []int{1, 2, 3}

	input := Converter(numbers...)

	firstStage := NewPipeline(identity)(input)

	secondStage := FanOut(firstStage, RoundRobin, NewPipeline(double), NewPipeline(square))

	merged := FanIn(secondStage...)

	thirdStage := NewPipeline(divideBy(2))(merged)

	result := Sink(thirdStage)

	t.Log(result)

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
			i := Converter(tt.input...)

			pip := NewPipeline(tt.args.functors...)(i)

			got := Sink(pip)

			for i, valueGotten := range got {
				if valueGotten != tt.want[i] {
					t.Errorf("NewPipeline() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
