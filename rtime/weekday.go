package rtime

import (
	"bufio"
	"github.com/dustinevan/chron"
	"github.com/dustinevan/chron/dura"
	"time"
)

type Weekday struct {
	current chron.Time
	day time.Weekday
}

func WeekdayOf(t chron.Time, w time.Weekday) Weekday {
	return Weekday{current:t, day:w}
}


func (w Weekday) Next() chron.Day {
	return .Increment(dura.Days(w.DaysUntil(t))).AsDay()
}

func (w Weekday) Prev(t chron.Time) chron.Day {
	return t.Decrement(dura.Days(w.DaysSince(t))).AsDay()
}

func (w Weekday) NthOccurence(t chron.Time, n int) chron.Day {
	if n > 0 {
		return t.Increment(dura.Days(w.DaysUntil(t) + (n-1)*7)).AsDay()
	}
	if n < 0 {
		return t.Decrement(dura.Days(w.DaysSince(t) + (n-1)*7)).AsDay()
	}
	return t.AsDay()
}

func (w Weekday) DaysUntil(t chron.Time) int {
	wd := int(w) - int(t.AsTime().Weekday())
	if wd > 0 {
		return wd
	}
	return 7 + wd
}

func (w Weekday) DaysSince(t chron.Time) int {
	wd := int(w) - int(t.AsTime().Weekday())
	if wd < 0 {
		return -wd
	}
	return 7 - wd
	bufio.NewScanner()
}
