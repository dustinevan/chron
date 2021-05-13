package rtime

import (
	"time"

	"github.com/dustinevan/chron"
)

// 0 is the last day of the month, 1 = 2nd to last day etc. rollover is possible
type EndOfMonth int

func (e EndOfMonth) Date(m chron.Month) chron.Day {
	return m.AddN(1).AsDay().AddN(int(e) - 1)
}

func LastWeekdayOfMonth(m chron.Month, w time.Weekday) chron.Day {
	nth := NewNthWeekDay(w, -1)
	return nth.Date(m.AddN(1).AsDay())
}

func LastWeekdayOfYear(y chron.Year, w time.Weekday) chron.Day {
	nth := NewNthWeekDay(w, -1)
	return nth.Date(y.AddN(1).AsDay())
}
