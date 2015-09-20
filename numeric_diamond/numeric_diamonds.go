package numeric_diamond

import (
	"fmt"
	"log"
	"math"
	"strings"
)

type NumericDiamond struct {
	Diamond [][]int
	Size    int
}

func NewDiamond(number int) *NumericDiamond {

	diamond := make([][]int, number*2-1)

	for i := 0; i < number*number; i++ {
		row := (i) % (number)
		row += int(math.Floor(float64(i / number)))

		// Insert into row (prepend)
		diamond[row] = append([]int{i + 1}, diamond[row]...)
	}

	return &NumericDiamond{
		Diamond: diamond,
		Size:    number,
	}
}

func (n *NumericDiamond) String() string {

	largestNumber := n.Size * n.Size
	digitSize := getNumDigits(largestNumber + 1)
	diamondString := "\n"

	for _, row := range n.Diamond {
		// Big whitespace before
		diamondString += strings.Repeat(" ", ((len(n.Diamond)+1)/2-len(row))*digitSize)
		for _, digit := range row {

			// Pad number with whitespace
			diamondString += strings.Repeat(" ", (digitSize - getNumDigits(digit)))
			diamondString += fmt.Sprintf("%d", digit)

			// Print gap between next number
			diamondString += strings.Repeat(" ", digitSize)

		}
		diamondString += "\n"
	}

	return diamondString
}

func getNumDigits(number int) int {
	return int(math.Ceil(math.Log10(float64(number + 1))))
}

func main() {

	numericDiamond := NewDiamond(10)
	log.Println(numericDiamond)
}
