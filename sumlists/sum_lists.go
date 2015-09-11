package sum_lists

import (
	"fmt"
	"math"
)

type IntLink struct {
	Digit int
	Next  *IntLink
}

type IntLinkedListSupervisor struct {
	Tail   *IntLink
	Length int
}

func (ll *IntLinkedListSupervisor) String() string {

	link := ll.Tail
	outString := "Output: ("
	for link != nil {
		outString += fmt.Sprintf("%d,", link.Digit)
		link = link.Next
	}
	outString += ")\n"

	return outString
}

func NewIntLinkedList(inputs ...int) *IntLinkedListSupervisor {

	ll := IntLinkedListSupervisor{}
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

func SumLists(inputs ...*IntLinkedListSupervisor) *IntLinkedListSupervisor {

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

	digitLength := int(math.Ceil(math.Log10(float64(total + 1))))
	digits := make([]int, digitLength)
	i := digitLength - 1
	for total > 0 {
		digits[i] = total % 10
		total = total / 10
		i--
	}

	return NewIntLinkedList(digits...)
}
