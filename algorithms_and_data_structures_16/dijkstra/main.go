package main

import (
	"fmt"
	"math"
)

const inf = math.MaxInt32

func minDist(dist []int, visited []bool) int {
	min := inf
	minIndex := -1

	for v := 0; v < len(dist); v++ {
		if !visited[v] && dist[v] <= min {
			min = dist[v]
			minIndex = v
		}
	}

	return minIndex
}

func Dijkstra(graph [][]int, start int) []int {
	n := len(graph)
	dist := make([]int, n)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		dist[i] = inf
		visited[i] = false
	}

	dist[start] = 0

	for i := 0; i < n-1; i++ {
		u := minDist(dist, visited)
		visited[u] = true

		for v := 0; v < n; v++ {
			if !visited[v] && graph[u][v] != 0 && dist[u] != inf && dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v]
			}
		}
	}

	return dist
}

func main() {
	graph := [][]int{
		{0, 10, 20, 0, 0},
		{10, 0, 0, 50, 10},
		{20, 0, 0, 20, 33},
		{0, 50, 20, 0, 20},
		{0, 10, 33, 20, 0},
	}
	start := 0
	distances := Dijkstra(graph, start)
	for i := range distances {
		fmt.Printf("%d -> %d = %d\n", start, i, distances[i])
	}
}
