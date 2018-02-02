package chron

import (
	"fmt"

	"github.com/dustinevan/chron/dura"
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
	return &Interval{
		start: start,
		end:   start.Increment(d).Decrement(dura.Nano),
		d:     d,
	}
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
