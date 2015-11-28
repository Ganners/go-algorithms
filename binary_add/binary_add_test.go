package binary_add

import (
	"reflect"
	"testing"
)

func TestAddBinary3_12(t *testing.T) {
	// 3
	a := []int{1, 1, 0, 0}
	// 12
	b := []int{0, 0, 1, 1}
	// 15
	c := []int{1, 1, 1, 1, 0}

	if r := AddBinary(a, b); !reflect.DeepEqual(r, c) {
		t.Error("Expected", c, "got", r)
	}
}

func TestAddBinary3_3(t *testing.T) {
	// 3
	a := []int{1, 1, 0, 0}
	// 3
	b := []int{1, 1, 0, 0}
	// 6
	c := []int{0, 1, 1, 0, 0}

	if r := AddBinary(a, b); !reflect.DeepEqual(r, c) {
		t.Error("Expected", c, "got", r)
	}
}

func TestAddBinary15_15(t *testing.T) {
	// 15
	a := []int{1, 1, 1, 1}
	// 15
	b := []int{1, 1, 1, 1}
	// 30
	c := []int{0, 1, 1, 1, 1}

	if r := AddBinary(a, b); !reflect.DeepEqual(r, c) {
		t.Error("Expected", c, "got", r)
	}
}
