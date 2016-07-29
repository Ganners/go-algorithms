package dash

import "errors"

var (
	// Possible errors
	ErrLowercaseAlphaOnly = errors.New("Please only use lowercase alphabetic characters only")
	ErrLengthsDoNotMatch  = errors.New("Lengths do not match")
	ErrStartWordNotFound  = errors.New("Start word could not be found in dictionary")
	ErrEndWordNotFound    = errors.New("End word could not be found in dictionary")
)

// DictionaryDash will contain a dictionary (stored as a graph) and
// allow it to be queried (with Dash) to find the shortest number of
// transformations to get from one word to another
type DictionaryDash struct {
	graph map[string]*node
}

// Returns a dictionary dash, it will contain a graph for which
// multiple dashes can be run on to find the minimum length of 1
// character transformations
func NewDictionaryDash(dictionary []string) *DictionaryDash {
	_, dictionaryMap := constructGraph(dictionary)
	return &DictionaryDash{
		graph: dictionaryMap,
	}
}

// Dash will validate what we're asking for and then run Dijkstras
// algorithm to find the minimum weighted path in the dictionary graph
func (dd *DictionaryDash) Dash(start, end string) (int, error) {

	for _, l := range start + end {
		if l < 'a' || l > 'z' {
			return 0, ErrLowercaseAlphaOnly
		}
	}

	if len(start) != len(end) {
		return 0, ErrLengthsDoNotMatch
	}

	// If they are equal, we have done it in zero steps
	if start == end {
		return 0, nil
	}

	// If the words don't exist in the dictionary, there's no point
	// trying to continue
	if _, found := dd.graph[start]; !found {
		return 0, ErrStartWordNotFound
	}
	if _, found := dd.graph[end]; !found {
		return 0, ErrEndWordNotFound
	}

	// Execute a dijkstra search, we only care about the tentative weight
	// between the nodes. If we want to see an output of the path that was
	// taken then we can return the first component
	transformations := dijkstra(dd.graph, dd.graph[start], dd.graph[end])

	return transformations, nil
}
