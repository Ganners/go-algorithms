package duplicates

import (
	"math"
	"runtime"
	"sync"
)

// Stores the duplicated results, used in a struct so that the return
// arguments can be composed into a channel
type duplicateResult struct {
	output        []int
	numDuplicates int
}

// Aims to run the remove duplicates more parallellelely
func removeDuplicatesChunked(in []int) ([]int, int) {

	mu := new(sync.Mutex)
	out := make([]int, 0)
	numDuplicates := 0

	numCPU := runtime.NumCPU()
	chunkSize := int(math.Ceil(float64(len(in) / numCPU)))
	dupesChan := make(chan duplicateResult)

	// Chunk based on the number of available CPUs
	for i := 1; i <= numCPU; i++ {
		go func(i int, in []int) {

			sliceFrom := chunkSize * (i - 1)
			sliceTo := chunkSize * i
			chunk := make([]int, chunkSize)
			copy(chunk, in[sliceFrom:sliceTo])

			result, numDuplicates := removeDuplicates(chunk)
			dupesChan <- duplicateResult{
				result, numDuplicates,
			}
		}(i, in)
	}

	// Read them in and merge safely
	for chunk := 1; chunk <= numCPU; chunk++ {
		mu.Lock()

		res := <-dupesChan
		out = append(out, res.output...)
		numDuplicates += res.numDuplicates

		mu.Unlock()
	}

	// Run removeDuplicates on merged results
	out, dupeCount := removeDuplicates(out)
	numDuplicates += dupeCount

	return out, numDuplicates
}

// Removes the duplicates, returns back a slice that has been purified
// (order is not maintained) and the number of duplicates that were
// found
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
