package chron

import (
	"fmt"

	"github.com/dustinevan/chron/dura"
)

type Span interface {
	Contains(Time) bool
	Before(Time) bool
	After(Time) bool
	Duration() dura.Duration
}

type Period struct {
	start Time
	end   Time
	d     dura.Duration
}

func NewSpan(start Time, d dura.Duration) *Timespan {
	return &Timespan{
		start: start,
		end:   start.Increment(len),
		d:     d,
	}
}

func (s *Timespan) Contains(t Time) bool {
	return !s.Before(t) && !s.After(t)
}

func (s *Timespan) Before(t Time) bool {
	return s.end.AsTime().Before(t.AsTime())
}

func (s *Timespan) After(t Time) bool {
	return s.start.AsTime().After(t.AsTime())
}

func (s *Timespan) Duration() dura.Duration {
	return s.d
}

func (s *Timespan) String() string {
	return fmt.Sprintf("start:%s, end:%s, len:%s", s.start, s.end, s.d)
}
