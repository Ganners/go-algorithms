package polynomial

import "testing"

func TestInterpolatingPolynomial(t *testing.T) {
	for _, testCase := range []struct {
		x  float64
		y  float64
		xs []float64
		ys []float64
	}{
		{
			x:  -3,
			y:  2,
			xs: []float64{2, 7},
			ys: []float64{3, 4},
		},
		{
			x:  12,
			y:  5,
			xs: []float64{2, 7},
			ys: []float64{3, 4},
		},
	} {
		y := InterpolatingPolynomial(testCase.x, testCase.xs, testCase.ys)
		if y != testCase.y {
			t.Errorf("expected polynomial output to be %f, was %f", testCase.y, y)
		}
	}

}
