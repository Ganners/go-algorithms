package fill_water

// This version of the function operates in about O(2N)
func FillWaterCount(env []int) int {

	if len(env) <= 2 {
		return 0
	}

	env = append([]int{0}, env...)
	total := 0
	fromLargest := 1
	toLargest := 0

	// While we haven't looked at every row
	for fromLargest != len(env)-1 {

		// Forward lookup
		for i := fromLargest + 1; i < len(env); i++ {
			if env[i] > env[toLargest] {
				toLargest = i
			}
			if env[toLargest] > env[fromLargest] {
				break
			}
		}

		// Calculate the minimum
		min := env[toLargest]
		if env[fromLargest] < env[toLargest] {
			min = env[fromLargest]
		}

		// Loop back and add to total
		for i := fromLargest + 1; i < toLargest; i++ {
			total += min - env[i]
		}

		fromLargest, toLargest = toLargest, 0
	}

	return total
}
