package dura

import (
	"fmt"
	"time"
)

type Duration interface {
	Years() int
	Months() int
	Days() int
	Duration() time.Duration
}

// Implementations of the chron.Length interface

type Interval struct {
	Year  int
	Month int
	Day   int
	Dur   time.Duration
}

func Years(y int) Interval {
	return Interval{Year: y}
}

func Months(m int) Interval {
	return Interval{Month: m}
}

func Days(d int) Interval {
	return Interval{Day: d}
}

func Hours(h int) Interval {
	return Interval{Dur: time.Duration(int(time.Hour) * h)}
}

func Mins(m int) Interval {
	return Interval{Dur: time.Duration(int(time.Minute) * m)}
}

func Secs(s int) Interval {
	return Interval{Dur: time.Duration(int(time.Second) * s)}
}

func Millis(m int) Interval {
	return Interval{Dur: time.Duration(int(time.Millisecond) * m)}
}

func Micros(m int) Interval {
	return Interval{Dur: time.Duration(int(time.Microsecond) * m)}
}

func Nanos(n int) Interval {
	return Interval{Dur: time.Duration(int(time.Nanosecond) * n)}
}

func (d Interval) Mult(n int) Interval {
	return Interval{
		Year:  d.Year * n,
		Month: d.Month * n,
		Day:   d.Day * n,
		Dur:   time.Duration(int(d.Dur) * n),
	}
}

func Sum(durs ...Interval) Interval {
	d := Interval{}
	for _, dur := range durs {
		d.Year = d.Year + dur.Year
		d.Month = d.Month + dur.Month
		d.Day = d.Day + dur.Day
		d.Dur = time.Duration(int(d.Dur) + int(dur.Dur))
	}
	return d
}

func (d Interval) Years() int {
	return d.Year
}

func (d Interval) Months() int {
	return d.Month
}

func (d Interval) Days() int {
	return d.Day
}

func (d Interval) Duration() time.Duration {
	return d.Dur
}

func (d Interval) String() string {
	return fmt.Sprintf("%vyrs %vmons %vdays %s", d.Year, d.Month, d.Day, d.Dur)
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

var durations = []Interval{
	Interval{},
	Interval{Year: 100},
	Interval{Year: 10},
	Interval{Year: 1},
	Interval{Month: 3},
	Interval{Month: 1},
	Interval{Day: 7},
	Interval{Day: 1},
	Interval{Dur: 60 * 60 * 1000 * 1000 * 1000},
	Interval{Dur: 60 * 1000 * 1000 * 1000},
	Interval{Dur: 1000 * 1000 * 1000},
	Interval{Dur: 1000 * 1000},
	Interval{Dur: 1000},
	Interval{Dur: 1},
}

/*var converters = []func(chron.Time) chron.Time{
	func(c chron.Time) chron.Time {

	}
}*/

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
