package vectorshift

// Shifts a vector in O(N) with 1 space complexity
func Shift(N []int, x int) []int {

	n := len(N)

	// Can short circuit if there will be no shifting
	if x == n || (x > n && mod(n, x) == 0) {
		return N
	}

	lastI := 0
	lastV := N[lastI]

	for {
		shift := mod(lastI+x, n)
		lastI = shift

		N[lastI], lastV = lastV, N[shift]

		// The last shift will always be 0 (I.e. we're back where we
		// started)
		if shift == 0 {
			break
		}
	}

	return N
}

// Golang's mod operates on integers not just reals
func mod(n, d int) int {
	m := n % d
	if m < 0 {
		return m + d
	} else {
		return m
	}
}
