package main

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	type args struct {
		graph [][]int
		start int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Simple graph",
			args: args{
				graph: [][]int{
					{0, 4, 0, 0, 0, 0, 0, 8, 0},
					{4, 0, 8, 0, 0, 0, 0, 11, 0},
					{0, 8, 0, 7, 0, 4, 0, 0, 2},
					{0, 0, 7, 0, 9, 14, 0, 0, 0},
					{0, 0, 0, 9, 0, 10, 0, 0, 0},
					{0, 0, 4, 14, 10, 0, 2, 0, 0},
					{0, 0, 0, 0, 0, 2, 0, 1, 6},
					{8, 11, 0, 0, 0, 0, 1, 0, 7},
					{0, 0, 2, 0, 0, 0, 6, 7, 0},
				},
				start: 0,
			},
			want: []int{0, 4, 12, 19, 21, 11, 9, 8, 14},
		},
		{
			name: "Small graph",
			args: args{
				graph: [][]int{
					{0, 1, 0},
					{1, 0, 1},
					{0, 1, 0},
				},
				start: 0,
			},
			want: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dijkstra(tt.args.graph, tt.args.start); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dijkstra() = %v, want %v", got, tt.want)
			}
		})
	}
}
