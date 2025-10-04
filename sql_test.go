package godate

import (
	"fmt"
	"testing"
	"time"
)

func TestDate_Scan(t *testing.T) {
	cases := []struct {
		row           any
		expected      Date
		expectedError bool
	}{
		{row: nil, expected: ZeroDate(), expectedError: false},
		{row: []byte("2000-01-01"), expected: MustNew(2000, 1, 1), expectedError: false},
		{row: "2000-01-01", expected: MustNew(2000, 1, 1), expectedError: false},
		{row: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), expected: MustNew(2000, 1, 1), expectedError: false},
		{row: uint32(zeroOffset), expected: ZeroDate(), expectedError: true},
		{row: uint64(zeroOffset), expected: ZeroDate(), expectedError: true},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("test_scan_type(%T)", c.row), func(t *testing.T) {
			temp := ZeroDate()
			err := temp.Scan(c.row)
			if c.expectedError {
				if err == nil {
					t.Error("expected error")
				}
			} else {
				if err != nil {
					t.Fatalf("not expected error: %s", err)
				}
				if !temp.Equal(c.expected) {
					t.Errorf("mut date(%s) is not equal expected date(%s)", temp, c.expected)
				}
			}
		})
	}
}

func TestDate_Value(t *testing.T) {
	cases := []struct {
		d        Date
		expected any
	}{
		{d: ZeroDate(), expected: nil},
		{d: MustNew(2000, 1, 1), expected: "2000-01-01"},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("test_value_date(%s)", c.d), func(t *testing.T) {
			value, err := c.d.Value()
			if err != nil {
				t.Fatalf("not expected error: %s", err)
			}
			if value != c.expected {
				t.Errorf("value(%v) is not equal expected value(%v)", value, c.expected)
			}
		})
	}
}
