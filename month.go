package chron

import (
	"time"

	"github.com/dustinevan/chron/dura"
)

type Month struct {
	time.Time
}

func NewMonth(year int, month time.Month) Month {
	return Month{time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)}
}

func ThisMonth() Month {
	return Now().AsMonth()
}

func MonthOf(time time.Time) Month {
	return NewMonth(time.Year(), time.Month())
}

func (m Month) AsYear() Year           { return YearOf(m.Time) }
func (m Month) AsMonth() Month         { return MonthOf(m.Time) }
func (m Month) AsDay() Day             { return DayOf(m.Time) }
func (m Month) AsHour() Hour           { return HourOf(m.Time) }
func (m Month) AsMinute() Minute       { return MinuteOf(m.Time) }
func (m Month) AsSecond() Second       { return SecondOf(m.Time) }
func (m Month) AsMilli() Milli         { return MilliOf(m.Time) }
func (m Month) AsMicro() Micro         { return MicroOf(m.Time) }
func (m Month) AsTimeExact() TimeExact { return TimeOf(m.Time) }
func (m Month) AsTime() time.Time      { return m.Time }

func (m Month) Increment(l dura.Time) TimeExact {
	return TimeExact{m.AddDate(l.Years(), l.Months(), l.Days()).Add(l.Duration())}
}

func (m Month) Decrement(l dura.Time) TimeExact {
	return TimeExact{m.AddDate(-1*l.Years(), -1*l.Months(), -1*l.Days()).Add(-1 * l.Duration())}
}

func (m Month) AddN(n int) Month {
	return Month{m.AddDate(0, n, 0)}
}

// span.Time implementation
func (m Month) Start() Time {
	return m.AsTimeExact()
}

func (m Month) End() Time {
	return m.AddN(1).Decrement(dura.Nano)
}

func (m Month) Contains(t Span) bool {
	return !m.Before(t) && !m.After(t)
}

func (m Month) Before(t Span) bool {
	return m.End().AsTime().Before(t.Start().AsTime())
}

func (m Month) After(t Span) bool {
	return m.Start().AsTime().After(t.End().AsTime())
}

func (m Month) Duration() dura.Time {
	return dura.Month
}

/*
func (m Month) AddMonth(m1 int) Month {
	return m.AsMonth()
}

func (m Month) AddDay(d int) Day {
	return NewDay(y.Year(), 1, d)
}

func (m Month) AddHour(h int) Hour {
	return NewHour(y.Year(), 1, 1, h)
}

func (m Month) AddMinute(m int) Minute {
}

func (m Month) AddSecond(s int) Second {
	return NewSecond(y.Year(), m, d)
}

func (m Month) AddMilli(m int) Milli {
	return NewMilli(y.Year)
}

func (m Month) AddMicro(m int) Micro {
	return NewMicro(y.Year)
}
*/
