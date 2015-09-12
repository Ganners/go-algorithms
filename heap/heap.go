package heap

import (
	"fmt"
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

// Outputs the tree so it can be visualised
//
// Example output:
//
// Input: [4, 6, 7, 8, 9, 10, 11] <- 1
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

	arrayHeight := ((numLevels) * 3)

	outputStr := make([]string, arrayHeight)

	for key, val := range mh.data {

		insertRow := int(
			math.Floor(
				math.Log2(float64(key + 1))))

		levelFromBottom := numLevels - insertRow
		numSpacesBefore := int(math.Pow(2.0, float64(levelFromBottom)))

		if outputStr[insertRow*3] != "" {
			numSpacesBefore = (numSpacesBefore * 2) - maxDigitLength
		}

		outputStr[insertRow*3] += strings.Repeat(" ", numSpacesBefore)
		outputStr[insertRow*3] += leftPad2Len(fmt.Sprintf("%d", val), "0", maxDigitLength)
	}

	return strings.Join(outputStr, "\n")
}

// Finds the maximum length of the digits, I.e. 1023 = 4
func (mh MinHeap) maxDigitLength() int {

	return int(
		math.Ceil(
			math.Log10(
				float64(mh.data[len(mh.data)-1] + 1))))
}

// Returns the number of nodes that exist
func (mh MinHeap) Count() int {
	return len(mh.data)
}

// Returns the number of levels that exist
func (mh MinHeap) NumLevels() int {
	return int(math.Ceil(math.Log2(float64(len(mh.data)))))
}
