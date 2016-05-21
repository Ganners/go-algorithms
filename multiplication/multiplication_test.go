package multiplication

import "testing"

func TestMultiply1(t *testing.T) {
	for a := 0; a <= 100; a += 1 {
		for b := 0; b <= 100; b += 1 {

			expected := a * b
			res := Multiply1(a, b)

			if res != expected {
				t.Errorf("[%d * %d] Result %d does not match expected %d", a, b, res, expected)
			}
		}
	}
}

func TestMultiply2(t *testing.T) {
	for a := 0; a <= 100; a += 1 {
		for b := 0; b <= 100; b += 1 {

			expected := a * b
			res := Multiply2(a, b)

			if res != expected {
				t.Errorf("[%d * %d] Result %d does not match expected %d", a, b, res, expected)
			}
		}
	}
}

// Compare recursion with iteration, iteration wins of course!
func BenchmarkMultiply1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply1(773, 103)
	}
}

func BenchmarkMultiply2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply2(773, 103)
	}
}
