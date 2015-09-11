package sum_lists

import (
	"reflect"
	"testing"
)

func TestNewIntLinkedList(t *testing.T) {

	fixture := &IntLinkedList{
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
		t.Errorf("Returned list %s does not match expected: %s", ll, fixture)
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
		t.Errorf("Returned list %s does not match expected: %s", result, c)
	}
}

func TestSumListsOrdered(t *testing.T) {

	a := NewIntLinkedList(6, 1, 7)
	// +
	b := NewIntLinkedList(2, 9, 5)
	// =
	c := NewIntLinkedList(2, 1, 9)

	result := SumListsOrdered(a, b)
	if !reflect.DeepEqual(result, c) {
		t.Errorf("Returned list %s does not match expected: %s", result, c)
	}
}
