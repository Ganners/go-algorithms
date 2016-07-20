package main

import (
	"fmt"
	"math/rand"
)

func montyHall(iterations int, doSwitch bool) (numCorrect, numIncorrect int) {
	for i := 0; i < iterations; i++ {
		// Start with 3 rooms
		rooms := []int{1, 2, 3}
		winnerI := rand.Intn(len(rooms))
		winner := rooms[winnerI]

		// Pick a room
		chooseI := rand.Intn(len(rooms) - 1)
		choose := rooms[chooseI]
		remaining := append(rooms[:chooseI], rooms[chooseI+1:]...)

		// Host opens a door
		hostDoorI := 0
		if remaining[0] == winner {
			hostDoorI = 1
		}

		// Stay or switch?
		if doSwitch {
			remaining = append(rooms[:hostDoorI], rooms[hostDoorI+1:]...)
			choose = remaining[0]
		}

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
