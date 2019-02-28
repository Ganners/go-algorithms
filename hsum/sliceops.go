package avx

func SumFloat(f []float64) float64 {
	sum := 0.0
	for i := 0; i < len(f); i++ {
		sum += f[i]
	}
	return sum
}

func SumFloatAVX(f []float64) float64

func SumFloat32AVX(f []float32) float32

func IDK(f []float64) []float64
