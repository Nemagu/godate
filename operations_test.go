package godate

import (
	"fmt"
	"testing"
	"time"
)

func TestDate_Sub(t *testing.T) {
	cases := []struct {
		d     Date
		od    Date
		delta Duration
	}{
		{d: MustNew(2001, time.January, 31), od: MustNew(2001, time.January, 1), delta: 30},
		{d: MustNew(2000, time.March, 1), od: MustNew(2000, time.February, 1), delta: 29},
		{d: MustNew(2001, time.January, 1), od: MustNew(2001, time.January, 31), delta: -30},
		{d: MustNew(2001, time.January, 1), od: MustNew(2001, time.January, 1), delta: 0},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("test_date(%s)_sub_date(%s)", c.d, c.od), func(t *testing.T) {
			delta := c.d.Sub(c.od)
			if delta != c.delta {
				t.Errorf("expected delta(%d) not right(%d)", delta, c.delta)
			}
		})
	}
}
