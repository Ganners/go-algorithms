// This is an experiment for calculating the day of the week, using the
// 'odd + 11' doomsday method in a purely pipelined way.
package doomsday

import (
	"fmt"
	"strconv"
)

// The input channel format and internal type for OddPlus11
type Date struct {
	Year  int
	Month int
	Day   int

	dayOfWeek int
}

// OddPlus11 will give you a channel which you can pass in a date, you'll then
// receive the day of the week for that date
func OddPlus11() (chan<- Date, <-chan string) {

	in := make(chan Date)

	pipe := twoYearDate(in)
	pipe = oddPlus11(pipe)
	pipe = divide2(pipe)
	pipe = oddPlus11(pipe)
	pipe = modulo7(pipe)
	pipe = subtractFrom7(pipe)
	pipe = driftDoomsday(pipe)
	pipe = driftClosestDay(pipe)
	pipe = modulo7(pipe)

	out := toString(pipe)

	return in, out
}

// Grabs the tens and units from the year (last two numbers)
func twoYearDate(in chan Date) chan Date {
	out := make(chan Date)
	go func() {
		for {
			d := <-in
			d.dayOfWeek = (d.Year - (d.Year/1000)*1000)
			d.dayOfWeek = (d.Year - (d.Year/100)*100)
			out <- d
		}
	}()
	return out
}

// If the number is odd, we add 11
func oddPlus11(in chan Date) chan Date {
	out := make(chan Date)
	go func() {
		for {
			d := <-in
			if d.dayOfWeek%2 == 1 {
				d.dayOfWeek += 11
			}
			out <- d
		}
	}()
	return out
}

// Divides the number by two
func divide2(in chan Date) chan Date {
	out := make(chan Date)
	go func() {
		for {
			d := <-in
			d.dayOfWeek /= 2
			out <- d
		}
	}()
	return out
}

// Performs a mod 7 on the number, in the positive set of numbers
func modulo7(in chan Date) chan Date {
	out := make(chan Date)
	go func() {
		for {
			d := <-in
			d.dayOfWeek %= 7
			if d.dayOfWeek < 0 {
				d.dayOfWeek += 7
			}
			out <- d
		}
	}()
	return out
}

// Subtracts the number from 7
func subtractFrom7(in chan Date) chan Date {
	out := make(chan Date)
	go func() {
		for {
			d := <-in
			d.dayOfWeek = 7 - d.dayOfWeek
			out <- d
		}
	}()
	return out
}

// Drifts to the doomsday
func driftDoomsday(in chan Date) chan Date {
	out := make(chan Date)
	go func() {
		for {
			d := <-in

			switch (d.Year / 200) % 4 {
			case 0:
				d.dayOfWeek += 5
			case 1:
				d.dayOfWeek += 3
			case 2:
				d.dayOfWeek += 2
			case 3:
				d.dayOfWeek += 6
			}

			out <- d
		}
	}()
	return out
}

// Drifts to the closest day of the month
func driftClosestDay(in chan Date) chan Date {
	out := make(chan Date)
	go func() {
		for {
			d := <-in
			isLeap := (d.Year % 4) == 0
			doomsday := 0

			switch d.Month {
			case 1:
				if isLeap {
					doomsday = 4
				} else {
					doomsday = 3
				}
			case 2:
				if isLeap {
					doomsday = 1
				} else {
					doomsday = 0
				}
			case 3:
				doomsday = 28
			case 4:
				doomsday = 4
			case 5:
				doomsday = 9
			case 6:
				doomsday = 6
			case 7:
				doomsday = 11
			case 8:
				doomsday = 8
			case 9:
				doomsday = 5
			case 10:
				doomsday = 10
			case 11:
				doomsday = 7
			case 12:
				doomsday = 12
			}

			d.dayOfWeek += d.Day - doomsday
			out <- d
		}
	}()
	return out
}

// Converts the dayOfWeek attached to the date and sends back a string
func toString(in chan Date) chan string {
	out := make(chan string)
	go func() {
		for {
			d := <-in

			switch d.dayOfWeek {
			case 0:
				out <- "Sunday"
			case 1:
				out <- "Monday"
			case 2:
				out <- "Tuesday"
			case 3:
				out <- "Wednesday"
			case 4:
				out <- "Thursday"
			case 5:
				out <- "Friday"
			case 6:
				out <- "Saturday"
			default:
				out <- "Error: " + strconv.Itoa(d.dayOfWeek)
			}
		}
	}()
	return out
}

// This can be used in the pipeline to spit out some debug information
func debug(in chan Date) chan Date {
	out := make(chan Date)
	go func() {
		for {
			d := <-in
			fmt.Printf("DEBUG: %+v\n", d)
			out <- d
		}
	}()
	return out
}
