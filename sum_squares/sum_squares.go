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

// Efficient
func SumCubes(n int) int {
	s := (n * (n + 1) / 2)
	return s * s
}

// Brute force, used for comparative testing
func SumCubesBrute(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sq := i * i * i
		sum += sq
	}
	return sum
}
