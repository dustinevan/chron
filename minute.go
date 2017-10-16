package chron

import (
	"time"

	"github.com/dustinevan/time/chron/length"
)

type Minute struct {
	time.Time
}

func NewMinute(year int, month time.Month, day, hour, min int) Minute {
	return Minute{time.Date(year, month, day, hour, min, 0, 0, time.UTC)}
}

func MinuteOf(time time.Time) Minute {
	return NewMinute(time.Year(), time.Month(), time.Day(), time.Hour(), time.Minute())
}

func (m Minute) AsYear() Year           { return YearOf(m.Time) }
func (m Minute) AsMonth() Month         { return MonthOf(m.Time) }
func (m Minute) AsDay() Day             { return DayOf(m.Time) }
func (m Minute) AsHour() Hour           { return HourOf(m.Time) }
func (m Minute) AsMinute() Minute       { return MinuteOf(m.Time) }
func (m Minute) AsSecond() Second       { return SecondOf(m.Time) }
func (m Minute) AsMilli() Milli         { return MilliOf(m.Time) }
func (m Minute) AsMicro() Micro         { return MicroOf(m.Time) }
func (m Minute) AsExactTime() ExactTime { return TimeOf(m.Time) }
func (m Minute) AsTime() time.Time      { return m.Time }

func (m Minute) Increment(l Length) ExactTime {
	return ExactTime{m.AddDate(l.Years(), l.Months(), l.Days()).Add(l.Duration())}
}

func (m Minute) Decrement(l Length) ExactTime {
	return ExactTime{m.AddDate(-1*l.Years(), -1*l.Months(), -1*l.Days()).Add(l.Duration())}
}

func (m Minute) AddN(n int) Minute {
	return Minute{m.Add(time.Duration(int(time.Minute) * n))}
}

// Period Implementation
func (m Minute) Contains(t ExactTime) bool {
	return t.Minute() == m.Minute()
}

func (m Minute) Before() ExactTime {
	return m.AsExactTime().Decrement(length.Nano)
}

func (m Minute) After() ExactTime {
	return m.AsExactTime().Increment(length.Nano)
}

func (m Minute) Len() Length {
	return length.Minute
}
