package vectorshift

import (
	"reflect"
	"testing"
)

type TestCase struct {
	In     []int
	Shift  int
	Expect []int
}

func getTestCases() []TestCase {
	return []TestCase{
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
}

func TestShift(t *testing.T) {
	for _, test := range getTestCases() {
		out := Shift(test.In, test.Shift)
		if !reflect.DeepEqual(test.Expect, out) {
			t.Errorf("Expected %v, got %v", test.Expect, out)
		}
	}
}

func TestShift2(t *testing.T) {
	for _, test := range getTestCases() {
		out := Shift2(test.In, test.Shift)
		if !reflect.DeepEqual(test.Expect, out) {
			t.Errorf("Expected %v, got %v", test.Expect, out)
		}
	}
}

func BenchmarkShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Shift([]int{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
			31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
			41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
		}, -7)
	}
}

func BenchmarkShift2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Shift2([]int{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
			31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
			41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
		}, -7)
	}
}
