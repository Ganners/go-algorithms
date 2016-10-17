package fibonacci

import "math"

var (
	FibCache = make(map[float64]float64, 100)
)

// Not thread safe
func FibonacciCached(n float64) float64 {
	res, ok := FibCache[n]
	if ok {
		return res
	}
	res = Fibonacci(n)
	FibCache[n] = res
	return res
}

// Closed form Fibonacci calculation
func Fibonacci(n float64) float64 {
	const (
		sqrt5  float64 = 2.2360679775  // = math.Sqrt(5.0)
		golden float64 = 1.61803398875 // = (1.0 + sqrt5) / 2.0
	)
	return math.Floor((math.Pow(golden, n) / sqrt5) + 0.5)
}
