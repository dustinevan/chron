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
func (m Micro) AsMicro() Micro         { return m }
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
func (m Micro) Start() TimeExact {
	return m.AsTimeExact()
}

func (m Micro) End() TimeExact {
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

func (m Micro) AddYears(y int) Micro {
	return m.Increment(dura.Years(y)).AsMicro()
}

func (m Micro) AddMonths(ms int) Micro {
	return m.Increment(dura.Months(ms)).AsMicro()
}

func (m Micro) AddDays(d int) Micro {
	return m.Increment(dura.Days(d)).AsMicro()
}

func (m Micro) AddHours(h int) Micro {
	return m.Increment(dura.Hours(h)).AsMicro()
}

func (m Micro) AddMinutes(ms int) Micro {
	return m.Increment(dura.Mins(ms)).AsMicro()
}

func (m Micro) AddSeconds(s int) Micro {
	return m.Increment(dura.Secs(s)).AsMicro()
}

func (m Micro) AddMillis(ms int) Micro {
	return m.Increment(dura.Millis(ms)).AsMicro()
}

func (m Micro) AddMicro(ms int) Micro {
	return m.AddN(ms)
}

func (m Micro) AddNano(n int) TimeExact {
	return m.AsTimeExact().AddN(n)
}
