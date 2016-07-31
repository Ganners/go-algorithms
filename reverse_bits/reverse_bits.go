package reverse_bits

import "math"

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

// Constant time bitwise reverse in int64 space
//
// The multiply operation creates five separate copies of the 8-bit byte
// pattern to fan-out into a 64-bit value. The AND operation selects the bits
// that are in the correct (reversed) positions, relative to each 10-bit groups
// of bits. The multiply and the AND operations copy the bits from the original
// byte so they each appear in only one of the 10-bit sets. The reversed
// positions of the bits from the original byte coincide with their relative
// positions within any 10-bit set. The last step, which involves modulus
// division by 2^10 - 1, has the effect of merging together each set of 10 bits
// (from positions 0-9, 10-19, 20-29, ...) in the 64-bit value. They do not
// overlap, so the addition steps underlying the modulus division behave like
// or operations.
//
// This method was attributed to Rich Schroeppel in the Programming Hacks
// section of Beeler, M., Gosper, R. W., and Schroeppel, R. HAKMEM. MIT AI Memo
// 239, Feb. 29, 1972.
func ReverseBitsBitwise(bits uint8) uint8 {
	//    02 02 02 02 02
	// 01 08 84 42 20 10
	b := (int64(bits) * 0x0202020202 & 0x010884422010) % 1023
	return uint8(b)
}
