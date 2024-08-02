package main

import (
	"fmt"
	"math"
)

type Edge struct {
	to     int
	weight int
}

func dijkstra(graph [][]Edge, start int) []int {
	n := len(graph)
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt32
	}

	dist[start] = 0
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		minDist := math.MaxInt32
		var u int
		for j := 0; j < n; j++ {
			if !visited[j] && dist[j] < minDist {
				minDist = dist[j]
				u = j
			}
		}
		visited[u] = true
		for _, e := range graph[u] {
			if alt := dist[u] + e.weight; alt < dist[e.to] {
				dist[e.to] = alt
			}
		}
	}

	return dist
}

func main() {

	graph := [][]Edge{
		{{1, 10}, {2, 3}},        // 0
		{{2, 1}, {3, 2}},         // 1
		{{1, 4}, {3, 8}, {4, 2}}, // 2
		{{4, 7}},                 // 3
		{{0, 3}},                 // 4
	}

	dist := dijkstra(graph, 0)
	fmt.Println(dist) // [0 7 3 9 5]
}
