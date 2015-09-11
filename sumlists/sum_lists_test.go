package sum_lists

import (
	"reflect"
	"testing"
)

func TestNewIntLinkedList(t *testing.T) {

	fixture := &IntLinkedListSupervisor{
		Tail: &IntLink{
			Digit: 1,
			Next: &IntLink{
				Digit: 2,
				Next: &IntLink{
					Digit: 3,
				},
			},
		},
	}
	ll := NewIntLinkedList(1, 2, 3)

	if !reflect.DeepEqual(*fixture, *ll) {
		t.Errorf(
			"Returned linked list didn't match expected %+v", fixture)
	}
}

func TestSumLists(t *testing.T) {

	a := NewIntLinkedList(7, 1, 6)
	// +
	b := NewIntLinkedList(5, 9, 2)
	// =
	c := NewIntLinkedList(9, 1, 2)

	result := SumLists(a, b)
	if !reflect.DeepEqual(result, c) {
		t.Errorf("Returned list does not match expected: %+v", c)
	}
}
