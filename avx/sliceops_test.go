package avx

import (
	"testing"
)

func generateSeries(l int) []float64 {
	input := make([]float64, l)
	for i := 0; i < l; i++ {
		input[i] = float64(i + 1)
	}
	return input
}

func generateSeries32(l int) []float32 {
	input := make([]float32, l)
	for i := 0; i < l; i++ {
		input[i] = float32(i + 1)
	}
	return input
}

func TestSumFloat32AVX(t *testing.T) {
	{
		f := generateSeries32(32)
		sum := SumFloat32AVX(f)
		if sum != 528 {
			t.Errorf("sum %v does not match expected %v", sum, 528)
		}
	}
	{
		f := generateSeries32(16)
		sum := SumFloat32AVX(f)
		if sum != 136 {
			t.Errorf("sum %v does not match expected %v", sum, 136)
		}
	}
	{
		f := generateSeries32(64)
		sum := SumFloat32AVX(f)
		if sum != 2080 {
			t.Errorf("sum %v does not match expected %v", sum, 2080)
		}
	}
	// {
	// 	f := generateSeries32(16)
	// 	sum := SumFloat32AVX(f)
	// 	if sum != 136 {
	// 		t.Errorf("sum %v does not match expected %v", sum, 136)
	// 	}
	// }
	// {
	// 	f := generateSeries32(8)
	// 	sum := SumFloat32AVX(f)
	// 	if sum != 36 {
	// 		t.Errorf("sum %v does not match expected %v", sum, 36)
	// 	}
	// }
}

func TestSumFloatAVX(t *testing.T) {
	{
		f := generateSeries(8)
		sum := SumFloatAVX(f)
		if sum != 36 {
			t.Errorf("sum %v does not match expected %v", sum, 36)
		}
	}
	{
		f := generateSeries(24)
		sum := SumFloatAVX(f)
		if sum != 300 {
			t.Errorf("sum %v does not match expected %v", sum, 300)
		}
	}
	{
		f := generateSeries(24)
		sum := SumFloatAVX(f)
		if sum != 300 {
			t.Errorf("sum %v does not match expected %v", sum, 300)
		}
	}
	{
		f := generateSeries(16)
		sum := SumFloatAVX(f)
		if sum != 136 {
			t.Errorf("sum %v does not match expected %v", sum, 136)
		}
	}
	{
		f := generateSeries(8)
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
		f := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
		sum := SumFloat(f)
		if sum != 136 {
			t.Errorf("sum %v does not match expected %v", sum, 136)
		}
	}
	{
		f := generateSeries(8)
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

func BenchmarkSumFloatAVX8(b *testing.B) {
	input := generateSeries(8)
	for i := 0; i < b.N; i++ {
		_ = SumFloatAVX(input)
	}
}

func BenchmarkSumFloat8(b *testing.B) {
	input := generateSeries(8)
	for i := 0; i < b.N; i++ {
		_ = SumFloat(input)
	}
}

func BenchmarkSumFloatAVX16(b *testing.B) {
	input := generateSeries(16)
	for i := 0; i < b.N; i++ {
		_ = SumFloatAVX(input)
	}
}

func BenchmarkSumFloat16(b *testing.B) {
	input := generateSeries(16)
	for i := 0; i < b.N; i++ {
		_ = SumFloat(input)
	}
}

func BenchmarkSumFloatAVX32(b *testing.B) {
	input := generateSeries(32)
	for i := 0; i < b.N; i++ {
		_ = SumFloatAVX(input)
	}
}

func BenchmarkSumFloat32(b *testing.B) {
	input := generateSeries(32)
	for i := 0; i < b.N; i++ {
		_ = SumFloat(input)
	}
}

func BenchmarkSumFloatAVX256(b *testing.B) {
	input := generateSeries(256)
	for i := 0; i < b.N; i++ {
		_ = SumFloatAVX(input)
	}
}

func BenchmarkSumFloat32AVX256(b *testing.B) {
	input := generateSeries32(256)
	for i := 0; i < b.N; i++ {
		_ = SumFloat32AVX(input)
	}
}

func BenchmarkSumFloat256(b *testing.B) {
	input := generateSeries(256)
	for i := 0; i < b.N; i++ {
		_ = SumFloat(input)
	}
}
