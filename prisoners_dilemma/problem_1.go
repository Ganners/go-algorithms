package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	numPrisoners := 1000
	numHats := 1001

	// there are 1000 prisoners
	prisoners := make([]int, numPrisoners)

	// and 1001 hats
	hats := make([]int, numHats)
	for i := 0; i < numHats; i++ {
		hats[i] = i + 1
	}

	// the hats are randomly sorted and one is removed
	rand.Shuffle(len(hats), func(i, j int) { hats[i], hats[j] = hats[j], hats[i] })
	hats = hats[:len(hats)-1]

	// the prisoners are all assigned hats
	for i := range prisoners {
		prisoners[i] = hats[i]
	}

	correctlyGuessedBefore := []int{}
	discardedThroughStrategy := []int{}
	priorGuess := 0
	prisonersSurvived := 0
	prisonersDied := 0

	// the prisoners take turns, from left to right, to make a choice
	for i := range prisoners {
		// we can look in front and see what we're missing from a set of 1001
		missing := make(map[int]struct{}, 1000)
		for j := 0; j < numHats; j++ {
			missing[j+1] = struct{}{}
		}
		for j := i + 1; j < len(prisoners); j++ {
			delete(missing, prisoners[j])
		}

		// remove prior guesses we've remembered
		for _, j := range correctlyGuessedBefore {
			delete(missing, j)
		}
		for _, j := range discardedThroughStrategy {
			delete(missing, j)
		}

		// we should be left with some number of options. I need to choose in
		// such a way that informs the person in front of what to choose
		missingSlice := make([]int, 0, len(missing))
		for j := range missing {
			missingSlice = append(missingSlice, j)
		}

		sort.Ints(missingSlice)
		guess := 0

		if i == 0 {
			// first prisoner does something a little different, they tell i + 1
			// their number which eliminates the remaining choice
			guess = prisoners[i+1]
		} else if i == 1 {

			// the second prisoner follows this strategy and can eliminate
			// everything
			guess = priorGuess

			// we can just discard all
			discardedThroughStrategy = append(discardedThroughStrategy, missingSlice...)
		} else {
			guess = missingSlice[len(missingSlice)-1]
		}

		priorGuess = guess

		// take a guess at the minimum
		if guess == prisoners[i] {
			correctlyGuessedBefore = append(correctlyGuessedBefore, guess)
			prisonersSurvived++
		} else {
			prisonersDied++
		}
	}

	fmt.Println(prisonersSurvived, prisonersDied)
}
