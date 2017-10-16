package chron

import (
	"time"

	"github.com/dustinevan/time/chron/length"
)

type Hour struct {
	time.Time
}

func NewHour(year int, month time.Month, day, hour int) Hour {
	return Hour{time.Date(year, month, day, hour, 0, 0, 0, time.UTC)}
}

func HourOf(time time.Time) Hour {
	return NewHour(time.Year(), time.Month(), time.Day(), time.Hour())
}

func (h Hour) AsYear() Year           { return YearOf(h.Time) }
func (h Hour) AsMonth() Month         { return MonthOf(h.Time) }
func (h Hour) AsDay() Day             { return DayOf(h.Time) }
func (h Hour) AsHour() Hour           { return HourOf(h.Time) }
func (h Hour) AsMinute() Minute       { return MinuteOf(h.Time) }
func (h Hour) AsSecond() Second       { return SecondOf(h.Time) }
func (h Hour) AsMilli() Milli         { return MilliOf(h.Time) }
func (h Hour) AsMicro() Micro         { return MicroOf(h.Time) }
func (h Hour) AsExactTime() ExactTime { return TimeOf(h.Time) }
func (h Hour) AsTime() time.Time      { return h.Time }

func (h Hour) Increment(l Length) ExactTime {
	return ExactTime{h.AddDate(l.Years(), l.Months(), l.Days()).Add(l.Duration())}
}

func (h Hour) Decrement(l Length) ExactTime {
	return ExactTime{h.AddDate(-1*l.Years(), -1*l.Months(), -1*l.Days()).Add(l.Duration())}
}

func (h Hour) AddN(n int) Hour {
	return Hour{h.Add(time.Duration(int(time.Hour) * n))}
}

// Period Implementation
func (h Hour) Contains(t ExactTime) bool {
	return t.Hour() == h.Hour()
}

func (h Hour) Before() ExactTime {
	return h.AsExactTime().Decrement(length.Nano)
}

func (h Hour) After() ExactTime {
	return h.AsExactTime().Increment(length.Nano)
}

func (h Hour) Len() Length {
	return length.Hour
}