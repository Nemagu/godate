package godate

import "testing"

func TestDuration(t *testing.T) {
	cases := []struct {
		duration int32
		days     int
	}{
		{duration: 0, days: 0},
		{duration: 1, days: 1},
		{duration: 2, days: 2},
		{duration: 3, days: 3},
	}
	for _, c := range cases {
		t.Run("test_days_in_year", func(t *testing.T) {
			res := Duration(c.duration).Days()
			if res != c.days {
				t.Errorf("invalid days count: res = %d, exp = %d", res, c.days)
			}
		})
	}
}
