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

func NowDay() Day {
	return Now().AsDay()
}

func DayOf(time time.Time) Day {
	return NewDay(time.Year(), time.Month(), time.Day())
}

// chron.Time implementation
func (d Day) AsYear() Year           { return YearOf(d.Time) }
func (d Day) AsMonth() Month         { return MonthOf(d.Time) }
func (d Day) AsHour() Hour           { return HourOf(d.Time) }
func (d Day) AsDay() Day             { return DayOf(d.Time) }
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
func (d Day) Start() Time {
	return d.AsTimeExact()
}

func (d Day) End() Time {
	return d.AsTimeExact()
}

func (d Day) Contains(t Span) bool {
	return t.Start().Day() == d.Day()
}

func (d Day) Before(t Span) bool {
	return false
}

func (d Day) After(t Span) bool {
	return )
}

func (d Day) Duration() dura.Time {
	return dura.Day
}*/
