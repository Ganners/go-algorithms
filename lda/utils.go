package lda

func VecSum(a []float64) float64 {
	sum := 0.
	for _, v := range a {
		sum += v
	}
	return sum
}

func VecAdd(a []float64, b float64) {
	for i := range a {
		a[i] += b
	}
}

func VecDiv(a []float64, b float64) {
	for i := range a {
		a[i] /= b
	}
}

func VecMul(a []float64, b []float64) []float64 {
	multiplied := make([]float64, len(a))
	for i := range a {
		multiplied[i] = a[i] * b[i]
	}
	return multiplied
}

func ArgMax(a []float64) int {
	maxV := 0.0
	maxI := 0
	for i, v := range a {
		if v > maxV {
			maxI = i
			maxV = v
		}
	}
	return maxI
}
