package avx

import (
	"testing"
)

func TestSumFloatAVX(t *testing.T) {
	{
		f := []float64{1, 2, 3, 4, 5, 6, 7, 8}
		sum := SumFloatAVX(f)
		if sum != 36 {
			t.Errorf("sum %v does not match expected %v", sum, 36)
		}
	}
	{
		f := []float64{7, 5, 8, 4, 1, 9, 8, 1}
		sum := SumFloatAVX(f)
		if sum != 43 {
			t.Errorf("sum %v does not match expected %v", sum, 45)
		}
	}
}

func TestSumFloat(t *testing.T) {
	{
		f := []float64{1, 2, 3, 4, 5, 6, 7, 8}
		sum := SumFloat(f)
		if sum != 36 {
			t.Errorf("sum %v does not match expected %v", sum, 36)
		}
	}
	{
		f := []float64{7, 5, 8, 4, 1, 9, 8, 1}
		sum := SumFloat(f)
		if sum != 43 {
			t.Errorf("sum %v does not match expected %v", sum, 45)
		}
	}
}

func BenchmarkSumFloatAVX(b *testing.B) {
	input := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < b.N; i++ {
		_ = sumFloatAVX(input)
	}
}

func BenchmarkSumFloat(b *testing.B) {
	input := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < b.N; i++ {
		_ = SumFloat(input)
	}
}
