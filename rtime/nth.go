package rtime

import (
	"time"

	"github.com/dustinevan/chron"
)

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
