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
				"hot",
				"log",
			},
			transformations: 3,
		},
	} {

		transformations := DictionaryDash(
			test.start,
			test.end,
			test.dictionary)

		if transformations != test.transformations {
			t.Errorf("Number of transformations %d does not match expected %d",
				transformations,
				test.transformations)
		}
	}
}
