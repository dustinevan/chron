package chron

import (
	"time"

	"github.com/dustinevan/time/chron/length"
)

type Milli struct {
	time.Time
}

func NewMilli(year int, month time.Month, day, hour, min, sec, milli int) Milli {
	return Milli{time.Date(year, month, day, hour, min, sec, milli*1000000, time.UTC)}
}

func MilliOf(t time.Time) Milli {
	return Milli{t.Truncate(time.Millisecond)}
}

func (m Milli) AsYear() Year           { return YearOf(m.Time) }
func (m Milli) AsMonth() Month         { return MonthOf(m.Time) }
func (m Milli) AsDay() Day             { return DayOf(m.Time) }
func (m Milli) AsHour() Hour           { return HourOf(m.Time) }
func (m Milli) AsMinute() Minute       { return MinuteOf(m.Time) }
func (m Milli) AsSecond() Second       { return SecondOf(m.Time) }
func (m Milli) AsMilli() Milli         { return MilliOf(m.Time) }
func (m Milli) AsMicro() Micro         { return MicroOf(m.Time) }
func (m Milli) AsTimeExact() TimeExact { return TimeOf(m.Time) }
func (m Milli) AsTime() time.Time      { return m.Time }

func (m Milli) Increment(l Length) TimeExact {
	return TimeExact{m.AddDate(l.Years(), l.Months(), l.Days()).Add(l.Duration())}
}

func (m Milli) Decrement(l Length) TimeExact {
	return TimeExact{m.AddDate(-1*l.Years(), -1*l.Months(), -1*l.Days()).Add(-1 * l.Duration())}
}

func (m Milli) AddN(n int) Milli {
	return Milli{m.Add(time.Duration(int(time.Millisecond) * n))}
}

// Period Implementation
func (m Milli) Contains(t TimeExact) bool {
	return (t.Nanosecond()/1000000)*1000000 == t.AsMilli().Nanosecond()
}

func (m Milli) Before() TimeExact {
	return m.AsExactTime().Decrement(length.Nano)
}

func (m Milli) After() TimeExact {
	return m.AsExactTime().Increment(length.Milli)
}

func (m Milli) Len() Length {
	return length.Milli
}
