package multiplication

// Iterative method where we grab the log of the smallest value
func Multiply2(a, b int) int {

	// Our short circuits, also make sure a is smallest
	if a == 0 || b == 0 {
		return 0
	}

	if a > b {
		b, a = a, b
	}

	if a == 1 {
		return b
	}

	// Start calculating c, build an adequately large stack
	//
	// We need to bring this into a stack as we have to work backwards
	// through it, essentially because we need to look at if it's odd
	// and apply the addition in reverse
	c := b
	stack := make([]int, 0, 8)

	for ; a > 1; a >>= 1 {
		stack = append(stack, a)
	}

	for i := len(stack) - 1; i >= 0; i-- {
		c <<= 1
		if (stack[i] & 1) == 1 {
			c += b
		}
	}

	return c
}

// Recursive method
func Multiply1(a, b int) int {

	// Answer if 0 if any inputs are 0
	if a == 0 || b == 0 {
		return 0
	}

	// If wrong, flip them
	if a > b {
		b, a = a, b
	}

	return multiplyRecursive(a, b)
}

func multiplyRecursive(a, b int) int {

	if a == 0 {
		return 0
	}

	if a == 1 {
		return b
	}

	c := multiplyRecursive(a>>1, b) << 1

	if (a & 1) == 1 {
		c += b
	}

	return c
}
