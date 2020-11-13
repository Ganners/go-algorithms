package sqrt

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	for _, tc := range []struct {
		x float64
		y float64
	}{
		{
			x: 4,
			y: 2,
		},
		{
			x: 4,
			y: 2,
		},
		{
			x: 9,
			y: 3,
		},
		{
			x: 25,
			y: 5,
		},
		{
			x: 100,
			y: 10,
		},
		{
			x: 45,
			y: math.Sqrt(45),
		},
		{
			x: 123456,
			y: math.Sqrt(123456),
		},
	} {
		if y := Sqrt(tc.x); math.Abs(y-tc.y) > 1e-2 {
			t.Errorf("expected sqrt for %.2f to be %.2f, got %.2f", tc.x, tc.y, y)
		}
	}
}
