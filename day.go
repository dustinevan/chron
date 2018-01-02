package chron

import (
	"time"

	"github.com/dustinevan/chron/dura"
)

type Day struct {
	time.Time
}

// Constructors
func NewDay(year int, month time.Month, day int) Day {
	return Day{time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
}

func Today() Day {
	return Now().AsDay()
}

func DayOf(time time.Time) Day {
	return NewDay(time.Year(), time.Month(), time.Day())
}

// chron.Time implementation
func (d Day) AsYear() Year           { return YearOf(d.Time) }
func (d Day) AsMonth() Month         { return MonthOf(d.Time) }
func (d Day) AsDay() Day             { return d }
func (d Day) AsHour() Hour           { return HourOf(d.Time) }
func (d Day) AsMinute() Minute       { return MinuteOf(d.Time) }
func (d Day) AsSecond() Second       { return SecondOf(d.Time) }
func (d Day) AsMilli() Milli         { return MilliOf(d.Time) }
func (d Day) AsMicro() Micro         { return MicroOf(d.Time) }
func (d Day) AsTimeExact() TimeExact { return TimeOf(d.Time) }
func (d Day) AsTime() time.Time      { return d.Time }

func (d Day) Increment(t dura.Time) TimeExact {
	return TimeExact{d.AddDate(t.Years(), t.Months(), t.Days()).Add(t.Duration())}
}

func (d Day) Decrement(t dura.Time) TimeExact {
	return TimeExact{d.AddDate(-1*t.Years(), -1*t.Months(), -1*t.Days()).Add(-1 * t.Duration())}
}

func (d Day) AddN(n int) Day {
	return Day{d.AddDate(0, 0, n)}
}

// span.Time implementation
func (d Day) Start() TimeExact {
	return d.AsTimeExact()
}

func (d Day) End() TimeExact {
	return d.AddN(1).Decrement(dura.Nano)
}

func (d Day) Contains(t Span) bool {
	return !d.Before(t) && !d.After(t)
}

func (d Day) Before(t Span) bool {
	return d.End().AsTime().Before(t.Start().AsTime())
}

func (d Day) After(t Span) bool {
	return d.Start().AsTime().After(t.End().AsTime())
}

func (d Day) Duration() dura.Time {
	return dura.Day
}

func (d Day) AddYears(y int) Day {
	return d.Increment(dura.Years(y)).AsDay()
}

// needs a global setting. i.e. july 31 - 1 month
func (d Day) AddMonths(m int) Day {
	return d.Increment(dura.Months(m)).AsDay()
}

func (d Day) AddDays(ds int) Day {
	return d.AddN(ds)
}

func (d Day) AddHours(h int) Hour {
	return d.AsHour().AddN(h)
}

func (d Day) AddMinutes(m int) Minute {
	return d.AsMinute().AddN(m)
}

func (d Day) AddSeconds(s int) Second {
	return d.AsSecond().AddN(s)
}

func (d Day) AddMillis(m int) Milli {
	return d.AsMilli().AddN(m)
}

func (d Day) AddMicro(m int) Micro {
	return d.AsMicro().AddN(m)
}

func (d Day) AddNano(n int) TimeExact {
	return d.AsTimeExact().AddN(n)
}
