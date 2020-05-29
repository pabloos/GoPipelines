package pipelines

import (
	"reflect"
	"testing"
)

func Test_noOrder(t *testing.T) {
	type args struct {
		input []Element
	}
	tests := []struct {
		name string
		args args
		want []Element
	}{
		{
			name: "1,2,3 => 1,2,3",
			args: args{
				[]Element{
					Element{
						value:    4,
						orderNum: 1,
					},
					Element{
						value:    32,
						orderNum: 2,
					},
					Element{
						value:    1,
						orderNum: 3,
					},
				},
			},
			want: []Element{
				Element{
					value:    4,
					orderNum: 1,
				},
				Element{
					value:    32,
					orderNum: 2,
				},
				Element{
					value:    1,
					orderNum: 3,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := noOrder(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("noOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedSink(t *testing.T) {
	input := Converter(3, 1, 2, 5)

	pipe := Pipeline(double)(input)

	results := Sink(pipe, QuickSort)

	wanted := []Element{
		Element{
			orderNum: 0,
			value:    6,
		},
		Element{
			orderNum: 1,
			value:    2,
		},
		Element{
			orderNum: 2,
			value:    4,
		},
		Element{
			orderNum: 3,
			value:    10,
		},
	}

	for i, result := range results {
		if result != wanted[i].value {
			t.Errorf("Wanted: %d, Got: %d", result, wanted[i])
		}
	}
}

func TestQuickSort(t *testing.T) {
	type args struct {
		arr []Element
	}
	tests := []struct {
		name string
		args args
		want []Element
	}{
		{
			name: "normal case",
			args: args{
				arr: []Element{
					Element{
						orderNum: 2,
						value:    3,
					},
					Element{
						orderNum: 0,
						value:    1,
					},
					Element{
						orderNum: 1,
						value:    2,
					},
				},
			},
			want: []Element{
				Element{
					orderNum: 0,
					value:    1,
				},
				Element{
					orderNum: 1,
					value:    2,
				},
				Element{
					orderNum: 2,
					value:    3,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
