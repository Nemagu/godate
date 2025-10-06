package godate

// Compare compares the date d with other.
// If t is before other, it returns -1;
// if t is after other, it returns +1;
// if they're the same, it returns 0.
func (d Date) Compare(other Date) int {
	switch {
	case d < other:
		return -1
	case d > other:
		return 1
	default:
		return 0
	}
}

// Equal reports whether d and other represent the same date.
func (d Date) Equal(other Date) bool {
	return d == other
}

// Before reports whether the date d is before other.
func (d Date) Before(other Date) bool {
	return d < other
}

// After reports whether the date d is after other.
func (d Date) After(other Date) bool {
	return d > other
}

// Sub returns the duration d-other.
func (d Date) Sub(other Date) Duration {
	return Duration(d) - Duration(other)
}

// Add returns the date d+other.
func (d Date) Add(duration Duration) Date {
	switch {
	case duration == 0:
		return d
	case duration < 0:
		temp := Date(-duration)
		if d <= temp {
			return Date(0)
		}
		return d - temp
	default:
		sum := d + Date(duration)
		if sum > maxDays {
			return Date(maxDays)
		}
		return sum
	}
}
