package sieve

import (
	"reflect"
	"testing"
)

// Execute some tests against the sieve
func TestSieveOfEratosthenes(t *testing.T) {

	testCases := []struct {
		UpTo   int
		Expect []int
	}{
		{
			UpTo:   2,
			Expect: []int{2},
		},
		{
			UpTo:   5,
			Expect: []int{2, 3, 5},
		},
		{
			UpTo:   11,
			Expect: []int{2, 3, 5, 7, 11},
		},
	}

	// Execute all test cases
	for _, test := range testCases {

		results := Aggregate(SieveOfEratosthenes(test.UpTo))

		// Test
		if !reflect.DeepEqual(results, test.Expect) {
			t.Errorf("Result %v did not match expected %v", results, test.Expect)
		}
	}
}
