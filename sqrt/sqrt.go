package sqrt

// Sqrt will solve the sqrt to 5 significant figures using Newton's method
func Sqrt(a float64) float64 {
	// solve x^2 - a = 0
	// a <- (x ^ 2 - a) / (2x)

	if a < 0 {
		return 0
	}

	x := a / 2
	for {
		fx := x*x - a
		if fx < 1e-5 {
			break
		}
		fxPrime := 2 * x
		x -= fx / fxPrime
	}
	return x
}
