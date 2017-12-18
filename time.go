package chron

import (
	"time"

	"github.com/dustinevan/chron/dura"
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

	Increment(dura.Time) TimeExact
	Decrement(dura.Time) TimeExact
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
func (t TimeExact) AsTimeExact() TimeExact { return t }
func (t TimeExact) AsTime() time.Time      { return t.Time }

func (t TimeExact) Increment(d dura.Time) TimeExact {
	return TimeExact{t.AddDate(d.Years(), d.Months(), d.Days()).Add(d.Duration())}
}

func (t TimeExact) Decrement(d dura.Time) TimeExact {
	return TimeExact{t.AddDate(-1*d.Years(), -1*d.Months(), -1*d.Days()).Add(-1 * d.Duration())}
}

// AddN adds n Nanoseconds to the TimeExact
func (t TimeExact) AddN(n int) TimeExact {
	return TimeOf(t.AsTime().Add(time.Duration(n)))
}

// span.Time implementation
func (t TimeExact) Start() Time {
	return t
}

func (t TimeExact) End() Time {
	return t
}

func (t TimeExact) Contains(s Span) bool {
	return !t.Before(s) && !t.After(s)
}

func (t TimeExact) Before(s Span) bool {
	return t.End().AsTime().Before(s.Start().AsTime())
}

func (t TimeExact) After(s Span) bool {
	return t.Start().AsTime().After(s.End().AsTime())
}

func (t TimeExact) Duration() dura.Time {
	return dura.Micro
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

func (t TimeExact) AddYears(y int) TimeExact {
	return t.Increment(dura.Years(y))
}

func (t TimeExact) AddMonths(m int) TimeExact {
	return t.Increment(dura.Months(m))
}

func (t TimeExact) AddDays(d int) TimeExact {
	return t.Increment(dura.Days(d))
}

func (t TimeExact) AddHours(h int) TimeExact {
	return t.Increment(dura.Hours(h))
}

func (t TimeExact) AddMinutes(m int) TimeExact {
	return t.Increment(dura.Mins(m))
}

func (t TimeExact) AddSeconds(s int) TimeExact {
	return t.Increment(dura.Secs(s))
}

func (t TimeExact) AddMillis(m int) TimeExact {
	return t.Increment(dura.Millis(m))
}

func (t TimeExact) AddMicro(m int) TimeExact {
	return t.Increment(dura.Micros(m))
}

func (t TimeExact) AddNano(n int) TimeExact {
	return t.AddN(n)
}
