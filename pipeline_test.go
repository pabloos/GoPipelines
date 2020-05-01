package main

import (
	"reflect"
	"testing"
)

func TestPipeline_Exec(t *testing.T) {
	type fields struct {
		source source
		end    end
		stages stages
	}
	type args struct {
		input []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "One Stage case",
			fields: fields{
				Converter,
				Sink,
				genStages(func(number int) int {
					return number + 1
				}),
			},
			args: args{
				input: []int{
					1, 2, 3,
				},
			},
			want: []int{
				2, 3, 4,
			},
		},
		{
			name: "Various stages case",
			fields: fields{
				Converter,
				Sink,
				genStages(
					func(number int) int {
						return number + 1
					},
					func(number int) int {
						return number * 2
					},
				),
			},
			args: args{
				input: []int{
					1, 2, 3,
				},
			},
			want: []int{
				4, 6, 8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pip := &Pipeline{
				source: tt.fields.source,
				stages: tt.fields.stages,
				end:    tt.fields.end,
			}
			if got := pip.Exec(tt.args.input...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pipeline.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecursivePipeline_Exec(t *testing.T) {
	type fields struct {
		source source
		end    end
		stages stages
	}
	type args struct {
		input []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "One Stage case",
			fields: fields{
				Converter,
				Sink,
				genStages(func(number int) int {
					return number + 1
				}),
			},
			args: args{
				input: []int{
					1, 2, 3,
				},
			},
			want: []int{
				3, 4, 5,
			},
		},
		{
			name: "Various stages case",
			fields: fields{
				Converter,
				Sink,
				genStages(
					func(number int) int {
						return number + 1
					},
					func(number int) int {
						return number * 2
					},
				),
			},
			args: args{
				input: []int{
					1, 2, 3,
				},
			},
			want: []int{
				10, 14, 18,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pip := &Pipeline{
				source: tt.fields.source,
				stages: tt.fields.stages,
				end:    tt.fields.end,
			}
			if got := pip.Exec(pip.Exec(tt.args.input...)...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pipeline.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
