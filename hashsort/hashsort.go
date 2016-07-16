package hashsort

import "math"

// SuperHash just works out a column and row, this functionality is
// inlined in the HashSort function
func SuperHash(key, n int) (int, int) {
	square := int(math.Ceil(math.Sqrt(float64(n))))
	col := key % square
	row := key / square
	return row, col
}

// HashSort operates theoretically in N time and 2 space, it doesn't
// replace but generates a new array (2D to deal with duplicates).
func HashSort(in []int, max int) []int {
	width := int(math.Ceil(math.Sqrt(float64(max))))
	hashArray := make([][]int, width*width+width)
	flatHashArray := make([]int, 0, len(hashArray))

	for _, v := range in {
		col := v % width
		row := v / width
		i := width*row + col
		hashArray[i] = append(hashArray[i], v)
	}

	for _, v := range hashArray {
		for _, v2 := range v {
			flatHashArray = append(flatHashArray, v2)
		}
	}

	return flatHashArray
}
