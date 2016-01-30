package main

import (
	"container/list"
	"math"
)

// Efficient linear(ish) time lookup for the largest rectangle
func LargestRectangle(histogram []int) int {

	// Prepend and append 0 so we have a start value
	histogram = append(append([]int{0}, histogram...), 0)

	// Create a stack and start it with a 0. Note here is the stack will keep
	// track of the histogram indexes, not the values
	stack := list.New()
	stack.PushBack(0)
	max := 0

	// Loop through our histogram values
	for i, x := range histogram {

		// While the current value is less than the value at the stack key
		for x < histogram[stack.Back().Value.(int)] {

			// Pop the last key from the stack and use as the key
			y := histogram[stack.Remove(stack.Back()).(int)]

			// Set the maximum value to the largest of itself and the
			// next stack item multiplied by this stack item
			max = int(math.Max(
				float64(max),
				float64((i-1-stack.Back().Value.(int))*y)))
		}

		// Push this key into the stack
		stack.PushBack(i)
	}
	return max
}
