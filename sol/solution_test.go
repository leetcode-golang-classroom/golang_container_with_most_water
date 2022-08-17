package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	for idx := 0; idx < b.N; idx++ {
		maxArea(height)
	}
}
func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "height = [1,8,6,2,5,4,8,3,7]",
			args: args{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			want: 49,
		},
		{
			name: "height = [1,1]",
			args: args{height: []int{1, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
