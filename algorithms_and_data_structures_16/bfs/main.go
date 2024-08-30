package main

import "fmt"

type Graph struct {
	adjacencyList map[int][]int
}

func (g *Graph) AddEdge(src, dest int) {
	g.adjacencyList[src] = append(g.adjacencyList[src], dest)
}

func BFS(g *Graph, start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		if !visited[vertex] {
			fmt.Printf("%d, ", vertex)
			visited[vertex] = true
			for _, neighbor := range g.adjacencyList[vertex] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}
		}
	}
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyList: map[int][]int{},
	}
}

func main() {
	graph := NewGraph()
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(4, 6)
	graph.AddEdge(4, 7)
	graph.AddEdge(4, 8)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 5)

	startVertex := 0
	BFS(graph, startVertex)
}
