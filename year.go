package chron

import (
	"time"

	"github.com/dustinevan/time/chron/length"
)

type Year struct {
	time.Time
}

func NewYear(year int) Year {
	return Year{time.Date(year, time.Month(1), 1, 0, 0, 0, 0, time.UTC)}
}

func YearOf(time time.Time) Year {
	return NewYear(time.Year())
}

func (y Year) AsYear() Year           { return YearOf(y.Time) }
func (y Year) AsMonth() Month         { return MonthOf(y.Time) }
func (y Year) AsDay() Day             { return DayOf(y.Time) }
func (y Year) AsHour() Hour           { return HourOf(y.Time) }
func (y Year) AsMinute() Minute       { return MinuteOf(y.Time) }
func (y Year) AsSecond() Second       { return SecondOf(y.Time) }
func (y Year) AsMilli() Milli         { return MilliOf(y.Time) }
func (y Year) AsMicro() Micro         { return MicroOf(y.Time) }
func (y Year) AsTimeExact() TimeExact { return TimeOf(y.Time) }
func (y Year) AsTime() time.Time      { return y.Time }

func (y Year) Increment(l Length) TimeExact {
	return TimeExact{y.AddDate(l.Years(), l.Months(), l.Days()).Add(l.Duration())}
}

func (y Year) Decrement(l Length) TimeExact {
	return TimeExact{y.AddDate(-1*l.Years(), -1*l.Months(), -1*l.Days()).Add(-1*l.Duration())}
}

func (y Year) AddN(n int) Year {
	return Year{y.AddDate(n, 0, 0)}
}

// Period Implementation
func (y Year) Contains(t TimeExact) bool {
	return t.Year() == y.Year()
}

func (y Year) Before() TimeExact {
	return y.AsTimeExact().Decrement(length.Nano)
}

func (y Year) After() TimeExact {
	return y.AsTimeExact().Increment(length.Year)
}

func (y Year) Len() Length {
	return length.Year
}

/*func (y Year) AddYear(n int) Year {
	return NewYear(y.Year() + n)
}

func (y Year) AddMonth(m int) Month {
	return y.AsMonth().AddMonth(m)
}

func (y Year) AddDay(d int) Day {
	return NewDay(y.Year(), 1, d)
}

func (y Year) AddHour(h int) Hour {
	return NewHour(y.Year(), 1, 1, h)
}

func (y Year) AddMinute(m int) Minute {
}

func (y Year) AddSecond(s int) Second {
	return NewSecond(y.Year(), m, d)
}

func (y Year) AddMilli(m int) Milli {
	return NewMilli(y.Year)
}

func (y Year) AddMicro(m int) Micro {
	return NewMicro(y.Year)
}
*/
