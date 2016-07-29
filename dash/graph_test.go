package dash

import (
	"strings"
	"testing"
)

func TestConstructGraph(t *testing.T) {

	for _, test := range []struct {
		dictionary []string
		graph      string
	}{
		{
			dictionary: []string{
				"hit",
				"hot",
				"dot",
				"dog",
				"log",
			},
			// Constructing the graph is a pain so I'll just test it
			// against the expected .String() output
			graph: strings.Join([]string{
				"* -> hit",
				"hit -> hot",
				"hot -> hit",
				"hot -> dot",
				"dot -> hot",
				"dot -> dog",
				"dog -> dot",
				"dog -> log",
				"log -> dog",
			}, "\n"),
		},
	} {
		graph, _ := constructGraph(test.dictionary)
		if graph.String() != test.graph {
			t.Errorf("Graph\n'%s'\ndoes not match\n'%s'\n", graph, test.graph)
		}
	}
}

func TestDistance(t *testing.T) {

	for _, test := range []struct {
		str1     string
		str2     string
		distance int
	}{
		{
			str1:     "cog",
			str2:     "log",
			distance: 1,
		},
		{
			str1:     "aaa",
			str2:     "bbb",
			distance: 3,
		},
		{
			str1:     "cog",
			str2:     "aof",
			distance: 2,
		},
		{
			// Only compares to the length of str1
			str1:     "cog",
			str2:     "aofabc",
			distance: 2,
		},
	} {
		dist := distance(test.str1, test.str2)
		if dist != test.distance {
			t.Errorf("Expected distance to be %d, got %d", test.distance, dist)
		}
	}
}
