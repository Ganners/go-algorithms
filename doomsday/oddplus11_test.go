package doomsday

import "testing"

func TestOddPlus11(t *testing.T) {

	testCases := []struct {
		in  Date
		out string
	}{
		{
			in: Date{
				Year:  1980,
				Month: 3,
				Day:   5,
			},
			out: "Wednesday",
		},
		{
			in: Date{
				Year:  1988,
				Month: 2,
				Day:   5,
			},
			out: "Friday",
		},
		{
			in: Date{
				Year:  1990,
				Month: 3,
				Day:   16,
			},
			out: "Friday",
		},
		{
			in: Date{
				Year:  2015,
				Month: 1,
				Day:   1,
			},
			out: "Thursday",
		},
		{
			in: Date{
				Year:  2016,
				Month: 5,
				Day:   15,
			},
			out: "Sunday",
		},
		{
			in: Date{
				Year:  2016,
				Month: 5,
				Day:   14,
			},
			out: "Saturday",
		},
		{
			in: Date{
				Year:  2016,
				Month: 6,
				Day:   10,
			},
			out: "Friday",
		},
		{
			in: Date{
				Year:  2016,
				Month: 12,
				Day:   25,
			},
			out: "Sunday",
		},
		{
			in: Date{
				Year:  2018,
				Month: 10,
				Day:   10,
			},
			out: "Wednesday",
		},
	}

	inputChan, answerChan := OddPlus11()

	for _, test := range testCases {

		// Test - send something in and way for something to come out
		inputChan <- test.in
		answer := <-answerChan

		if answer != test.out {
			t.Errorf("[%d-%d-%d] Answer %s does not match expected %s",
				test.in.Year, test.in.Month, test.in.Day,
				answer, test.out)
		}
	}
}
