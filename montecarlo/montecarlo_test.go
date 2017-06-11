package montecarlo

import (
	"fmt"
	"testing"
)

func TestMonteCarlo(t *testing.T) {
	funcs := []func(int) float64{
		monteCarloNaive,
		monteCarloParallelNaive,
	}
	for i, f := range funcs {
		ratio := f(1000000)
		if fmt.Sprintf("%.2f", ratio) != "3.14" {
			t.Errorf("%d: ratio incorrect: %f", i, ratio)
		}
	}
}

func TestMonteCarloBatch(t *testing.T) {
	funcs := []func(int, int) float64{
		monteCarloBatch1,
		monteCarloBatch2,
		monteCarloBatch3,
	}
	for i, f := range funcs {
		ratio := f(1000000, 1000)
		if fmt.Sprintf("%.2f", ratio) != "3.14" {
			t.Errorf("%d: ratio incorrect: %f", i, ratio)
		}
	}
}

func BenchmarkMonteCarloNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		monteCarloNaive(1000000)
	}
}

func BenchmarkMonteCarloParallelNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		monteCarloParallelNaive(1000000)
	}
}

func BenchmarkMonteCarloBatch1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		monteCarloBatch1(1000000, 1000)
	}
}

func BenchmarkMonteCarloBatch2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		monteCarloBatch2(1000000, 1000)
	}
}

func BenchmarkMonteCarloBatch3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		monteCarloBatch3(1000000, 1000)
	}
}

func BenchmarkMonteCarloBatch3_1000000_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		monteCarloBatch3(1000000, 1)
	}
}

func BenchmarkMonteCarloBatch3_1000000_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		monteCarloBatch3(1000000, 2)
	}
}

func BenchmarkMonteCarloBatch3_1000000_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		monteCarloBatch3(1000000, 4)
	}
}

func BenchmarkMonteCarloBatch3_1000000_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		monteCarloBatch3(1000000, 8)
	}
}

func BenchmarkMonteCarloBatch3_1000000_16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		monteCarloBatch3(1000000, 16)
	}
}

func BenchmarkMonteCarloBatch3_1000000_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		monteCarloBatch3(1000000, 32)
	}
}

func BenchmarkMonteCarloBatch3_1000000_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		monteCarloBatch3(1000000, 64)
	}
}

func BenchmarkMonteCarloBatch3_1000000_128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		monteCarloBatch3(1000000, 128)
	}
}

func BenchmarkMonteCarloBatch3_1000000_256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		monteCarloBatch3(1000000, 256)
	}
}

func BenchmarkMonteCarloBatch3_1000000_512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		monteCarloBatch3(1000000, 512)
	}
}

func BenchmarkMonteCarloBatch3_1000000_1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		monteCarloBatch3(1000000, 1024)
	}
}

func BenchmarkMonteCarloBatch3_1000000_2056(b *testing.B) {
	for i := 0; i < b.N; i++ {
		monteCarloBatch3(1000000, 2056)
	}
}
