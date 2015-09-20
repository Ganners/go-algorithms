package bst_array

import "testing"

func TestInArray(t *testing.T) {

	haystack := []int{1, 3, 6, 9, 11, 24, 33, 100, 101}

	needle := 1
	found := InArray(needle, haystack)
	if !found {
		t.Errorf("Expected element %s to be found, it wasn't", needle)
	}

	needle = 101
	found = InArray(needle, haystack)
	if !found {
		t.Errorf("Expected element %s to be found, it wasn't", needle)
	}

	needle = 3
	found = InArray(needle, haystack)
	if !found {
		t.Errorf("Expected element %s to be found, it wasn't", needle)
	}

	needle = 100
	found = InArray(needle, haystack)
	if !found {
		t.Errorf("Expected element %s to be found, it wasn't", needle)
	}
}

func TestInArrayIterative(t *testing.T) {

	haystack := []int{1, 3, 6, 9, 11, 24, 33, 100, 101}

	needle := 1
	found := InArrayIterative(needle, haystack)
	if !found {
		t.Errorf("Expected element %s to be found, it wasn't", needle)
	}

	needle = 101
	found = InArrayIterative(needle, haystack)
	if !found {
		t.Errorf("Expected element %s to be found, it wasn't", needle)
	}

	needle = 3
	found = InArrayIterative(needle, haystack)
	if !found {
		t.Errorf("Expected element %s to be found, it wasn't", needle)
	}

	needle = 100
	found = InArrayIterative(needle, haystack)
	if !found {
		t.Errorf("Expected element %s to be found, it wasn't", needle)
	}
}
