package search

func BinarySearch(needle int, haystack []int) int {

	if len(haystack) == 0 {
		return -1
	}

	mid := len(haystack) / 2

	if haystack[mid] > needle {
		return BinarySearch(needle, haystack[:mid])
	} else if haystack[mid] < needle {
		return BinarySearch(needle, haystack[mid+1:])
	} else {
		return needle
	}
}
