package main

import "log"

func main() {
	c := multiply(15, 30)
	log.Println(c)
}

func multiply(a, b int) int {

	// Defaults to a smaller than b
	smallest := a
	largest := b

	// If wrong, flip them
	if a > b {
		smallest, largest = largest, smallest
	}

	if smallest == 0 {
		return 0
	}

	return multiplyRecursive(smallest, largest)
}

func multiplyRecursive(smallest, largest int) int {

	if smallest == 0 {
		return 0
	}

	if smallest == 1 {
		return largest
	}

	halfProduct := multiplyRecursive(smallest>>1, largest)

	if (smallest & 1) == 1 {
		return (halfProduct << 1) + largest
	} else {
		return (halfProduct << 1)
	}
}
