package rebase

import "testing"

// TestRebaseDecimalise will test out a number of bases
// with a large number of integers to make sure that they can
// convert back and forth
func TestRebaseDecimalise(t *testing.T) {
	for base := 2; base < 62; base++ {
		for i := 0; i < 1000; i += 5 {
			b := Rebase(i, base)
			d := Decimalise(b, base)
			if i != d {
				t.Errorf("%d does not match input %d", d, i)
			}
		}
	}
}
