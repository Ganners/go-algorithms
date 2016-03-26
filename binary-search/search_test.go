package search

import "testing"

func TestBinarySEarch(t *testing.T) {

	testCases := []struct {
		Needle   int
		Haystack []int
		Expected int
	}{
		{
			Needle:   1,
			Haystack: []int{1, 3, 5, 7, 9},
			Expected: 1,
		},
		{
			Needle:   3,
			Haystack: []int{1, 3, 5, 7, 9},
			Expected: 3,
		},
		{
			Needle:   5,
			Haystack: []int{1, 3, 5, 7, 9},
			Expected: 5,
		},
		{
			Needle:   7,
			Haystack: []int{1, 3, 5, 7, 9},
			Expected: 7,
		},
		{
			Needle:   9,
			Haystack: []int{1, 3, 5, 7, 9},
			Expected: 9,
		},
		{
			Needle:   6,
			Haystack: []int{1, 3, 5, 7, 9},
			Expected: -1,
		},
		{
			Needle:   2,
			Haystack: []int{1, 3, 5, 7, 9},
			Expected: -1,
		},
		{
			Needle:   1,
			Haystack: []int{1, 3, 5, 7, 9, 100},
			Expected: 1,
		},
		{
			Needle:   9,
			Haystack: []int{1, 3, 5, 7, 9, 100},
			Expected: 9,
		},
		{
			Needle:   0,
			Haystack: []int{1, 3, 5, 7, 9, 100},
			Expected: -1,
		},
		{
			Needle:   -1,
			Haystack: []int{1, 3, 5, 7, 9, 100},
			Expected: -1,
		},
	}

	for _, test := range testCases {
		res := BinarySearch(test.Needle, test.Haystack)

		if res != test.Expected {
			t.Errorf("Expected result to be %d, got %d", test.Expected, res)
		}
	}
}
