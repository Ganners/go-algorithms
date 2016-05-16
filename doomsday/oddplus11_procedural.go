// The procedural version, for performance comparison
package doomsday

import "strconv"

func OddPlus11Procedural(d Date) string {

	d.dayOfWeek = (d.Year - (d.Year/1000)*1000)
	d.dayOfWeek = (d.Year - (d.Year/100)*100)

	if d.dayOfWeek%2 == 1 {
		d.dayOfWeek += 11
	}

	d.dayOfWeek /= 2

	if d.dayOfWeek%2 == 1 {
		d.dayOfWeek += 11
	}

	d.dayOfWeek = 7 - d.dayOfWeek

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

	d.dayOfWeek %= 7
	if d.dayOfWeek < 0 {
		d.dayOfWeek += 7
	}

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
	default:
	}

	return "Error: " + strconv.Itoa(d.dayOfWeek)
}
