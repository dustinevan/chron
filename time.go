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
	return NewYear(0).AsTimeExact()
}

// see: https://stackoverflow.com/questions/25065055/what-is-the-maximum-time-time-in-go
// and time.Unix() implementation
var unixToInternal = int64((1969*365 + 1969/4 - 1969/100 + 1969/400) * 24 * 60 * 60)
var max = time.Unix(1<<63-1-unixToInternal, 999999999).UTC()
var min = time.Unix(-1*int64(^uint(0)>>1)-1+unixToInternal, 0).UTC()

func MaxValue() TimeExact {
	return TimeOf(max)
}

func MinValue() TimeExact {
	return TimeOf(min)
}
