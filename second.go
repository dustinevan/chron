package chron

import (
	"time"

	"github.com/dustinevan/chron/dura"
)

type Second struct {
	time.Time
}

func NewSecond(year int, month time.Month, day, hour, min, sec int) Second {
	return Second{time.Date(year, month, day, hour, min, sec, 0, time.UTC)}
}

func ThisSecond() Second {
	return Now().AsSecond()
}

func SecondOf(t time.Time) Second {
	return Second{t.Truncate(time.Second)}
}

func (s Second) AsYear() Year           { return YearOf(s.Time) }
func (s Second) AsMonth() Month         { return MonthOf(s.Time) }
func (s Second) AsDay() Day             { return DayOf(s.Time) }
func (s Second) AsHour() Hour           { return HourOf(s.Time) }
func (s Second) AsMinute() Minute       { return MinuteOf(s.Time) }
func (s Second) AsSecond() Second       { return SecondOf(s.Time) }
func (s Second) AsMilli() Milli         { return MilliOf(s.Time) }
func (s Second) AsMicro() Micro         { return MicroOf(s.Time) }
func (s Second) AsTimeExact() TimeExact { return TimeOf(s.Time) }
func (s Second) AsTime() time.Time      { return s.Time }

func (s Second) Increment(l dura.Time) TimeExact {
	return TimeExact{s.AddDate(l.Years(), l.Months(), l.Days()).Add(l.Duration())}
}

func (s Second) Decrement(l dura.Time) TimeExact {
	return TimeExact{s.AddDate(-1*l.Years(), -1*l.Months(), -1*l.Days()).Add(-1 * l.Duration())}
}

func (s Second) AddN(n int) Second {
	return Second{s.Add(time.Duration(int(time.Second) * n))}
}

/// span.Time implementation
func (s Second) Start() Time {
	return s.AsTimeExact()
}

func (s Second) End() Time {
	return s.AddN(1).Decrement(dura.Nano)
}

func (s Second) Contains(t Span) bool {
	return !s.Before(t) && !s.After(t)
}

func (s Second) Before(t Span) bool {
	return s.End().AsTime().Before(t.Start().AsTime())
}

func (s Second) After(t Span) bool {
	return s.Start().AsTime().After(t.End().AsTime())
}

func (s Second) Duration() dura.Time {
	return dura.Second
}
