package sum_list

import "testing"

func BenchmarkSumSeriesRecursive1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSeriesRecursive(1000)
	}
}

func BenchmarkSumSeriesLoop1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSeriesLoop(1000)
	}
}

func BenchmarkSumSeriesFormula1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSeriesFormula(1000)
	}
}

func BenchmarkSumSeriesRecursive1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSeriesRecursive(1000000)
	}
}

func BenchmarkSumSeriesLoop1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSeriesLoop(1000000)
	}
}

func BenchmarkSumSeriesFormula1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSeriesFormula(1000000)
	}
}

func TestSumSeries(t *testing.T) {

	testCases := []struct {
		In  int
		Out int
	}{
		{
			In:  1,
			Out: 1,
		},
		{
			In:  2,
			Out: 3,
		},
		{
			In:  3,
			Out: 6,
		},
		{
			In:  4,
			Out: 10,
		},
		{
			In:  5,
			Out: 15,
		},
		{
			In:  100,
			Out: 5050,
		},
	}

	for _, test := range testCases {
		{
			n := SumSeriesRecursive(test.In)
			if n != test.Out {
				t.Errorf("Expected %d, got %d", test.Out, n)
			}
		}
		{
			n := SumSeriesLoop(test.In)
			if n != test.Out {
				t.Errorf("Expected %d, got %d", test.Out, n)
			}
		}
		{
			n := SumSeriesFormula(test.In)
			if n != test.Out {
				t.Errorf("Expected %d, got %d", test.Out, n)
			}
		}
	}
}
