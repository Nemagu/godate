package gregorian

import (
	"fmt"
	"testing"
	"time"
)

func TestDaysInYear(t *testing.T) {
	cases := []struct {
		year int
		days int
	}{
		{year: 2001, days: 365},
		{year: 2002, days: 365},
		{year: 2003, days: 365},
		{year: 2004, days: 366},
		{year: 2100, days: 365},
		{year: 2400, days: 366},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("test_days_in_year(%d)", c.year), func(t *testing.T) {
			res := DaysInYear(c.year)
			if res != c.days {
				t.Errorf("invalid days count: res = %d, exp = %d", res, c.days)
			}
		})
	}
}

func TestDaysInMonth(t *testing.T) {
	cases := []struct {
		year  int
		month time.Month
		days  int
	}{
		{year: 2001, month: time.January, days: 31},
		{year: 2001, month: time.February, days: 28},
		{year: 2001, month: time.March, days: 31},
		{year: 2001, month: time.April, days: 30},
		{year: 2001, month: time.May, days: 31},
		{year: 2001, month: time.June, days: 30},
		{year: 2001, month: time.July, days: 31},
		{year: 2001, month: time.August, days: 31},
		{year: 2001, month: time.September, days: 30},
		{year: 2001, month: time.October, days: 31},
		{year: 2001, month: time.November, days: 30},
		{year: 2001, month: time.December, days: 31},
		{year: 2004, month: time.February, days: 29},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("test_days_in_year(%d)_month(%d)", c.year, c.month), func(t *testing.T) {
			res := DaysInMonth(c.year, c.month)
			if res != c.days {
				t.Errorf("invalid days count: res = %d, exp = %d", res, c.days)
			}
		})
	}
}
