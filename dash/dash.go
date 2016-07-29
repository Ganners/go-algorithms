package dash

import "errors"

const (
	inf = 0x7FF0000000000000
)

var (
	ErrLengthsDoNotMatch = errors.New("Lengths do not match")
	ErrStartWordNotFound = errors.New("Start word could not be found in dictionary")
	ErrEndWordNotFound   = errors.New("End word could not be found in dictionary")
)

// DictionaryDash will take a start and end word, and a dictionary.
// Using that dictionary it will find the length with the shortest
// transformation
func DictionaryDash(start, end string, dictionary []string) (int, error) {

	if len(start) != len(end) {
		return 0, ErrLengthsDoNotMatch
	}

	// If they are equal, we have done it in zero steps
	if start == end {
		return 0, nil
	}

	// Construct a graph from the dictionary
	_, dictMap := constructGraph(dictionary)

	// If the words don't exist in the dictionary, there's no point
	// trying to continue
	if _, found := dictMap[start]; !found {
		return 0, ErrStartWordNotFound
	}
	if _, found := dictMap[end]; !found {
		return 0, ErrEndWordNotFound
	}

	// Execute a dijkstra search, we only care about the tentative weight
	// between the nodes
	_, steps := dijkstra(dictMap, dictMap[start], dictMap[end])

	return steps, nil
}

// Performs dijkstra search to find the shortest path between the start
// and end node
func dijkstra(graph map[string]*node, start, end *node) ([]*node, int) {

	// Reset graph parents
	for _, n := range graph {
		n.parent = nil
	}

	q := make(map[*node]int, len(graph))
	visited := make(map[*node]int, len(graph))

	// Set the start node to visited with weight 0
	visited[start] = 0
	current := start

	for {
		for _, neighbour := range current.neighbours {
			if _, found := visited[neighbour]; found {
				continue
			}
			if weight, found := q[neighbour]; found && weight < visited[current]+1 {
				continue
			}
			q[neighbour] = visited[current] + 1
		}

		// Break out when we've used up the q
		if len(q) == 0 || visited[end] != 0 {
			break
		}

		min := inf
		for node, weight := range q {
			if weight <= min {
				current = node
				min = weight
			}
		}

		// Remove current from the queue
		delete(q, current)

		// Set the visited to the min
		visited[current] = min
	}

	// The path for which it was found in
	path := []*node{}

	if weight, found := visited[end]; found {

		path = make([]*node, weight+1)

		for v, w := range visited {
			if w <= weight {
				path[w] = v
			}
		}

		return path, weight
	}

	return path, inf
}
