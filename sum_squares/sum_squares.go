package sum_squares

// Efficient
func SumSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
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
