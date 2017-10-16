package chron

import (
	"time"

	"github.com/dustinevan/time/chron/length"
)

type Month struct {
	time.Time
}

func NewMonth(year int, month time.Month) Month {
	return Month{time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)}
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

func (m Month) Increment(l Length) TimeExact {
	return TimeExact{m.AddDate(l.Years(), l.Months(), l.Days()).Add(l.Duration())}
}

func (m Month) Decrement(l Length) TimeExact {
	return TimeExact{m.AddDate(-1*l.Years(), -1*l.Months(), -1*l.Days()).Add(-1 * l.Duration())}
}

func (m Month) AddN(n int) Month {
	return Month{m.AddDate(0, n, 0)}
}

// Period Implementation
func (m Month) Contains(t TimeExact) bool {
	return t.Month() == m.Month()
}

func (m Month) Before() TimeExact {
	return m.AsTimeExact().Decrement(length.Nano)
}

func (m Month) After() TimeExact {
	return m.AsTimeExact().Increment(length.Month)
}

func (m Month) Len() Length {
	return length.Month
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
