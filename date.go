package godate

import (
	"fmt"
	"time"

	"github.com/Nemagu/godate/gregorian"
)

const (
	maxDays       = 3_652_060
	zeroOffset    = 719_162
	secondsPerDay = 24 * 60 * 60
)

var epoch = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)

type Date uint32

func ZeroDate() Date {
	return Date(0)
}

func New(year int, month time.Month, day int) (Date, error) {
	if err := ValidateDate(year, month, day); err != nil {
		return Date(0), err
	}
	return fromTime(time.Date(year, month, day, 0, 0, 0, 0, time.UTC)), nil
}

func MustNew(year int, month time.Month, day int) Date {
	if err := ValidateDate(year, month, day); err != nil {
		panic(err)
	}
	return fromTime(time.Date(year, month, day, 0, 0, 0, 0, time.UTC))
}

func FromTime(t time.Time) Date {
	return fromTime(t)
}

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

func Today() Date {
	return fromTime(time.Now())
}

func TodayUTC() Date {
	return fromTime(time.Now().UTC())
}

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

func (d Date) ToTime(loc *time.Location) time.Time {
	return d.toTime()
}

func (d Date) toTime() time.Time {
	return epoch.AddDate(0, 0, int(d))
}

func (d Date) Date() (int, time.Month, int) {
	return d.toTime().Date()
}

func (d Date) Year() int {
	if d.IsZero() {
		return 1
	}
	return d.toTime().Year()
}

func (d Date) Month() time.Month {
	if d.IsZero() {
		return time.January
	}
	return d.toTime().Month()
}

func (d Date) Day() int {
	if d.IsZero() {
		return 1
	}
	return d.toTime().Day()
}

func (d Date) IsZero() bool {
	return d == 0
}

func (d Date) String() string {
	if d.IsZero() {
		return "0001-01-01"
	}
	year, month, day := d.Date()
	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}
