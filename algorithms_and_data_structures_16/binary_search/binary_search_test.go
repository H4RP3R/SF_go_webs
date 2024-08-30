package binarysearch

import "testing"

func Test_binSearch(t *testing.T) {
	type args struct {
		arr []int
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test case 1", args{[]int{1, 2, 3, 4, 5}, 3}, 2},
		{"Test case 2", args{[]int{1, 2, 3, 4, 5}, 10}, -1},
		{"Test case 3", args{[]int{1, 2, 3, 4, 5}, 0}, -1},
		{"Test case 4", args{[]int{1, 2, 3, 4, 5}, 5}, 4},
		{"Test case 5", args{[]int{1, 2, 3, 4, 5}, 4}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binSearch(tt.args.arr, tt.args.val); got != tt.want {
				t.Errorf("binSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
