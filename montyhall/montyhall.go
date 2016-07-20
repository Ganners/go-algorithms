package main

import (
	"fmt"
	"math/rand"
)

func montyHall(iterations int, doSwitch bool) (numCorrect, numIncorrect int) {
	for i := 0; i < iterations; i++ {
		// Start with 3 rooms
		rooms := []int{0, 1, 2}
		winner := rand.Intn(len(rooms))

		// Pick a room
		choose := rand.Intn(len(rooms) - 1)
		remaining := append(rooms[:choose], rooms[choose+1:]...)

		// Host opens a door, pick the first available
		hostDoor := 0
		for j := 0; j < len(remaining); j++ {
			if remaining[j] != winner {
				hostDoor = j
			}
		}

		// Stay or switch?
		if doSwitch {
			remaining = append(rooms[:hostDoor], rooms[hostDoor+1:]...)
			choose = remaining[0]
		}

		// Track
		if choose == winner {
			numCorrect++
		} else {
			numIncorrect++
		}
	}
	return
}

func main() {
	i := 10000
	scale := 100.0 / float64(i)
	{
		c, i := montyHall(i, false)
		p := scale * float64(c)
		fmt.Printf("No Switch: Correct/Incorrect: %d/%d (%f%%)\n", c, i, p)
	}
	{
		c, i := montyHall(i, true)
		p := scale * float64(c)
		fmt.Printf("Switch:    Correct/Incorrect: %d/%d (%f%%)\n", c, i, p)
	}
}
