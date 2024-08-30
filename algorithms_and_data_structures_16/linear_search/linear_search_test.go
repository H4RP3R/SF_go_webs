package linearsearch

import "testing"

func Test_lSearch(t *testing.T) {
	type args struct {
		arr []int
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty slice", args{[]int{}, 42}, -1},
		{"not found", args{[]int{1, 2, 3, 4, 5}, 42}, -1},
		{"found at first", args{[]int{42, 1, 2, 3}, 42}, 0},
		{"found at end", args{[]int{1, 2, 3, 42}, 42}, 3},
		{"found at middle", args{[]int{1, 2, 42, 3}, 42}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lSearch(tt.args.arr, tt.args.val); got != tt.want {
				t.Errorf("lSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
