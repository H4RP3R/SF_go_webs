package main

import "fmt"

type Graph struct {
	adjacencyList map[int][]int
}

func (g *Graph) AddEdge(src, dest int) {
	g.adjacencyList[src] = append(g.adjacencyList[src], dest)
}

func DFS(g *Graph, start int, visited map[int]bool) {
	if visited[start] {
		return
	}

	visited[start] = true
	fmt.Printf("Visit %d\n", start)
	for _, vertex := range g.adjacencyList[start] {
		DFS(g, vertex, visited)
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
	visited := make(map[int]bool)
	DFS(graph, startVertex, visited)
	fmt.Println(visited)
}
