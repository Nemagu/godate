package godate

import (
	"fmt"
	"testing"
)

func TestData_MarshalText(t *testing.T) {
	cases := []struct {
		d             Date
		expected      string
		expectedError bool
	}{
		{d: MustNew(2000, 4, 10), expected: "2000-04-10", expectedError: false},
		{d: MustNew(2, 4, 11), expected: "0002-04-11", expectedError: false},
		{d: MustNew(9999, 4, 10), expected: "9999-04-10", expectedError: false},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("test_marshal_text_date(%s)", c.d), func(t *testing.T) {
			res, err := c.d.MarshalText()
			if c.expectedError {
				if err == nil {
					t.Error("expected error")
				}
			} else {
				if err != nil {
					t.Errorf("not expected error: %s", err)
				}
				temp := string(res)
				if temp != c.expected {
					t.Errorf("res(%s) is not equal expect(%s)", temp, c.expected)
				}
			}
		})
	}
}

func TestData_MarshalJSON(t *testing.T) {
	cases := []struct {
		d             Date
		expected      string
		expectedError bool
	}{
		{d: MustNew(2000, 4, 10), expected: "\"2000-04-10\"", expectedError: false},
		{d: MustNew(2, 4, 11), expected: "\"0002-04-11\"", expectedError: false},
		{d: MustNew(9999, 4, 10), expected: "\"9999-04-10\"", expectedError: false},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("test_marshal_json_date(%s)", c.d), func(t *testing.T) {
			res, err := c.d.MarshalJSON()
			if c.expectedError {
				if err == nil {
					t.Error("expected error")
				}
			} else {
				if err != nil {
					t.Fatalf("not expected error: %s", err)
				}
				temp := string(res)
				if temp != c.expected {
					t.Errorf("res(%s) is not equal expect(%s)", temp, c.expected)
				}
			}
		})
	}
}

func TestData_UnmarshalText(t *testing.T) {
	cases := []struct {
		d             Date
		expected      Date
		expectedError bool
	}{
		{d: ZeroDate(), expected: MustNew(2000, 04, 10), expectedError: false},
		{d: ZeroDate(), expected: MustNew(0002, 04, 11), expectedError: false},
		{d: ZeroDate(), expected: MustNew(9999, 04, 10), expectedError: false},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("test_unmarshal_text_date(%s)", c.expected), func(t *testing.T) {
			err := c.d.UnmarshalText([]byte(c.expected.String()))
			if c.expectedError {
				if err == nil {
					t.Error("expected error")
				}
			} else {
				if err != nil {
					t.Errorf("not expected error: %s", err)
				}
				if !c.d.Equal(c.expected) {
					t.Errorf("res(%s) is not equal expect(%s)", c.d.String(), c.expected)
				}
			}
		})
	}
}

func TestData_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		d             Date
		expected      Date
		row           []byte
		expectedError bool
	}{
		{d: ZeroDate(), expected: MustNew(2000, 04, 10), row: []byte("\"2000-04-10\""), expectedError: false},
		{d: ZeroDate(), expected: MustNew(0002, 04, 11), row: []byte("\"0002-04-11\""), expectedError: false},
		{d: ZeroDate(), expected: MustNew(9999, 04, 10), row: []byte("\"9999-04-10\""), expectedError: false},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("test_unmarshal_text_date(%s)", c.expected), func(t *testing.T) {
			err := c.d.UnmarshalJSON(c.row)
			if c.expectedError {
				if err == nil {
					t.Error("expected error")
				}
			} else {
				if err != nil {
					t.Errorf("not expected error: %s", err)
				}
				if !c.d.Equal(c.expected) {
					t.Errorf("res(%s) is not equal expect(%s)", c.d.String(), c.expected)
				}
			}
		})
	}
}
