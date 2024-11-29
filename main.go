// main.go
package main

import (
	"container/list"
	"fmt"
)

// Graph represents an undirected graph
type Graph struct {
	adjacencyList map[int]*list.List
}

// NewGraph creates a new empty graph
func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int]*list.List),
	}
}

// AddVertex adds a new vertex to the graph
func (g *Graph) AddVertex(v int) {
	g.adjacencyList[v] = list.New()
}

// AddEdge adds an edge between two vertices
func (g *Graph) AddEdge(v1, v2 int) {
	// check if the vertex already exist
	// if not add a vertex
	if g.adjacencyList[v1] == nil {
		g.adjacencyList[v1] = list.New()
	}
	if g.adjacencyList[v2] == nil {
		g.adjacencyList[v2] = list.New()
	}
	g.adjacencyList[v1].PushBack(v2)
	g.adjacencyList[v2].PushBack(v1)
}

// BFS performs a breadth-first search starting from the given vertex
func BFS(g *Graph, start int) []int {
	// A map to track which vertices have been visited to avoid cycles.
	visited := make(map[int]bool)
	result := []int{}
	queue := list.New()

	// Mark the start node as visited and enqueue it
	visited[start] = true
	queue.PushFront(start)
	// The loop continues until there are no more vertices in the queue.
	for queue.Len() > 0 {
		// get the first element in the list with it value
		current := queue.Front().Value.(int)
		// after that remove it from the list
		queue.Remove(queue.Front())
		// If the current vertex is not already in the result, it is added.

		if len(result) == 0 || current != result[len(result)-1] {
			result = append(result, current)
		}

		// Dequeue all nodes adjacent to this node
		for neighborElement := g.adjacencyList[current].Front(); neighborElement != nil; neighborElement = neighborElement.Next() {
			neighbor := neighborElement.Value.(int)
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.PushFront(neighbor)
			}
		}
	}

	return result
}

func main() {
	graph := NewGraph()
	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)

	fmt.Println("BFS traversal:", BFS(graph, 0))
}
