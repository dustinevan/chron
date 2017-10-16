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
	AsExactTime() ExactTime
	AsTime() time.Time

	Increment(length Length) ExactTime
	Decrement(length Length) ExactTime
}

type ExactTime struct {
	time.Time
}

func Now() ExactTime {
	return TimeOf(time.Now())
}

func NewTime(year int, month time.Month, day, hour, min, sec, nano int) ExactTime {
	return ExactTime{time.Date(year, time.Month(month), day, hour, min, sec, nano, time.UTC)}
}

func TimeOf(t time.Time) ExactTime {
	return ExactTime{t}
}

func (t ExactTime) AsYear() Year           { return YearOf(t.Time) }
func (t ExactTime) AsMonth() Month         { return MonthOf(t.Time) }
func (t ExactTime) AsDay() Day             { return DayOf(t.Time) }
func (t ExactTime) AsHour() Hour           { return HourOf(t.Time) }
func (t ExactTime) AsMinute() Minute       { return MinuteOf(t.Time) }
func (t ExactTime) AsSecond() Second       { return SecondOf(t.Time) }
func (t ExactTime) AsMilli() Milli         { return MilliOf(t.Time) }
func (t ExactTime) AsMicro() Micro         { return MicroOf(t.Time) }
func (t ExactTime) AsExactTime() ExactTime { return TimeOf(t.Time) }
func (t ExactTime) AsTime() time.Time      { return t.Time }

func (t ExactTime) Increment(l Length) ExactTime {
	return ExactTime{t.AddDate(l.Years(), l.Months(), l.Days()).Add(l.Duration())}
}

func (t ExactTime) Decrement(l Length) ExactTime {
	return ExactTime{t.AddDate(-1*l.Years(), -1*l.Months(), -1*l.Days()).Add(-1 * l.Duration())}
}

func ZeroValue() ExactTime {
	return ExactTime{time.Time{}}
}
