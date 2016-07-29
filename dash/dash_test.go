package dash

import "testing"

// Effectively tests the Dijkstra search, but with the validation
// wrapper of the Dash function
func TestDictionaryDash(t *testing.T) {
	for _, test := range []struct {
		start           string
		end             string
		dictionary      []string
		transformations int
	}{
		{
			start: "hit",
			end:   "cog",
			dictionary: []string{
				"hit",
				"dot",
				"dog",
				"cog",
				"log",
				"hot",
			},
			transformations: 4,
		},
		{
			start: "hit",
			end:   "log",
			dictionary: []string{
				"hit",
				"dot",
				"dog",
				"cog",
				"log",
				"hot",
			},
			transformations: 4,
		},
		{
			start: "hit",
			end:   "hot",
			dictionary: []string{
				"hit",
				"dot",
				"dog",
				"cog",
				"hot",
				"log",
			},
			transformations: 1,
		},
		{
			start: "log",
			end:   "cog",
			dictionary: []string{
				"hit",
				"dot",
				"dog",
				"cog",
				"hot",
				"log",
			},
			transformations: 1,
		},
		{
			start: "hot",
			end:   "hot",
			dictionary: []string{
				"hit",
				"dot",
				"dog",
				"cog",
				"hot",
				"log",
			},
			transformations: 0,
		},
		{
			start: "hit",
			end:   "bag",
			dictionary: []string{
				"hit",
				"hut",
				"hot",
				"dot",
				"dog",
				"cog",
				"cag",
				"bog",
				"bag",
			},
			transformations: 5,
		},
	} {

		// Construct the graph and execute a single dash on it
		dd := NewDictionaryDash(test.dictionary)
		transformations, err := dd.Dash(test.start, test.end)

		if err != nil {
			t.Errorf("Did not expect error, got %s", err)
		}

		if transformations != test.transformations {
			t.Errorf("Number of transformations %d does not match expected %d",
				transformations,
				test.transformations)
		}
	}
}
