package chron

import (
	"time"

	"github.com/dustinevan/chron/length"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/dustinevan/chron/dura"
)

type Day struct {
	time.Time
}

func NewDay(year int, month time.Month, day int) Day {
	return Day{time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
}

func Today() Day {
	return Now().AsDay()
}

func DayOf(time time.Time) Day {
	return NewDay(time.Year(), time.Month(), time.Day())
}

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

func (d Day) Increment(du dura.Duration) TimeExact {
	return TimeExact{d.AddDate(du.Years(), du.Months(), du.Days()).Add(du.Duration())}
}

func (d Day) Decrement(du dura.Duration) TimeExact {
	return TimeExact{d.AddDate(-1*du.Years(), -1*du.Months(), -1*du.Days()).Add(-1 * du.Duration())}
}

func (d Day) AddN(n int) Day {
	return Day{d.AddDate(0, 0, n)}
}

// Period Implementation
func (d Day) Contains(t TimeExact) bool {
	return t.Day() == d.Day()
}

func (d Day) Before() TimeExact {
	return d.AsTimeExact().Decrement(dura.Nano)
}

func (d Day) After() TimeExact {
	return d.AsTimeExact().Increment(duration.Day)
}

func (d Day) Len() Length {
	return duration.Day
}
