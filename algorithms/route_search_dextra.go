package main

import (
	"fmt"
	"math"
)

// Graph structure
type Graph struct {
	nodes map[int]map[int]int
}

// NewGraph creates a new graph
func NewGraph() *Graph {
	return &Graph{nodes: make(map[int]map[int]int)}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(u, v, weight int) {
	if g.nodes[u] == nil {
		g.nodes[u] = make(map[int]int)
	}
	g.nodes[u][v] = weight
}

// Dijkstra finds the shortest path using Dijkstra's Algorithm
func (g *Graph) Dijkstra(start int) map[int]int {
	distances := make(map[int]int)
	for node := range g.nodes {
		distances[node] = math.MaxInt64
	}
	distances[start] = 0

	visited := make(map[int]bool)

	for len(visited) < len(g.nodes) {
		// Find the unvisited node with the smallest distance
		minNode := -1
		for node := range g.nodes {
			if _, seen := visited[node]; !seen {
				if minNode == -1 || distances[node] < distances[minNode] {
					minNode = node
				}
			}
		}

		// Mark the node as visited
		visited[minNode] = true

		// Update the distance to each neighbor
		for neighbor, weight := range g.nodes[minNode] {
			if newDist := distances[minNode] + weight; newDist < distances[neighbor] {
				distances[neighbor] = newDist
			}
		}
	}

	return distances
}

func main() {
	graph := NewGraph()

	graph.AddEdge(0, 1, 4)
	graph.AddEdge(0, 2, 1)
	graph.AddEdge(2, 1, 2)
	graph.AddEdge(1, 3, 1)
	graph.AddEdge(2, 3, 5)

	distances := graph.Dijkstra(0)
	fmt.Println("Shortest distances from node 0:", distances)
}
