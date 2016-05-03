package vectorshift

import (
	"reflect"
	"testing"
)

func TestShift(t *testing.T) {

	testCases := []struct {
		In     []int
		Shift  int
		Expect []int
	}{
		{
			In:     []int{1, 2, 3, 4, 5},
			Shift:  -1,
			Expect: []int{2, 3, 4, 5, 1},
		},
		{
			In:     []int{1, 2, 3, 4, 5},
			Shift:  -3,
			Expect: []int{4, 5, 1, 2, 3},
		},
		{
			In:     []int{1, 2, 3, 4, 5},
			Shift:  -33,
			Expect: []int{4, 5, 1, 2, 3},
		},
		{
			In:     []int{1, 2, 3, 4, 5},
			Shift:  3,
			Expect: []int{3, 4, 5, 1, 2},
		},
		{
			In:     []int{1, 2, 3, 4, 5},
			Shift:  5,
			Expect: []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range testCases {
		out := Shift(test.In, test.Shift)
		if !reflect.DeepEqual(test.Expect, out) {
			t.Errorf("Expected %v, got %v", test.Expect, out)
		}
	}
}
