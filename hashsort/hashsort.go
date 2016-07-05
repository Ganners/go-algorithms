package hashsort

import "math"

func SuperHash(key, n int) (int, int) {
	square := int(math.Ceil(math.Sqrt(float64(n))))
	col := key % square
	row := key / square
	return row, col
}

func HashSort(in []int) []int {
	size := len(in)
	width := int(math.Ceil(math.Sqrt(float64(size))))
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
