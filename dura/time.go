package dura

import (
	"fmt"
	"time"
)

type Time interface {
	Years() int
	Months() int
	Days() int
	Duration() time.Duration
}

type Duration struct {
	Yrs  int
	Mons int
	Dys  int
	Dur  time.Duration
}

func NewDuration(year, month, day int, dur time.Duration) Duration {
	return Duration{year, month, day, dur}
}

func Years(y int) Duration {
	return Duration{Yrs: y}
}

func Months(m int) Duration {
	return Duration{Mons: m}
}

func Days(d int) Duration {
	return Duration{Dys: d}
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
		Yrs:  d.Yrs * n,
		Mons: d.Mons * n,
		Dys:  d.Dys * n,
		Dur:  time.Duration(int(d.Dur) * n),
	}
}

func Sum(durs ...Duration) Duration {
	d := Duration{}
	for _, dur := range durs {
		d.Yrs = d.Yrs + dur.Yrs
		d.Mons = d.Mons + dur.Mons
		d.Dys = d.Dys + dur.Dys
		d.Dur = time.Duration(int(d.Dur) + int(dur.Dur))
	}
	return d
}

func (d Duration) Years() int {
	return d.Yrs
}

func (d Duration) Months() int {
	return d.Mons
}

func (d Duration) Days() int {
	return d.Dys
}

func (d Duration) Duration() time.Duration {
	return d.Dur
}

func (d Duration) String() string {
	return fmt.Sprintf("%vy%vm%vd%s", d.Yrs, d.Mons, d.Dys, d.Dur)
}

// TimeUnit represents a length of time for use in switch statements.
type Unit int

const (
	Zero Unit = iota
	Century
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
	"Zero Unit",
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
	Duration{Yrs: 100},
	Duration{Yrs: 10},
	Duration{Yrs: 1},
	Duration{Mons: 3},
	Duration{Mons: 1},
	Duration{Dys: 7},
	Duration{Dys: 1},
	Duration{Dur: 60 * 60 * 1000 * 1000 * 1000},
	Duration{Dur: 60 * 1000 * 1000 * 1000},
	Duration{Dur: 1000 * 1000 * 1000},
	Duration{Dur: 1000 * 1000},
	Duration{Dur: 1000},
	Duration{Dur: 1},
}

func (u Unit) Years() int {
	return durations[int(u)].Yrs
}

func (u Unit) Months() int {
	return durations[int(u)].Mons
}

func (u Unit) Days() int {
	return durations[int(u)].Dys
}

func (u Unit) Duration() time.Duration {
	return durations[int(u)].Dur
}

func (u Unit) String() string {
	return units[int(u)]
}
