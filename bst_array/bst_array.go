package bst_array

import "math"

// Performs binary sort on a sorted array, recursively so is both O(log
// n) in time and space
func InArray(needle int, haystack []int) bool {

	midPoint := int(math.Ceil(float64(len(haystack) / 2)))

	if len(haystack) == 0 {
		return false
	}

	if needle == haystack[midPoint] {
		return true
	}

	if needle < haystack[midPoint] {
		return InArray(needle, haystack[:midPoint])
	}

	if needle > haystack[midPoint] {
		return InArray(needle, haystack[midPoint+1:])
	}

	return false
}

// Performs binary sort on a sorted array iteratively so is both O(log
// n) in time and space
func InArrayIterative(needle int, haystack []int) bool {

	for {
		midPoint := int(math.Ceil(float64(len(haystack) / 2)))

		if len(haystack) == 0 {
			return false
		}

		if needle == haystack[midPoint] {
			return true
		}

		if needle < haystack[midPoint] {
			haystack = haystack[:midPoint]
			continue
		}

		if needle > haystack[midPoint] {
			haystack = haystack[midPoint+1:]
			continue
		}
	}
}
