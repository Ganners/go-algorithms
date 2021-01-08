package polynomial

func InterpolatingPolynomial(x float64, xs, ys []float64) float64 {
	total := 0.0
	for i := 0; i < len(xs); i++ {
		mul := 1.0
		for j := 0; j < len(xs); j++ {
			if j == i {
				continue
			}
			mul *= (x - xs[j]) / (xs[i] - xs[j])
		}
		total += ys[i] * mul
	}
	return total
}
