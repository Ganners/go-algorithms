package avx

import (
	"testing"
)

func TestSum(t *testing.T) {
	f := []int{1, 2, 3, 4, 5}
	sum := Sum(f)
	if sum != 15 {
		t.Errorf("sum %v does not match expected %v", sum, 15)
	}

	f = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum = Sum(f)
	if sum != 45 {
		t.Errorf("sum %v does not match expected %v", sum, 45)
	}
}

func TestSumFloat(t *testing.T) {
	f := []float64{1, 2, 3, 4, 5}
	sum := SumFloat(f)
	if sum != 15 {
		t.Errorf("sum %v does not match expected %v", sum, 15)
	}

	f = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum = SumFloat(f)
	if sum != 45 {
		t.Errorf("sum %v does not match expected %v", sum, 45)
	}
}

func TestSumFloatAVX(t *testing.T) {
	f := []float64{1, 2, 3, 4, 5}
	sum := SumFloatAVX(f)
	if sum != 15 {
		t.Errorf("sum %v does not match expected %v", sum, 15)
	}

	f = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum = SumFloatAVX(f)
	if sum != 45 {
		t.Errorf("sum %v does not match expected %v", sum, 45)
	}
}
