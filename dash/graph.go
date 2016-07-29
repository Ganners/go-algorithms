package dash

import (
	"fmt"
	"strings"
)

// Constructs a node
type node struct {
	word       string
	neighbours []*node

	// Used during search
	parent *node
}

// Adds a neighbour to the list of neighbour edges
func (n *node) addNeighbour(neighbour *node) {
	n.neighbours = append(n.neighbours, neighbour)
}

// Prints out each nodes link, I.e.
// A -> B
// B -> A
// B -> C
// C -> B
func (n *node) String() string {
	seen := make(map[string]struct{}, 50)
	return strings.TrimSpace(toString(n, seen))
}

// Recursive string generator for String()
func toString(n *node, seen map[string]struct{}) string {
	str := "\n"
	if len(n.neighbours) == 0 {
		return ""
	}
	// If it has been seen as a root node
	for _, c := range n.neighbours {
		link := fmt.Sprintf("%s -> %s", n.word, c.word)
		if _, alreadySeen := seen[link]; !alreadySeen {
			seen[link] = struct{}{}
			str += link + toString(c, seen)
		}
	}
	return str
}

// We'll construct a graph from the dictionary, the graph representing all of
// the 1 letter changes that can exist from a given word.
//
// We also return a map to the nodes, so we can look up the start and end
//
// Where M is length of word, N is cardinality of dictionary
// The time complexity is O(M *N * N/2), space is O(N)
func constructGraph(dictionary []string) (*node, map[string]*node) {

	nodeMap := make(map[string]*node)
	root := &node{word: "*"}

	for _, word := range dictionary {
		if _, duplicate := nodeMap[word]; duplicate {
			continue
		}

		neighbour := &node{
			word: word,
		}

		if len(nodeMap) == 0 {
			root.addNeighbour(neighbour)
			nodeMap[word] = neighbour
			continue
		}

		for previous, _ := range nodeMap {
			if distance(word, previous) == 1 {
				// Set up forward and backwards edges
				neighbour.addNeighbour(nodeMap[previous])
				nodeMap[previous].addNeighbour(neighbour)

				// Set the words node
				nodeMap[word] = neighbour
			} else {
				// If it doesn't exists, we want to add it without any
				// edges. They'll get filled in if they can be
				if _, exists := nodeMap[word]; !exists {
					nodeMap[word] = neighbour
				}
			}
		}
	}

	return root, nodeMap
}

// Calculates the number of different letters between two strings
func distance(str1, str2 string) int {
	differences := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			differences++
		}
	}
	return differences
}
