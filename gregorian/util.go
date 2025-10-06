package gregorian

import "time"

var daysInMonth = []int{
	0,
	31, // January
	28,
	31, // March
	30,
	31, // May
	30,
	31, // July
	31,
	30, // September
	31,
	30, // November
	31,
}

// DaysInYear returns the number of days in a given year.
func DaysInYear(year int) int {
	if isLeap(year) {
		return 366
	}
	return 365
}

// DaysInYear returns the number of days in a given month in given year.
func DaysInMonth(year int, month time.Month) int {
	if month == time.February && isLeap(year) {
		return 29
	}
	return daysInMonth[month]
}

func isLeap(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
