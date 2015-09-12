package heap

import (
	"fmt"
	"math"
	"strings"
)

func padStringLeft(s string, padStr string, overallLen int) string {
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
//        04
//
//
//    06      07
//
//
//  08  09  10  11
//
func (mh MinHeap) String() string {

	if mh.Count() < 1 {
		return ""
	}

	// Declare array width and height
	maxDigitLength := mh.maxDigitLength()
	numLevels := mh.NumLevels()
	arrayHeight := (numLevels * 3) + 1

	outputStr := make([]string, arrayHeight)

	for key, val := range mh.data {

		insertRow := int(
			math.Floor(
				math.Log2(float64(key + 1))))

		levelFromBottom := numLevels - insertRow
		numSpacesBefore := int(math.Pow(2.0, float64(levelFromBottom)))

		insertIndex := insertRow * 3

		if outputStr[insertIndex] != "" {
			numSpacesBefore = (numSpacesBefore * 2) - maxDigitLength
		}

		outputStr[insertIndex] += strings.Repeat(" ", numSpacesBefore)
		outputStr[insertIndex] +=
			fmt.Sprintf("%d", val)
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

func (mh *MinHeap) swap(index1, index2 int) {

	mh.data[index1], mh.data[index2] = mh.data[index2], mh.data[index1]
}

func (mh *MinHeap) parent(i int) (int, int) {

	index := (i - 1) / 2
	value := mh.data[index]

	return index, value
}

// Shuffles up the integer until it's in the right place
func (mh *MinHeap) shuffleUp(i int) {

	val := mh.data[i]

	for {
		// Look at the parent
		pIndex, pVal := mh.parent(i)

		// If the parent is greater, swap
		if pVal > val {
			mh.swap(pIndex, i)
			i = pIndex
		} else {
			break
		}
	}
}

// The real algorithm stuff
func (mh *MinHeap) Push(x int) {

	mh.data = append(mh.data, x)
	mh.shuffleUp(mh.Count() - 1)
}
