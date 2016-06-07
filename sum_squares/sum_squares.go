package sum_squares

import "math"

// Efficient method
func SumSquares(n int) int {
	// Convert n into a float
	f := float64(n)

	// Calculate the triangular number (sum of 0 to n)
	tri := (f * (f + 1)) / 2

	// The ratio is (n * 2/3) + 1/3, multiplied by the triangular
	// number
	sum := ((f * (2.0 / 3.0)) + (1.0 / 3.0)) * tri

	// Round the number to the nearest decimal
	return int(math.Floor(sum + 0.5))
}

// Brute force, used for comparative testing
func SumSquaresBrute(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sq := i * i
		sum += sq
	}
	return sum
}
