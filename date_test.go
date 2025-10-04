package godate

import (
	"fmt"
	"testing"
	"time"
)

func TestDate_New(t *testing.T) {
	cases := []struct {
		year        int
		month       time.Month
		day         int
		shouldError bool
	}{
		{year: 1, month: time.January, day: 1, shouldError: false},
		{year: 1999, month: time.January, day: 31, shouldError: false},
		{year: 1999, month: time.February, day: 1, shouldError: false},
		{year: 2000, month: time.February, day: 29, shouldError: false},
		{year: 2001, month: time.February, day: 1, shouldError: false},
		{year: 2001, month: time.February, day: 41, shouldError: true},
		{year: 0, month: time.January, day: 1, shouldError: true},
		{year: 10_000, month: time.January, day: 1, shouldError: true},
	}
	for _, c := range cases {
		t.Run(
			fmt.Sprintf("test_new_date(%04d-%02d-%02d)", c.year, c.month, c.day),
			func(t *testing.T) {
				_, err := New(c.year, c.month, c.day)
				if c.shouldError {
					if err == nil {
						t.Error("expected error")
					}
				} else {
					if err != nil {
						t.Errorf("not expected error: %s", err)
					}
				}
			},
		)
	}
}

func TestDate_fromTime(t *testing.T) {
	cases := []struct {
		t     time.Time
		year  int
		month time.Month
		day   int
	}{
		{t: time.Date(2025, time.January, 1, 2, 2, 2, 2, time.UTC), year: 2025, month: time.January, day: 1},
		{t: time.Date(2025, time.January, 20, 2, 2, 2, 2, time.UTC), year: 2025, month: time.January, day: 20},
		{t: time.Date(2028, time.February, 29, 2, 2, 2, 2, time.UTC), year: 2028, month: time.February, day: 29},
		{t: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), year: 1, month: time.January, day: 1},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("test_date_from_time(%s)", c.t), func(t *testing.T) {
			d := fromTime(c.t)
			if d.Year() != c.year {
				t.Errorf("date(%s) year(%d) not equal expected(%d)", d, d.Year(), c.year)
			}
			if d.Month() != c.month {
				t.Errorf("date(%s) month(%s) not equal expected(%s)", d, d.Month(), c.month)
			}
			if d.Day() != c.day {
				t.Errorf("date(%s) year(%d) not equal expected(%d)", d, d.Day(), c.day)
			}
		})
	}
}
