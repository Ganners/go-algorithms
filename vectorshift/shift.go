package vectorshift

// Shifts a vector in O(N) with 1 space complexity
func Shift(N []int, x int) []int {

	ӏNӏ := len(N)

	// Can short circuit if there will be no shifting
	if x == ӏNӏ || (x > ӏNӏ && mod(ӏNӏ, x) == 0) {
		return N
	}

	lastI := 0
	lastV := N[lastI]

	for {
		shift := mod(lastI+x, ӏNӏ)
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

// Using Go's slicing internals, we can compute this in constant time
// (much more efficient, computes in 40% of the time of Shift)
//
// It does use more memory, however. Append will allocate
func Shift2(N []int, x int) []int {
	ӏNӏ := len(N)
	split := ӏNӏ - mod(x, ӏNӏ)
	return append(N[split:], N[:split]...)
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
