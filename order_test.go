package pipelines

import (
	"reflect"
	"sort"
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
			if sort.SliceStable(tt.args.input, NoOrder(tt.args.input)); !reflect.DeepEqual(tt.args.input, tt.want) {
				t.Errorf("noOrder() = %v, want %v", tt.args.input, tt.want)
			}
		})
	}
}

func Test_InOrder(t *testing.T) {
	type args struct {
		input []Element
	}
	tests := []struct {
		name string
		args args
		want []Element
	}{
		{
			name: "1,3,2 => 1,2,3",
			args: args{
				[]Element{
					Element{
						value:    4,
						orderNum: 1,
					},
					Element{
						value:    1,
						orderNum: 3,
					},
					Element{
						value:    32,
						orderNum: 2,
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
			if sort.SliceStable(tt.args.input, InOrder(tt.args.input)); !reflect.DeepEqual(tt.args.input, tt.want) {
				t.Errorf("noOrder() = %v, want %v", tt.args.input, tt.want)
			}
		})
	}
}

func Test_Reverse(t *testing.T) {
	type args struct {
		input []Element
	}
	tests := []struct {
		name string
		args args
		want []Element
	}{
		{
			name: "2,3,1 => 1,2,3",
			args: args{
				[]Element{

					Element{
						value:    32,
						orderNum: 2,
					},
					Element{
						value:    1,
						orderNum: 3,
					},
					Element{
						value:    4,
						orderNum: 1,
					},
				},
			},
			want: []Element{
				Element{
					value:    1,
					orderNum: 3,
				},
				Element{
					value:    32,
					orderNum: 2,
				},
				Element{
					value:    4,
					orderNum: 1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if sort.SliceStable(tt.args.input, Reverse(tt.args.input)); !reflect.DeepEqual(tt.args.input, tt.want) {
				t.Errorf("noOrder() = %v, want %v", tt.args.input, tt.want)
			}
		})
	}
}

func TestOrderedSink(t *testing.T) {
	input := Converter(3, 1, 2, 5)

	pipe := Pipeline(double)(input)

	results := SinkWithOrder(pipe, InOrder)

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
