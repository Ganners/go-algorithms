package sort

import (
	"reflect"
	"testing"
)

func TestInsertionSortAsc(t *testing.T) {

	var a = []int{5, 2, 4, 6, 1, 3}
	var sorted = []int{1, 2, 3, 4, 5, 6}
	if s := InsertionSort(a, ASC); !reflect.DeepEqual(s, sorted) {
		t.Error("Expected", sorted, "got", s)
	}

	// Excercise 2.1-1
	a = []int{31, 41, 59, 26, 41, 58}
	sorted = []int{26, 31, 41, 41, 58, 59}
	if s := InsertionSort(a, ASC); !reflect.DeepEqual(s, sorted) {
		t.Error("Expected", sorted, "got", s)
	}
}

func TestInsertionSortDesc(t *testing.T) {

	// Excercise 2.1-2
	var a = []int{5, 2, 4, 6, 1, 3}
	var sorted = []int{6, 5, 4, 3, 2, 1}
	if s := InsertionSort(a, DESC); !reflect.DeepEqual(s, sorted) {
		t.Error("Expected", sorted, "got", s)
	}
}

func TestLinearSearch(t *testing.T) {

	// Excercise 2.1-3
	var a = []int{12, 32, 14, 6, 18, 2, 2}

	if r := LinearSearch(a, 7); r != nil {
		t.Error("Expected nil", "got", r)
	}

	if r := LinearSearch(a, 2); r != 5 {
		t.Error("Expected 2", "got", r)
	}
}

// Test our merge function, which is an auxiliary procedure of our merge sort
func TestMerge(t *testing.T) {

	// Test 1
	A := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 4, 5, 7, 1, 2, 3, 6, 0, 0, 0, 0}
	R := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 2, 3, 4, 5, 6, 7, 0, 0, 0, 0}

	if r := merge(A, 9, 12, 16); !reflect.DeepEqual(r, R) {
		t.Error("Expected", R, "got", r)
	}

	// Test 2, with p - q and q+1 - r reversed
	A = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 6, 2, 4, 5, 7, 0, 0, 0, 0}
	R = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 2, 3, 4, 5, 6, 7, 0, 0, 0, 0}

	if r := merge(A, 9, 12, 16); !reflect.DeepEqual(r, R) {
		t.Error("Expected", R, "got", r)
	}
}

// Test our merge function, which is an auxiliary procedure of our merge sort
func TestMergeLengthThree(t *testing.T) {
	A := []int{6, 2, 3}
	R := []int{2, 3, 6}

	if r := merge(A, 0, 0, 2); !reflect.DeepEqual(r, R) {
		t.Error("Expected", R, "got", r)
	}
}

func TestMergeSort(t *testing.T) {

	A := []int{5, 9, 4, 7, 1, 5, 4, 5, 4, 3, 2, 11, 8}
	R := []int{1, 2, 3, 4, 4, 4, 5, 5, 5, 7, 8, 9, 11}

	if r := MergeSort(A, 0, len(A)-1); !reflect.DeepEqual(r, R) {
		t.Error("Expected", R, "got", r)
	}
}

func TestBubbleSort(t *testing.T) {

	A := []int{5, 9, 4, 7, 1, 5, 4, 5, 4, 3, 2, 11, 8}
	R := []int{1, 2, 3, 4, 4, 4, 5, 5, 5, 7, 8, 9, 11}

	if r := BubbleSort(A); !reflect.DeepEqual(r, R) {
		t.Error("Expected", R, "got", r)
	}
}
