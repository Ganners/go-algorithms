package dash

import "testing"

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
	} {

		transformations, err := DictionaryDash(
			test.start,
			test.end,
			test.dictionary)

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
