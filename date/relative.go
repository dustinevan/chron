package date

import (
	"time"

	"github.com/dustinevan/time/chron"
)

// 0 is the last day of the month, 1 = 2nd to last day etc. rollover is possible
type EndOfMonth int

func (e EndOfMonth) Date(m chron.Month) chron.Day {
	return m.AddN(1).AsDay().AddN(int(e) - 1)
}

type NthWeekDay struct {
	// the weekday we're looking for
	weekday time.Weekday
	// the occurence number 0 is nothing, 1 is next weekday, -1 is prev weekday
	n int
}

func NewNthWeekDay(wday time.Weekday, n int) NthWeekDay {
	return NthWeekDay{weekday: wday, n: n}
}

func (nth NthWeekDay) Date(d chron.Day) chron.Day {
	if nth.n == 0 {
		return d
	}

	givenwd := d.Weekday()
	days := 0

	// negative case
	if nth.n < 0 {
		days += 7 * (nth.n + 1) // adding a negative number
		if nth.weekday >= givenwd {
			days -= 7 - int(nth.weekday) + int(givenwd)
		} else {
			days -= int(givenwd) - int(nth.weekday)
		}
		return d.AddN(days)
	}

	// positive case
	days += 7 * (nth.n - 1)
	if nth.weekday >= givenwd {
		days += int(nth.weekday) - int(givenwd)
	} else {
		days += 7 - int(givenwd) + int(nth.weekday)
	}

	return d.AddN(days)
}

func (nth NthWeekDay) OfMonth(m chron.Month) (chron.Day, bool) {
	result := nth.Date(m.AsDay())
	if m.Month() != result.Month() {
		return chron.ZeroValue().AsDay(), false
	}
	return result, true
}

func (nth NthWeekDay) OfYear(y chron.Year) (chron.Day, bool) {
	result := nth.Date(y.AsDay())
	if y.Year() != result.Year() {
		return chron.ZeroValue().AsDay(), false
	}
	return result, true
}

func LastWeekdayOfMonth(m chron.Month, w time.Weekday) chron.Day {
	nth := NewNthWeekDay(w, -1)
	return nth.Date(m.AddN(1).AsDay())
}

func LastWeekdayOfYear(y chron.Year, w time.Weekday) chron.Day {
	nth := NewNthWeekDay(w, -1)
	return nth.Date(y.AddN(1).AsDay())
}

// Example: The M-F date of the US observance of Christmas Day when Christmas falls on a weekend.
func ClosestNonWeekend(d chron.Day) chron.Day {
	if d.Weekday() == time.Saturday {
		return d.AddN(-1)
	}
	if d.Weekday() == time.Sunday {
		return d.AddN(1)
	}
	return d
}
