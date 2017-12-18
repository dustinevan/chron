package chron

import (
	"time"

	"github.com/dustinevan/chron/dura"
)

type Micro struct {
	time.Time
}

func NewMicro(year int, month time.Month, day, hour, min, sec, micro int) Micro {
	return Micro{time.Date(year, month, day, hour, min, sec, micro*1000, time.UTC)}
}

func ThisMicro() Micro {
	return Now().AsMicro()
}

func MicroOf(t time.Time) Micro {
	return Micro{t.Truncate(time.Microsecond)}
}

func (m Micro) AsYear() Year           { return YearOf(m.Time) }
func (m Micro) AsMonth() Month         { return MonthOf(m.Time) }
func (m Micro) AsDay() Day             { return DayOf(m.Time) }
func (m Micro) AsHour() Hour           { return HourOf(m.Time) }
func (m Micro) AsMinute() Minute       { return MinuteOf(m.Time) }
func (m Micro) AsSecond() Second       { return SecondOf(m.Time) }
func (m Micro) AsMilli() Milli         { return MilliOf(m.Time) }
func (m Micro) AsMicro() Micro         { return MicroOf(m.Time) }
func (m Micro) AsTimeExact() TimeExact { return TimeOf(m.Time) }
func (m Micro) AsTime() time.Time      { return m.Time }

func (m Micro) Increment(l dura.Time) TimeExact {
	return TimeExact{m.AddDate(l.Years(), l.Months(), l.Days()).Add(l.Duration())}
}

func (m Micro) Decrement(l dura.Time) TimeExact {
	return TimeExact{m.AddDate(-1*l.Years(), -1*l.Months(), -1*l.Days()).Add(-1 * l.Duration())}
}

func (m Micro) AddN(n int) Micro {
	return Micro{m.Add(time.Duration(int(time.Microsecond) * n))}
}

// span.Time implementation
func (m Micro) Start() Time {
	return m.AsTimeExact()
}

func (m Micro) End() Time {
	return m.AddN(1).Decrement(dura.Nano)
}

func (m Micro) Contains(t Span) bool {
	return !m.Before(t) && !m.After(t)
}

func (m Micro) Before(t Span) bool {
	return m.End().AsTime().Before(t.Start().AsTime())
}

func (m Micro) After(t Span) bool {
	return m.Start().AsTime().After(t.End().AsTime())
}

func (m Micro) Duration() dura.Time {
	return dura.Micro
}
