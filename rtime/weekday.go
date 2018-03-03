package rtime

import (
	"github.com/dustinevan/chron"
	"github.com/dustinevan/chron/dura"
)

//// chron.Relative represents a time or span that isn't definite until it is given
//// a time to relate to. e.g. Wednesdays, three hours ago, or a random second with the hour
//type Relative interface {
//	Next(chron.Time) (chron.Span, error)
//	Previous(chron.Time) (chron.Span, error)
//}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (w Weekday) Next(t chron.Time) chron.Day {
	return t.Increment(dura.Days(w.DaysUntil(t))).AsDay()
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
}
