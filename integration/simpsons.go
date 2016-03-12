package integration

type F func(x float64) float64

// Implementation of Simpson's rule:
//  \int_{a}^{b} f(x) \, dx \approx
//      \tfrac{b-a}{6}\left[f(a) + 4f\left(\tfrac{a+b}{2}\right)+f(b)\right].
func Simpson(f F, a, b, n float64) float64 {

	if n <= 1 {
		return ((b - a) / 6.0) * (f(a) + (4 * f(a+b/2.0)) + f(b))
	}

	h := (b - a) / n
	s := f(a) + f(b)

	for i := 1.0; i <= n; i += 2.0 {
		s += 4.0 * f(a+i*h)
	}

	for i := 2.0; i <= n-1.0; i += 2.0 {
		s += 2 * f(a+i*h)
	}

	return s * h / 3.0
}
