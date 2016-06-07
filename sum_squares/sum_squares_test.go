package sum_squares

import "testing"

func TestSumSquares(t *testing.T) {
	for i := 0; i < 100; i++ {
		brute := SumSquaresBrute(i)
		opt := SumSquares(i)

		if brute != opt {
			t.Errorf(
				"Brute force %d does not match optimized %d for %d",
				brute, opt, i,
			)
		}
	}
}
