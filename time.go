package chron

import (
	"time"
)

type Time interface {
	AsYear() Year
	AsMonth() Month
	AsDay() Day
	AsHour() Hour
	AsMinute() Minute
	AsSecond() Second
	AsMilli() Milli
	AsMicro() Micro
	AsTimeExact() TimeExact
	AsTime() time.Time

	Increment(length Length) TimeExact
	Decrement(length Length) TimeExact
}

type TimeExact struct {
	time.Time
}

func Now() TimeExact {
	return TimeOf(time.Now().In(time.UTC))
}

func NewTime(year int, month time.Month, day, hour, min, sec, nano int) TimeExact {
	return TimeExact{time.Date(year, time.Month(month), day, hour, min, sec, nano, time.UTC)}
}

func TimeOf(t time.Time) TimeExact {
	return TimeExact{t}
}

func (t TimeExact) AsYear() Year           { return YearOf(t.Time) }
func (t TimeExact) AsMonth() Month         { return MonthOf(t.Time) }
func (t TimeExact) AsDay() Day             { return DayOf(t.Time) }
func (t TimeExact) AsHour() Hour           { return HourOf(t.Time) }
func (t TimeExact) AsMinute() Minute       { return MinuteOf(t.Time) }
func (t TimeExact) AsSecond() Second       { return SecondOf(t.Time) }
func (t TimeExact) AsMilli() Milli         { return MilliOf(t.Time) }
func (t TimeExact) AsMicro() Micro         { return MicroOf(t.Time) }
func (t TimeExact) AsTimeExact() TimeExact { return TimeOf(t.Time) }
func (t TimeExact) AsTime() time.Time      { return t.Time }

func (t TimeExact) Increment(l Length) TimeExact {
	return TimeExact{t.AddDate(l.Years(), l.Months(), l.Days()).Add(l.Duration())}
}

func (t TimeExact) Decrement(l Length) TimeExact {
	return TimeExact{t.AddDate(-1*l.Years(), -1*l.Months(), -1*l.Days()).Add(-1 * l.Duration())}
}

func ZeroValue() TimeExact {
	return TimeExact{time.Time{}}
}
