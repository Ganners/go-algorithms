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
}

func TestMinHeapString(t *testing.T) {

	fixture := strings.Join([]string{
		"        04",
		"",
		"",
		"    06      07",
		"",
		"",
		"  08  09  10  11",
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

func TestPushShuffle(t *testing.T) {

	heap := &MinHeap{}
	heap.Push(5)
	heap.Push(10)
	heap.Push(7)
	heap.Push(1)

	fixture := MinHeap{
		data: []int{1, 5, 7, 10},
	}

	if !reflect.DeepEqual(heap, fixture) {

		t.Errorf("Expected heap to look like %s, got %s", fixture, heap)
	}
}
