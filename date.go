package godate

import (
	"fmt"
	"time"

	"github.com/Nemagu/godate/gregorian"
)

type Date struct {
	year  int
	month time.Month
	day   int
}

func NewDate(year int, month time.Month, day int) (Date, error) {
	if err := ValidateDate(year, month, day); err != nil {
		return Date{}, err
	}
	return Date{
		year:  year,
		month: month,
		day:   day,
	}, nil
}

func Today() Date {
	timeNow := time.Now()
	return Date{
		year:  timeNow.Year(),
		month: timeNow.Month(),
		day:   timeNow.Day(),
	}
}

func FromTime(t time.Time) Date {
	return Date{
		year:  t.Year(),
		month: t.Month(),
		day:   t.Day(),
	}
}

func ValidateDate(year int, month time.Month, day int) error {
	if year < 0 || year > 9999 {
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

func (d Date) Year() int {
	return d.year
}

func (d Date) Month() time.Month {
	return d.month
}

func (d Date) Day() int {
	return d.day
}

func (d Date) ToTime(loc *time.Location) time.Time {
	return time.Date(d.year, d.month, d.day, 0, 0, 0, 0, loc)
}

func (d *Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.year, d.month, d.day)
}
