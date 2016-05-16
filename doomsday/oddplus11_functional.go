package doomsday

import "strconv"

func OddPlus11Functional(d Date) string {
	return toString(
		modulo7(
			driftClosestDay(
				driftDoomsday(
					subtractFrom7(
						oddPlus11(
							divide2(
								oddPlus11(
									twoYearDate(d)))))))))
}

// Grabs the tens and units from the year (last two numbers)
func twoYearDate(d Date) Date {
	d.dayOfWeek = (d.Year - (d.Year/1000)*1000)
	d.dayOfWeek = (d.Year - (d.Year/100)*100)
	return d
}

// If the number is odd, we add 11
func oddPlus11(d Date) Date {
	if d.dayOfWeek%2 == 1 {
		d.dayOfWeek += 11
	}
	return d
}

// Divides the number by two
func divide2(d Date) Date {
	d.dayOfWeek /= 2
	return d
}

// Performs a mod 7 on the number, in the positive set of numbers
func modulo7(d Date) Date {
	d.dayOfWeek %= 7
	if d.dayOfWeek < 0 {
		d.dayOfWeek += 7
	}
	return d
}

// Subtracts the number from 7
func subtractFrom7(d Date) Date {
	d.dayOfWeek = 7 - d.dayOfWeek
	return d
}

// Drifts to the doomsday
func driftDoomsday(d Date) Date {
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
	return d
}

// Drifts to the closest day of the month
func driftClosestDay(d Date) Date {
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
	return d
}

// Converts the dayOfWeek attached to the date and sends back a string
func toString(d Date) string {
	switch d.dayOfWeek {
	case 0:
		return "Sunday"
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	}

	return "Error: " + strconv.Itoa(d.dayOfWeek)
}
