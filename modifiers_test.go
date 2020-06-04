package pipelines

// func Test_add2(t *testing.T) {
// 	type args struct {
// 		number int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{
// 			"one",
// 			args{
// 				1,
// 			},
// 			3,
// 		},
// 		{
// 			"two",
// 			args{
// 				2,
// 			},
// 			4,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got, _ := addTo(2)(tt.args.number); got != tt.want {
// 				t.Errorf("add2() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_square(t *testing.T) {
// 	type args struct {
// 		number int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{
// 			"three",
// 			args{
// 				3,
// 			},
// 			9,
// 		},
// 		{
// 			"two",
// 			args{
// 				2,
// 			},
// 			4,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got, _ := square(tt.args.number); got != tt.want {
// 				t.Errorf("square() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
