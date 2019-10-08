package bitwise_add

import "testing"

func TestBitwiseAdd(t *testing.T) {
	for a := 0; a < 1000; a++ {
		for b := 0; b < 1000; b++ {
			expected := a + b
			c := BitwiseAdd(a, b)
			if c != expected {
				t.Fatalf("%d + %d should equal %d, got %d", a, b, expected, c)
			}
		}
	}
}
