package main

func removeDuplicates(in []int) ([]int, int) {

	duplicatesFound := 0

	for i := 0; i < len(in)-duplicatesFound; i++ {
		for j := i + 1; j < len(in)-duplicatesFound; j++ {
			if in[i] == in[j] {

				// Swap them round, does not maintain order
				duplicatesFound++
				in[j], in[len(in)-(duplicatesFound)] =
					in[len(in)-(duplicatesFound)], in[j]

				// Retest with the swapped value
				j--
			}
		}
	}

	// Cut it to the number of duplicates found
	return in[0 : len(in)-duplicatesFound], duplicatesFound
}
