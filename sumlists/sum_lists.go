package sum_lists

import (
	"fmt"
	"math"
)

type IntLink struct {
	Digit int
	Next  *IntLink
}

type IntLinkedList struct {
	Tail *IntLink
}

func (ll *IntLinkedList) String() string {

	link := ll.Tail
	outString := "Output: ("
	for link != nil {
		outString += fmt.Sprintf("%d,", link.Digit)
		link = link.Next
	}
	outString += ")\n"

	return outString
}

func NewIntLinkedList(inputs ...int) *IntLinkedList {

	ll := IntLinkedList{}
	var lastNode *IntLink

	for _, input := range inputs {

		if ll.Tail == nil {
			ll.Tail = &IntLink{
				Digit: input,
			}
			lastNode = ll.Tail
		} else {
			lastNode.Next = &IntLink{
				Digit: input,
			}
			lastNode = lastNode.Next
		}
	}
	return &ll
}

const (
	Forwards = iota
	Backwards
)

func intToDigitSlice(input int, direction int) []int {

	digitLength := int(math.Ceil(math.Log10(float64(input + 1))))
	digits := make([]int, digitLength)

	i := 0
	if direction == Backwards {
		i = digitLength - 1
	}

	for input > 0 {
		digits[i] = input % 10
		input = input / 10

		if direction == Forwards {
			i++
		} else {
			i--
		}
	}

	return digits
}

// Given a number of linked lists, this will sum them all together
// automagically
func SumLists(inputs ...*IntLinkedList) *IntLinkedList {

	total := 0

	for _, inputList := range inputs {

		link := inputList.Tail
		unit := 0

		for link != nil {
			toAdd := link.Digit * int(math.Pow(10.0, float64(unit)))
			total += toAdd
			unit++
			link = link.Next
		}
	}

	digits := intToDigitSlice(total, Backwards)
	return NewIntLinkedList(digits...)
}

// Given a number of linked lists, this will sum them all together
// automagically
func SumListsOrdered(inputs ...*IntLinkedList) *IntLinkedList {

	total := 0

	for _, inputList := range inputs {

		link := inputList.Tail
		unit := 9 // Should be sufficiently high
		localTotal := 0

		for link != nil {
			toAdd := link.Digit * int(math.Pow(10.0, float64(unit)))
			localTotal += toAdd
			unit--
			link = link.Next
		}

		// Then divide off how far over we are
		localTotal = localTotal / int(math.Pow(10.0, float64(unit+1)))
		total += localTotal
	}

	digits := intToDigitSlice(total, Forwards)
	return NewIntLinkedList(digits...)
}
