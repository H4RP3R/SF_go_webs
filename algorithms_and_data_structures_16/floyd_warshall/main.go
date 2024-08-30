package main

import (
	"fmt"
	"math"
)

const inf = math.MaxInt32

func FloydWarshall(graph [][]int) [][]int {
	n := len(graph)
	dist := make([][]int, n)

	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		copy(dist[i], graph[i])
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}

func main() {
	graph := [][]int{
		{0, 5, inf, 10},
		{inf, 0, 3, inf},
		{inf, inf, 0, 1},
		{inf, inf, inf, 0},
	}

	result := FloydWarshall(graph)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			if result[i][j] == inf {
				fmt.Printf("Inf ")
			} else {
				fmt.Printf("%d ", result[i][j])
			}
		}
		fmt.Println()
	}
}
