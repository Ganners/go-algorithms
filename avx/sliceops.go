package avx

func SumFloat(f []float64) float64 {
	sum := 0.0
	for i := 0; i < 8; i++ {
		sum += f[i]
	}
	return sum
}

func SumFloatAVX(f []float64) float64 {
	sum := sumFloatAVX(f)
	return sum
}

func sumFloatAVX(f []float64) float64
