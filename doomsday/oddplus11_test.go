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

	// Build the pipelined OddPlus11 outside of the loop
	inputChan, answerChan := OddPlus11Pipelined()

	for _, test := range testCases {

		// Test - send something in and way for something to come out
		{
			inputChan <- test.in
			answer := <-answerChan

			if answer != test.out {
				t.Errorf("[%d-%d-%d] Answer %s does not match expected %s",
					test.in.Year, test.in.Month, test.in.Day,
					answer, test.out)
			}
		}

		// Test the procedural
		{
			out := OddPlus11Procedural(test.in)
			if out != test.out {
				t.Errorf("[%d-%d-%d] Answer %s does not match expected %s",
					test.in.Year, test.in.Month, test.in.Day,
					out, test.out)
			}
		}

		// Test the functional
		{
			out := OddPlus11Functional(test.in)
			if out != test.out {
				t.Errorf("[%d-%d-%d] Answer %s does not match expected %s",
					test.in.Year, test.in.Month, test.in.Day,
					out, test.out)
			}
		}
	}
}

func BenchmarkOddPlus1Pipelined(b *testing.B) {
	inputChan, answerChan := OddPlus11Pipelined()
	for i := 0; i < b.N; i++ {
		inputChan <- Date{2016, 3, 16, 0}
		<-answerChan
	}
}

func BenchmarkOddPlus1Procedural(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OddPlus11Procedural(Date{2016, 3, 16, 0})
	}
}

func BenchmarkOddPlus1Functional(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OddPlus11Functional(Date{2016, 3, 16, 0})
	}
}
