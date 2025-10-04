package godate

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

func (d Date) Equal(other Date) bool {
	return d == other
}

func (d Date) Before(other Date) bool {
	return d < other
}

func (d Date) After(other Date) bool {
	return d > other
}

func (d Date) Sub(other Date) Duration {
	return Duration(d) - Duration(other)
}

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
