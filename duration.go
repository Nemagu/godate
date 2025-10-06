package godate

// A Duration represents the elapsed date between two dates as an int32 days count.
type Duration int32

const Day = Duration(1)

// Days returns the duration as a integer number of days.
func (d Duration) Days() int {
	return int(d)
}
