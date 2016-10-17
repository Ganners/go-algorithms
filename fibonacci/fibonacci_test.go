package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {

	testCases := []struct {
		input  float64
		output float64
	}{
		{
			input:  0,
			output: 0,
		},
		{
			input:  1,
			output: 1,
		},
		{
			input:  2,
			output: 1,
		},
		{
			input:  3,
			output: 2,
		},
		{
			input:  4,
			output: 3,
		},
		{
			input:  5,
			output: 5,
		},
		{
			input:  6,
			output: 8,
		},
		{
			input:  7,
			output: 13,
		},
		{
			input:  8,
			output: 21,
		},
		{
			input:  9,
			output: 34,
		},
	}

	for _, tc := range testCases {
		out := Fibonacci(tc.input)
		if out != tc.output {
			t.Errorf("result %f does not match expected %f", out, tc.output)
		}
	}
}

func BenchmarkFibonacci20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(20.0)
	}
}

func BenchmarkFibonacci40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(40.0)
	}
}

func BenchmarkFibonacci80(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(80.0)
	}
}

func BenchmarkFibonacciCached20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciCached(20.0)
	}
}

func BenchmarkFibonacciCached40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciCached(40.0)
	}
}

func BenchmarkFibonacciCached80(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciCached(80.0)
	}
}
