package reverse_bits

import "math"

const (
	inf = 129
)

// Reverses the bits on an int64
// Operates in O(N) time and O(1) space
func ReverseBits(bits uint8) uint8 {
	reversed := uint8(0)
	j := 0.0
	for i := 7.0; i >= 0; i-- {
		unit := uint8(math.Pow(2.0, i))
		if bits >= unit {
			reversed += uint8(math.Pow(2.0, j))
			bits -= unit
		}
		j++
	}
	return reversed
}
