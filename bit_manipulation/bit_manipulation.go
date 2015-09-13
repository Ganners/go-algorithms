package main

import "log"

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
	integer ^= (1 << 5)
	log.Println(34, integer)

	// Check the 5th bit is turned on
	bit := ((integer >> 5) & 1) == 1
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
}
