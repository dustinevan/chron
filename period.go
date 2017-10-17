package chron

import (
	"github.com/dustinevan/chron/length"
	"fmt"
)

type Period interface {
	Contains(Time) bool
	Before() Time
	After() Time
	Len() Length
}

type Timespan struct {
	t1  Time
	t2  Time
	len Length
}

func NewSpan(start Time, len Length) *Timespan {
	return &Timespan{
		t1:  start,
		t2:  start.Increment(len),
		len: len,
	}
}

func (s *Timespan) Contains(t Time) bool {
	return (s.t1.AsTime().Before(t.AsTime()) && s.t2.AsTime().After(t.AsTime())) ||
		s.t1.AsTime().Equal(t.AsTime()) || s.t2.AsTime().Equal(t.AsTime())
}

func (s *Timespan) Before() Time {
	return s.t1.Decrement(length.Nano)
}

func (s *Timespan) After() Time {
	return s.t2.Increment(s.len)
}

func (s *Timespan) Len() Length {
	return s.len
}

func (s *Timespan) String() string {
	return fmt.Sprintf("t1:%s, t2:%s, len:%s", s.t1, s.t2, s.len)
}
