package main

import (
	"testing"
)

func TestNewPipeline(t *testing.T) {
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
