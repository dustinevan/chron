package chron

import (
	"fmt"

	"github.com/dustinevan/chron/dura"
	"time"
)

type Span interface {
	Start() Chron
	End() Chron
	Duration() dura.Time
	Comparer
}

type Comparer interface {
	Before(Span) bool
	After(Span) bool
	Contains(Span) bool
}

type Interval struct {
	start Chron
	end   Chron
	d     dura.Time
}

func NewInterval(start Chron, d dura.Time) *Interval {
	if d.Duration() == time.Nanosecond && d.Years() == 0 && d.Months() == 0 && d.Days() == 0 {
		return &Interval{start: start, end: start, d: dura.Nano}
	}
	return &Interval{
		start: start,
		end:   start.Increment(d).Decrement(dura.Nano),
		d:     d,
	}
}

func TimeRange(start, end Chron) *Interval {

}

func (s Interval) Contains(t Span) bool {
	return !s.Before(t) && !s.After(t)
}

func (s Interval) Before(t Span) bool {
	return s.End().AsTime().Before(t.Start().AsTime())
}

func (s Interval) After(t Span) bool {
	return s.Start().AsTime().After(t.End().AsTime())
}

func (s Interval) Duration() dura.Time {
	return s.d
}

func (s Interval) Start() Chron {
	return s.start
}

func (s Interval) End() Chron {
	return s.end
}

func (s Interval) String() string {
	return fmt.Sprintf("start:%s, end:%s, len:%s", s.start, s.end, s.d)
}


func Diff(start, end Time) dura.Duration {
	if start.AsTime().Equal(end.AsTime()) {
		return dura.Duration{}
	}
	if start.AsTime().After(end.AsTime()) {
		return Diff(end, start).Mult(-1)
	}
	s := start.AsChron()
	e := end.AsChron()

	yrs := e.AsYear().Year() - s.AsYear().Year()
	if s.AddYears(yrs).After(e) {
		yrs = yrs - 1
	}
	s = s.AddYears(yrs)

	dur := e.AsTime().Sub(s.AsTime())

	return dura.Duration{Yrs: yrs, Dur: dur}
}

func LeapDays(start, end Time) int {
	//TODO: i'm already looping, just count the days.

	if start.AsTime().After(end.AsTime()){
		return LeapDays(end, start)
	}
	s := start.AsChron()
	e := end.AsChron()
	yearsToCheck := make([]int, 0)
	if s.Before(s.AddMonths(2).AddDays(28)) {
		yearsToCheck = append(yearsToCheck, s.Year())
	}
	curr := s.AsYear().AddN(1)
	for curr.Year() < e.Year() {
		yearsToCheck = append(yearsToCheck, curr.Year())
		curr = curr.AddYears(1)
	}

}