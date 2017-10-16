package chron

import (
	"time"

	"github.com/dustinevan/time/chron/length"
)

type Micro struct {
	time.Time
}

func NewMicro(year int, month time.Month, day, hour, min, sec, micro int) Micro {
	return Micro{time.Date(year, month, day, hour, min, sec, micro*1000, time.UTC)}
}

func MicroOf(t time.Time) Micro {
	return Micro{t.Truncate(time.Microsecond)}
}

func (m Micro) AsYear() Year           { return YearOf(m.Time) }
func (m Micro) AsMonth() Month         { return MonthOf(m.Time) }
func (m Micro) AsDay() Day             { return DayOf(m.Time) }
func (m Micro) AsHour() Hour           { return HourOf(m.Time) }
func (m Micro) AsMinute() Minute       { return MinuteOf(m.Time) }
func (m Micro) AsSecond() Second       { return SecondOf(m.Time) }
func (m Micro) AsMilli() Milli         { return MilliOf(m.Time) }
func (m Micro) AsMicro() Micro         { return MicroOf(m.Time) }
func (m Micro) AsExactTime() ExactTime { return TimeOf(m.Time) }
func (m Micro) AsTime() time.Time      { return m.Time }

func (m Micro) Increment(l Length) ExactTime {
	return ExactTime{m.AddDate(l.Years(), l.Months(), l.Days()).Add(l.Duration())}
}

func (m Micro) Decrement(l Length) ExactTime {
	return ExactTime{m.AddDate(-1*l.Years(), -1*l.Months(), -1*l.Days()).Add(l.Duration())}
}

func (m Micro) AddN(n int) Micro {
	return Micro{m.Add(time.Duration(int(time.Microsecond) * n))}
}

// Period Implementation
func (m Micro) Contains(t ExactTime) bool {
	return (t.Nanosecond()/1000)*1000 == t.AsMilli().Nanosecond()
}

func (m Micro) Before() ExactTime {
	return m.AsExactTime().Decrement(length.Nano)
}

func (m Micro) After() ExactTime {
	return m.AsExactTime().Increment(length.Nano)
}

func (m Micro) Len() Length {
	return length.Micro
}