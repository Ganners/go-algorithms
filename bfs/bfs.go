package main

import (
	"container/list"
	"log"
)

type Node struct {
	Neighbours []*Node
	Data       int
}

type Graph struct {
	Root *Node
}

func New(nodes []int, adjacent map[int][]int) *Graph {

	if len(nodes) == 0 {
		return &Graph{}
	}

	// Inefficient for memory but convenient
	nodesCreated := make(map[int]*Node, len(nodes))

	var rootNode *Node

	// First generate all of the nodes
	for _, nodeVal := range nodes {
		node := &Node{
			Data: nodeVal,
		}

		if rootNode == nil {
			rootNode = node
		}

		nodesCreated[nodeVal] = node
	}

	// Second, generate edges
	for _, node := range nodesCreated {

		for _, adjVal := range adjacent[node.Data] {

			if n, found := nodesCreated[adjVal]; found {
				node.Neighbours = append(node.Neighbours, n)
			}
		}
	}

	graph := &Graph{
		Root: rootNode,
	}
	return graph
}

// Look for a target in the fewest possible moves, return the list of moves
func (g *Graph) BreadthFirstSearch(target int) map[int]int {

	seen := make(map[*Node]int)
	frontier := list.New()
	parents := make(map[int]int)

	// Set start point
	frontier.PushBack(g.Root)
	var nodeFound *Node

	for frontier.Len() > 0 {

		// pop()
		currentNode := frontier.Remove(frontier.Front()).(*Node)
		for _, node := range currentNode.Neighbours {
			if _, found := seen[node]; !found {

				seen[node] = 1
				parents[node.Data] = currentNode.Data

				if node.Data == target {
					// We're done, delete all from frontier and break out
					nodeFound = node
					frontier = list.New()
					break
				}

				// push()
				frontier.PushBack(node)
			}
		}
	}

	// Produce a reverse map to show how to get to root node from target
	if nodeFound != nil {
		linkToParent := make(map[int]int)
		for nodeFound.Data != g.Root.Data {
			linkToParent[nodeFound.Data] = parents[nodeFound.Data]
			nodeFound.Data = parents[nodeFound.Data]
		}
		return linkToParent
	}

	return parents
}

func main() {

	// Builds a directed graph (edges are one way)
	nodes := []int{1, 3, 7, 9, 11, 17, 23}
	adj := make(map[int][]int)
	adj[1] = []int{3, 7}
	adj[3] = []int{7}
	adj[7] = []int{9}
	adj[9] = []int{11, 23}
	adj[23] = []int{17}

	graph := New(nodes, adj)

	log.Println(graph.BreadthFirstSearch(17))
}
