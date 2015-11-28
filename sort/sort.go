package sort

const (
	ASC = iota
	DESC
)

// Implementation of insertion sort, with further code to flip
// the direct it sorts from. This has an Θ(n²) worst case
// running time.
func InsertionSort(a []int, dir int) []int {

	for j := 1; j < len(a); j++ {

		key := a[j]
		i := j - 1

		if dir == ASC {
			for ; i >= 0 && a[i] > key; i-- {
				a[i+1] = a[i]
			}
		}
		if dir == DESC {
			for ; i >= 0 && a[i] < key; i-- {
				a[i+1] = a[i]
			}
		}
		a[i+1] = key
	}

	return a
}

// Performs a linear search on an integer slice for an int.
// Returns either the found item or nil
func LinearSearch(a []int, search int) interface{} {

	for i := 0; i < len(a); i++ {

		if a[i] == search {
			return i
		}
	}
	return nil
}

// Merge sort, uses the Merge() function above. Recursively
// searches with divisions of two which, if len(A) were a
// square number would yield a worst case of Θ(n log(n)). This
// is a divide-and-conquer algorithm.
func MergeSort(A []int, p int, r int) []int {

	if p < r {
		q := (p + r) / 2
		A = MergeSort(A, p, q)
		A = MergeSort(A, q+1, r)
		A = merge(A, p, q, r)
	}

	return A
}

// Auxiliary procedure to MergeSort(), merges two pre-sorted
// sets from an array, delimited by q and ranging from p to r.
// Returns an array.
func merge(A []int, p int, q int, r int) []int {

	n1 := q - p + 1
	n2 := r - q

	L := make([]int, n1)
	R := make([]int, n2)

	for i := 0; i < n1; i++ {
		L[i] = A[p+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = A[q+j+1]
	}

	i := 0
	j := 0

	for k := p; k <= r; k++ {

		if len(L) > i && len(R) > j {
			if L[i] <= R[j] {

				A[k] = L[i]
				i++
			} else {

				A[k] = R[j]
				j++
			}
		} else if len(L) > i && len(R) <= j {

			A[k] = L[i]
			i++
		} else if len(L) <= i && len(R) > j {

			A[k] = R[j]
			j++
		}
	}

	return A
}

// Implementation of the Bubble Sort algorithm, really simple but really
// inefficient. Averages and worst cases Θ(n²)
func BubbleSort(A []int) []int {

	for i := 0; i < len(A); i++ {
		for j := len(A) - 1; j > i; j-- {

			if A[j] < A[j-1] {

				// Exchange variables
				A[j], A[j-1] = A[j-1], A[j]
			}
		}
	}
	return A
}
