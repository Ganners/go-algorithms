package floodfill

import "sync"

// Point gets passed along the channels, it represents a starting
// position whose row needs filling
type point struct {
	y int
	x int
}

// Performs a flood fill
func FloodFill(input [][]int64, fill int64, x, y int) [][]int64 {

	// Validate
	lenY := len(input)
	if lenY-1 < y {
		return input
	}

	lenX := len(input[0])
	if lenX-1 < y {
		return input
	}

	// Effectively a queue of things to fill left->right
	fillChan := make(chan point)

	value := input[y][x]

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// Send initial point
	go func() { fillChan <- point{y, x} }()

	// Helper function to fill from -> to (left->right or right->left),
	// and also perform checks which will feed data into a channel
	//
	// Inlined for a simpler API, it uses variables outside of this
	fillFromTo := func(fromX, toX, atY int) {

		// Initial value for X
		x := fromX

		// Is it a forward or backward incrementer
		increment := true
		if fromX > toX {
			increment = false
		}

		// Optimize above and below with some flags
		checkUp := false
		checkDown := false

		for {
			if x == toX {
				break
			}

			// Stop if we hit a block
			if input[atY][x] != value {
				break
			}

			// Fill the colour
			input[atY][x] = fill

			// Check above
			if !checkUp && atY-1 >= 0 && input[atY-1][x] == value {
				wg.Add(1)
				checkUp = true
				go func(p point) { fillChan <- p }(point{atY - 1, x})
			} else {
				checkUp = false
			}

			// Check below
			if !checkDown && atY+1 < lenY && input[atY+1][x] == value {
				wg.Add(1)
				checkDown = true
				go func(p point) { fillChan <- p }(point{atY + 1, x})
			} else {
				checkDown = false
			}

			// Need to cater for decrementing or incrementing
			if increment {
				x++
			} else {
				x--
			}
		}
	}

	// Fill left and right whenever a point comes in
	go func() {
		for {
			select {
			case p := <-fillChan:
				fillFromTo(p.x-1, 0, p.y)    // Middle to left
				fillFromTo(p.x, lenX-1, p.y) // Middle to right
				wg.Done()
			}
		}
	}()

	wg.Wait()
	return input
}
