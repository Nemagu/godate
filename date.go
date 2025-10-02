package godate

import (
	"time"
)

type Date uint64

func NewDate(year int, month time.Month, day int) (Date, error) {
	return Date(0), nil
}

func Now() Date {
	return Date(0)
}

func ZeroDate() Date {
	return Date(0)
}

func validateDate(count uint64) error {
	return nil
}
