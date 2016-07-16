package hashsort

import (
	"reflect"
	"sort"
	"testing"
)

func TestSuperHash(t *testing.T) {

	testCases := []struct {
		InKey  int
		InN    int
		OutRow int
		OutCol int
	}{
		{
			InKey:  5,
			InN:    100,
			OutRow: 0,
			OutCol: 5,
		},
		{
			InKey:  50,
			InN:    184,
			OutRow: 3,
			OutCol: 8,
		},
		{
			InKey:  51,
			InN:    1234,
			OutRow: 1,
			OutCol: 15,
		},
	}

	for _, test := range testCases {

		row, col := SuperHash(test.InKey, test.InN)

		if row != test.OutRow {
			t.Errorf(
				"Expected row to be %d, got %d for <%d,%d>",
				test.OutRow, row, test.InKey, test.InN)
		}

		if col != test.OutCol {
			t.Errorf(
				"Expected column to be %d, got %d for <%d,%d>",
				test.OutCol, col, test.InKey, test.InN)
		}
	}
}

func TestHashSort(t *testing.T) {

	testCases := []struct {
		In  []int
		Out []int
	}{
		{
			In:  []int{5, 9, 4, 7, 1, 5, 4, 5, 4, 3, 2, 11, 8},
			Out: []int{1, 2, 3, 4, 4, 4, 5, 5, 5, 7, 8, 9, 11},
		},
		{
			In:  fillReverse(10),
			Out: fill(10),
		},
		{
			In:  fillReverse(100),
			Out: fill(100),
		},
		{
			In:  fillReverse(1000),
			Out: fill(1000),
		},
		{
			In:  fillReverse(10000),
			Out: fill(10000),
		},
		{
			In:  fillReverse(100000),
			Out: fill(100000),
		},
	}

	for _, test := range testCases {
		if out := HashSort(test.In, test.In[0]); !reflect.DeepEqual(out, test.Out) {
			t.Error("Expected %v got %v", test.Out, out)
		}
	}
}

func BenchmarkHashSort100(b *testing.B) {
	benchSortHelper(b, 100)
}

func BenchmarkBuiltinSort100(b *testing.B) {
	benchSortHelper(b, 100)
}

func BenchmarkHashSort1000(b *testing.B) {
	benchSortHelper(b, 1000)
}

func BenchmarkBuiltinSort1000(b *testing.B) {
	benchSortHelper(b, 1000)
}

func BenchmarkHashSort10000(b *testing.B) {
	benchSortHelper(b, 10000)
}

func BenchmarkBuiltinSort10000(b *testing.B) {
	benchSortHelper(b, 10000)
}

func BenchmarkHashSort100000(b *testing.B) {
	benchSortHelper(b, 100000)
}

func BenchmarkBuiltinSort100000(b *testing.B) {
	benchSortHelper(b, 100000)
}

func BenchmarkHashSort1000000(b *testing.B) {
	benchSortHelper(b, 1000000)
}

func BenchmarkBuiltinSort1000000(b *testing.B) {
	benchSortHelper(b, 1000000)
}

// Fills an array from 1 ... n
func fill(to int) []int {
	out := make([]int, 0, to)
	for i := 1; i <= to; i++ {
		out = append(out, i*2)
	}
	return out
}

// Fills an array from n ... 1
func fillReverse(to int) []int {
	out := make([]int, 0, to)
	for i := to; i > 0; i-- {
		out = append(out, i*2)
	}
	return out
}

func benchSortHelper(b *testing.B, N int) {
	input := fillReverse(N)
	for i := 0; i < b.N; i++ {
		HashSort(input, N*2)
	}
}

func builtinSortHelper(b *testing.B, N int) {
	input := fillReverse(N)
	for i := 0; i < b.N; i++ {
		sort.Ints(input)
	}
}
