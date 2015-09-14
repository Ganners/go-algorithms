package main

import (
	"errors"
	"log"
	"math"
)

// Memory:
// Checking a bit.       isset := ((integer >> bitPos) & 1) == 1
// Toggling a bit.       integer ^= 1 << bitPos
// Turning on a bit.     integer |= 1 << bitPos
// Clearing a bit.       integer &^= 1 << bitPos
// Setting bit to value. integer ^= (-value ^ integer) & (1 << bitPos)

func main() {

	// 0
	var integer int
	log.Println(integer)

	// 1
	integer = integer | (1)
	log.Println(1, integer)

	// The XOR operator (^) can be used to toggle a bit

	// 3
	integer ^= (1 << 1)
	log.Println(3, integer)

	// 2
	integer ^= (1 << 0)
	log.Println(2, integer)

	// Set the 5th bit on (should be 34)
	integer |= (1 << 5)
	log.Println(34, integer)

	// Check the 5th bit is turned on
	bit := ((integer >> 5) & 1) == 1
	log.Println(true, bit)

	// Another way to do it (bit mask with 5 turned on)
	bit = (integer)&(1<<5) != 0
	log.Println(true, bit)

	// To clear a bit, &^ it :-)

	// Set the 5th bit OFF (clearing it)
	integer &^= (1 << 5)
	log.Println(2, integer)

	// Check the 5th bit is turned on
	bit = ((integer >> 5) & 1) == 1
	log.Println(false, bit)

	// Twos compliment is equivilent to flipping all bits and adding 1

	// Twos compliment on number
	integer = 12
	integer = (^integer) + 1
	log.Println(-12, integer)

	// Setting nth bit to 1 or 0 (number ^= (-x ^ number) & (1 << n))
	// Conver back to 12
	integer = ^(integer - 1)

	// Flip the 1st bit to a 1
	integer ^= (-1 ^ integer) & (1 << 0)
	log.Println(13, integer)
}

// Insert M bits into N, starting at bit J and ending at bit I
func insertInto(n, m, i, j int) (int, error) {

	nLength := int(math.Floor(math.Log2(float64(n))))
	mLength := j - i // Assume we can't always use log2 of M

	if mLength+i > nLength {
		return 0, errors.New(
			"m cannot fit into n from position i")
	}

	// Create a mask of 1s
	nMask := (1 << (uint(mLength) + 1)) - 1

	// &not the 1s to set first to 0s
	return n&^(nMask<<uint(i)) | (m << uint(i)), nil
}

// Q. What does this function do?
func someFunction(a, b int) interface{} {
	return !(((a << 1) &^ (b << 1)) != 0)
}
