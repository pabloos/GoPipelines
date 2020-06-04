package pipelines

import (
	"context"
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
					{
						value:    4,
						orderNum: 1,
					},
					{
						value:    32,
						orderNum: 2,
					},
					{
						value:    1,
						orderNum: 3,
					},
				},
			},
			want: []Element{
				{
					value:    4,
					orderNum: 1,
				},
				{
					value:    32,
					orderNum: 2,
				},
				{
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
					{
						value:    4,
						orderNum: 1,
					},
					{
						value:    1,
						orderNum: 3,
					},
					{
						value:    32,
						orderNum: 2,
					},
				},
			},
			want: []Element{
				{
					value:    4,
					orderNum: 1,
				},
				{
					value:    32,
					orderNum: 2,
				},
				{
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
					{
						value:    32,
						orderNum: 2,
					},
					{
						value:    1,
						orderNum: 3,
					},
					{
						value:    4,
						orderNum: 1,
					},
				},
			},
			want: []Element{
				{
					value:    1,
					orderNum: 3,
				},
				{
					value:    32,
					orderNum: 2,
				},
				{
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

func TestNoOrderedSink(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	input := Converter(3, 1, 2, 5)

	pipe := Pipeline(ctx, double)(input)

	results := SinkWithOrder(pipe, NoOrder)

	wanted := []Element{
		{
			orderNum: 0,
			value:    6,
		},
		{
			orderNum: 1,
			value:    2,
		},
		{
			orderNum: 2,
			value:    4,
		},
		{
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

func TestOrderedSink(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	input := Converter(3, 1, 2, 5)

	pipe := Pipeline(ctx, double)(input)

	results := SinkWithOrder(pipe, InOrder)

	wanted := []Element{
		{
			orderNum: 0,
			value:    6,
		},
		{
			orderNum: 1,
			value:    2,
		},
		{
			orderNum: 2,
			value:    4,
		},
		{
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

func TestReverseSink(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	input := Converter(3, 1, 2, 5)

	pipe := Pipeline(ctx, double)(input)

	results := SinkWithOrder(pipe, Reverse)

	wanted := []Element{
		{
			orderNum: 3,
			value:    10,
		},
		{
			orderNum: 2,
			value:    4,
		},
		{
			orderNum: 1,
			value:    2,
		},
		{
			orderNum: 0,
			value:    6,
		},
	}

	for i, result := range results {
		if result != wanted[i].value {
			t.Errorf("Wanted: %d, Got: %d", wanted[i], result)
		}
	}
}
