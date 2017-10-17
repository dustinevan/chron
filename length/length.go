package length

import (
	"fmt"
	"time"
)

// Implementations of the chron.Length interface

type Duration struct {
	Year  int
	Month int
	Day   int
	Dur   time.Duration
}

func (d Duration) Mult(n int) Duration {
	return Duration{
		Year:  d.Year * n,
		Month: d.Month * n,
		Day:   d.Day * n,
		Dur:   time.Duration(int(d.Dur) * n),
	}
}

func (d Duration) Years() int {
	return d.Year
}

func (d Duration) Months() int {
	return d.Month
}

func (d Duration) Days() int {
	return d.Day
}

func (d Duration) Duration() time.Duration {
	return d.Dur
}

func (d Duration) String() string {
	return fmt.Sprintf("%syrs %smons %sdays %snanos", d.Year, d.Month, d.Day, d.Dur)
}

// TimeUnit represents a length of time for use in switch statements.
type Unit int

const (
	Century Unit = iota + 1
	Decade
	Year
	Quarter
	Month
	Week
	Day
	Hour
	Minute
	Second
	Milli
	Micro
	Nano
)

var units = []string{
	"Invalid Unit",
	"Century",
	"Decade",
	"Year",
	"Quarter",
	"Month",
	"Week",
	"Day",
	"Hour",
	"Minute",
	"Second",
	"Milli",
	"Micro",
	"Nano",
}

var durations = []Duration{
	Duration{},
	Duration{Year: 100},
	Duration{Year: 10},
	Duration{Year: 1},
	Duration{Month: 3},
	Duration{Month: 1},
	Duration{Day: 7},
	Duration{Day: 1},
	Duration{Dur: 60 * 60 * 1000 * 1000 * 1000},
	Duration{Dur: 60 * 1000 * 1000 * 1000},
	Duration{Dur: 1000 * 1000 * 1000},
	Duration{Dur: 1000 * 1000},
	Duration{Dur: 1000},
	Duration{Dur: 1},
}

/*var converters = []func(chron.Time) chron.Time{
	func(c chron.Time) chron.Time {

	}
}*/

func (u Unit) Mult(n int) Duration {
	return Duration{
		Year:  u.Years() * n,
		Month: u.Months() * n,
		Day:   u.Days() * n,
		Dur:   time.Duration(int(u.Duration()) * n),
	}
}

func (u Unit) Years() int {
	return durations[int(u)].Year
}

func (u Unit) Months() int {
	return durations[int(u)].Month
}

func (u Unit) Days() int {
	return durations[int(u)].Day
}

func (u Unit) Duration() time.Duration {
	return durations[int(u)].Dur
}

func (u Unit) String() string {
	return units[int(u)]
}
