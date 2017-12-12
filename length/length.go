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

func Years(y int) Duration {
	return Duration{Year: y}
}

func Months(m int) Duration {
	return Duration{Month: m}
}

func Days(d int) Duration {
	return Duration{Day: d}
}

func Hours(h int) Duration {
	return Duration{Dur: time.Duration(int(time.Hour) * h)}
}

func Mins(m int) Duration {
	return Duration{Dur: time.Duration(int(time.Minute) * m)}
}

func Secs(s int) Duration {
	return Duration{Dur: time.Duration(int(time.Second) * s)}
}

func Millis(m int) Duration {
	return Duration{Dur: time.Duration(int(time.Millisecond) * m)}
}

func Micros(m int) Duration {
	return Duration{Dur: time.Duration(int(time.Microsecond) * m)}
}

func Nanos(n int) Duration {
	return Duration{Dur: time.Duration(int(time.Nanosecond) * n)}
}

func (d Duration) Mult(n int) Duration {
	return Duration{
		Year:  d.Year * n,
		Month: d.Month * n,
		Day:   d.Day * n,
		Dur:   time.Duration(int(d.Dur) * n),
	}
}

func Sum(durs ...Duration) Duration {
	d := Duration{}
	for _, dur := range durs {
		d.Year = d.Year + dur.Year
		d.Month = d.Month + dur.Month
		d.Day = d.Day + dur.Day
		d.Dur = time.Duration(int(d.Dur) + int(dur.Dur))
	}
	return d
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
