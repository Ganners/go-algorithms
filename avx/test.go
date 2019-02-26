package avx

func grow(f []float64) []float64 {
	increase := len(f) % 4
	for i := 0; i < increase; i++ {
		f = append(f, 0.0)
	}
	return f
}
