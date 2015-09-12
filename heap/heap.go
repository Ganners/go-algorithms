package heap

import (
	"fmt"
	"log"
	"math"
	"strings"
)

func leftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

type MinHeap struct {
	data []int
}

// Example output:
//
// input: [4, 6, 7, 8, 9, 10, 11] <- 1
//         ^  ^  ^  ^  ^   ^   ^
//         v__|__|  |  |   |   |  <- 2
//            v__|__|__|   |   |  <- 3
//               v_________|___|  <- 3
// Output:
//         04
//
//
//    06        07
//
//
// 08    09   10    11
//
func (mh MinHeap) String() string {

	// Declare array width and height
	maxDigitLength := mh.maxDigitLength()
	numLevels := mh.NumLevels()

	arrayHeight := (numLevels + 1) * 2
	arrayWidth := arrayHeight * maxDigitLength
	log.Println(arrayWidth)

	outputStr := make([]string, arrayHeight)

	for key, val := range mh.data {

		insertRow := int(
			math.Floor(
				math.Log2(float64(key + 1))))

		numSpaces := (numLevels - insertRow) * (maxDigitLength * 2)

		outputStr[insertRow*(numLevels)] += strings.Repeat(" ", numSpaces)
		outputStr[insertRow*(numLevels)] += leftPad2Len(fmt.Sprintf("%d", val), "0", maxDigitLength)
	}

	joined := strings.Join(outputStr, "\n")
	return joined
}

func (mh MinHeap) maxDigitLength() int {

	return int(
		math.Ceil(
			math.Log10(
				float64(mh.data[len(mh.data)-1] + 1))))
}

func (mh MinHeap) NumNodes() int {
	return len(mh.data)
}

func (mh MinHeap) NumLevels() int {
	return int(math.Ceil(float64(len(mh.data) / 2)))
}
