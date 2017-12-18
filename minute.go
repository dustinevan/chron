package chron

import (
	"time"

	"github.com/dustinevan/chron/dura"
)

type Minute struct {
	time.Time
}

func NewMinute(year int, month time.Month, day, hour, min int) Minute {
	return Minute{time.Date(year, month, day, hour, min, 0, 0, time.UTC)}
}

func ThisMinute() Minute {
	return Now().AsMinute()
}

func MinuteOf(time time.Time) Minute {
	return NewMinute(time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute())
}

func (m Minute) AsYear() Year           { return YearOf(m.Time) }
func (m Minute) AsMonth() Month         { return MonthOf(m.Time) }
func (m Minute) AsDay() Day             { return DayOf(m.Time) }
func (m Minute) AsHour() Hour           { return HourOf(m.Time) }
func (m Minute) AsMinute() Minute       { return MinuteOf(m.Time) }
func (m Minute) AsSecond() Second       { return SecondOf(m.Time) }
func (m Minute) AsMilli() Milli         { return MilliOf(m.Time) }
func (m Minute) AsMicro() Micro         { return MicroOf(m.Time) }
func (m Minute) AsTimeExact() TimeExact { return TimeOf(m.Time) }
func (m Minute) AsTime() time.Time      { return m.Time }

func (m Minute) Increment(l dura.Time) TimeExact {
	return TimeExact{m.AddDate(l.Years(), l.Months(), l.Days()).Add(l.Duration())}
}

func (m Minute) Decrement(l dura.Time) TimeExact {
	return TimeExact{m.AddDate(-1*l.Years(), -1*l.Months(), -1*l.Days()).Add(-1 * l.Duration())}
}

func (m Minute) AddN(n int) Minute {
	return Minute{m.Add(time.Duration(int(time.Minute) * n))}
}

// span.Time implementation
func (m Minute) Start() Time {
	return m.AsTimeExact()
}

func (m Minute) End() Time {
	return m.AddN(1).Decrement(dura.Nano)
}

func (m Minute) Contains(t Span) bool {
	return !m.Before(t) && !m.After(t)
}

func (m Minute) Before(t Span) bool {
	return m.End().AsTime().Before(t.Start().AsTime())
}

func (m Minute) After(t Span) bool {
	return m.Start().AsTime().After(t.End().AsTime())
}

func (m Minute) Duration() dura.Time {
	return dura.Minute
}
