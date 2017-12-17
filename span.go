package chron

import (
	"fmt"

	"github.com/dustinevan/chron/dura"
)

type Span interface {
	Start() TimeExact
	End() TimeExact
	Before(Time) bool
	After(Time) bool
	Contains(Time) bool
	Duration() dura.Time
}

type TimeSpan struct {
	start TimeExact
	end   TimeExact
	d     dura.Time
}

func NewTimeSpan(start TimeExact, d dura.Time) *TimeSpan {
	return &TimeSpan{
		start: start,
		end:   start.Increment(d).Decrement(dura.Nano),
		d:     d,
	}
}

func (s TimeSpan) Contains(t Span) bool {
	return !s.Before(t) && !s.After(t)
}

func (s TimeSpan) Before(t Span) bool {
	return s.End().AsTime().Before(t.Start().AsTime())
}

func (s TimeSpan) After(t Span) bool {
	return s.Start().AsTime().After(t.End().AsTime())
}

func (s TimeSpan) Duration() dura.Time {
	return s.d
}

func (s TimeSpan) Start() TimeExact {
	return s.start
}

func (s TimeSpan) End() TimeExact {
	return s.end
}

func (s TimeSpan) String() string {
	return fmt.Sprintf("start:%s, end:%s, len:%s", s.start, s.end, s.d)
}
