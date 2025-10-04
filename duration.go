package godate

type Duration int32

const Day = Duration(1)

func (d Duration) Days() int {
	return int(d)
}
