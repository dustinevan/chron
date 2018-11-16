package rtime

import (
	"time"

	"github.com/dustinevan/chron"
)

type Relative func(chron.Time) (chron.Time, bool)

func Months(m ...time.Month) Relative {
	return func(exact chron.TimeExact) (chron.TimeExact, bool) {
		thisM := exact.Month()
		nextM := thisM
		for mon := range m {

		}
	}
}

type Month func(y chron.Time) chron.Month
type MonthDay func(y chron.Time) chron.Day
type DayOfMonth func(m chron.Time) chron.Day
type HourOfDay func(d chron.Time) chron.Hour
type TimeOfDay func(d chron.Time) chron.Minute
type MinOfHour func(d chron.Time) chron.Minute
type SecOfMin func(d chron.Time) chron.Second

func NewTimeOfDay(hr, min int) TimeOfDay {
	return
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
