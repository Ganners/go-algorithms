package heap

import (
	"reflect"
	"strings"
	"testing"
)

var heap MinHeap = MinHeap{
	data: []int{4, 6, 7, 8, 9, 10, 11},
}

func TestCount(t *testing.T) {

	result := heap.Count()
	if result != 7 {
		t.Errorf("Expected to be 7, got %d", result)
	}
}

func TestMaxDigitLength(t *testing.T) {

	result := heap.maxDigitLength()
	if result != 2 {
		t.Errorf("Expected to be 2, got %d", result)
	}
}

func TestNumLevels(t *testing.T) {

	result := heap.NumLevels()
	if result != 3 {
		t.Errorf("Expected to be 3, got %d", result)
	}

	result = MinHeap{
		data: []int{4, 6, 7, 8},
	}.NumLevels()
	if result != 2 {
		t.Errorf("Expected to be 3, got %d", result)
	}
}

func TestMinHeapString(t *testing.T) {

	fixture := strings.Join([]string{
		"        4",
		"",
		"",
		"    6      7",
		"",
		"",
		"  8  9  10  11",
		"",
		"",
		"",
	}, "\n")

	heapOutput := heap.String()

	if heapOutput != fixture {

		t.Errorf(
			"Output did not match the expected.\nGot:\n\n%s\n\nExpected:\n\n%s\n",
			heapOutput,
			fixture)
	}
}

// Build a min heap by pushing stuff onto it and see if the data
// structure matches our expectations
func TestPushShuffle(t *testing.T) {

	heap := &MinHeap{}
	heap.Push(5)
	heap.Push(9)
	heap.Push(7)
	heap.Push(1)
	heap.Push(3)
	heap.Push(20)
	heap.Push(22)

	// Check it is what we expect here before we play with it some more
	fixture := MinHeap{
		data: []int{1, 3, 7, 9, 5, 20, 22},
	}

	if !reflect.DeepEqual(heap.data, fixture.data) {

		t.Errorf("Expected heap to look like \n%s\n, got \n%s\n",
			fixture.String(), heap.String())
	}

	// Check how 1 propogates up the heap
	heap.Push(1)

	fixture = MinHeap{
		data: []int{1, 1, 7, 3, 5, 20, 22, 9},
	}
	if !reflect.DeepEqual(heap.data, fixture.data) {

		t.Errorf("Expected heap to look like \n%s\n, got \n%s\n",
			fixture.String(), heap.String())
	}
}
