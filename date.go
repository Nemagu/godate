package godate

import (
	"fmt"
	"time"

	"github.com/Nemagu/godate/v1/gregorian"
)

const (
	maxDays       = 3_652_060
	zeroOffset    = 719_162
	secondsPerDay = 24 * 60 * 60
)

var epoch = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)

// A Date represent one date from range from "0001-01-01" to "9999-12-31"
//
// A Date value can be used by multiple goroutines simultaneously except
// that the methods [Date.UnmarshalJSON] and [Date.UnmarshalText] are not concurrency-safe.
//
// Date instances can be compared using the [Date.Before], [Date.After] and [Date.Equal] methods.
// The [Date.Sub] method subtracts two instances, producing a [Duration].
// The [Date.Add] method adds a Date and a Duration, producing a Date.
//
// Representations of a Date value saved by the [Date.MarshalJSON] and [Date.MarshalText] methods.
type Date uint32

// ZeroDate returns Date instance which corresponds zero value Date - "0001-01-01" or 0.
func ZeroDate() Date {
	return Date(0)
}

// New returns Date instance and error from given year, month and day.
// If year was out of range [1, 9999]
// and month was out of range [1, 12]
// and day was more than days count in month function returned error.
func New(year int, month time.Month, day int) (Date, error) {
	if err := ValidateDate(year, month, day); err != nil {
		return Date(0), err
	}
	return fromTime(time.Date(year, month, day, 0, 0, 0, 0, time.UTC)), nil
}

// MustNew returns Date instance and error from given year, month and day.
// If year was out of range [1, 9999]
// and month was out of range [1, 12]
// and day was more than days count in month function called panic!
func MustNew(year int, month time.Month, day int) Date {
	if err := ValidateDate(year, month, day); err != nil {
		panic(err)
	}
	return fromTime(time.Date(year, month, day, 0, 0, 0, 0, time.UTC))
}

// FromTime returns Date which corresponds given time.
func FromTime(t time.Time) Date {
	return fromTime(t)
}

// FromString returns the Date instance which corresponds given string.
// It function use time.Parse with time.DateOnly("2006-01-02") layout string
// so if it returns err FromString returns this error.
func FromString(s string) (Date, error) {
	t, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return Date(0), err
	}
	return fromTime(t), nil
}

func fromTime(t time.Time) Date {
	if t.IsZero() {
		return Date(0)
	}
	_, offset := t.Zone()
	secs := t.Unix() + int64(offset)
	return Date(zeroOffset + secs/secondsPerDay)
}

// Today returns the Date instance based on your UTC offset.
func Today() Date {
	return fromTime(time.Now())
}

// TodayUTC returns the Date instance based on zero UTC offset.
func TodayUTC() Date {
	return fromTime(time.Now().UTC())
}

// ValidateDate validate given year, month and day for exists.
func ValidateDate(year int, month time.Month, day int) error {
	if year < 1 || year > 9999 {
		return fmt.Errorf("invalid year: %d", year)
	}
	if month < 1 || month > 12 {
		return fmt.Errorf("invalid month: %d", month)
	}
	if day < 1 || day > gregorian.DaysInMonth(year, month) {
		return fmt.Errorf("invalid day: %d", day)
	}
	return nil
}

// ToTime returns time.Time instance
// which corresponds Date with 0 minutes, 0 seconds etc and based on zero UTC offset.
func (d Date) ToTime() time.Time {
	return d.toTime()
}

func (d Date) toTime() time.Time {
	return epoch.AddDate(0, 0, int(d))
}

// Date returns (year, month, day) for Date instance.
func (d Date) Date() (int, time.Month, int) {
	return d.toTime().Date()
}

// Year returns year for Date instance.
func (d Date) Year() int {
	if d.IsZero() {
		return 1
	}
	return d.toTime().Year()
}

// Month returns month for Date instance.
func (d Date) Month() time.Month {
	if d.IsZero() {
		return time.January
	}
	return d.toTime().Month()
}

// Day returns day for Date instance.
func (d Date) Day() int {
	if d.IsZero() {
		return 1
	}
	return d.toTime().Day()
}

// IsZero reports whether d represents the zero Date("0001-01-01").
func (d Date) IsZero() bool {
	return d == 0
}

// String returns Date instance formatted using the ISO standard("YYYY-MM-DD")
func (d Date) String() string {
	if d.IsZero() {
		return "0001-01-01"
	}
	year, month, day := d.Date()
	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}
