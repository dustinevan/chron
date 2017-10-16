package chron

import (
	"time"

	"github.com/dustinevan/chron/dura"
	"fmt"
	"reflect"
	"database/sql/driver"
	"strings"
)

//
type Year struct {
	time.Time
}

func NewYear(year int) Year {
	return Year{time.Date(year, time.Month(1), 1, 0, 0, 0, 0, time.UTC)}
}

func ThisYear() Year {
	return Now().AsYear()
}

func YearOf(t time.Time) Year {
	t.UTC()
	return NewYear(t.Year())
}

func (y Year) AsYear() Year       { return y }
func (y Year) AsMonth() Month     { return MonthOf(y.Time) }
func (y Year) AsDay() Day         { return DayOf(y.Time) }
func (y Year) AsHour() Hour       { return HourOf(y.Time) }
func (y Year) AsMinute() Minute   { return MinuteOf(y.Time) }
func (y Year) AsSecond() Second   { return SecondOf(y.Time) }
func (y Year) AsMilli() Milli     { return MilliOf(y.Time) }
func (y Year) AsMicro() Micro     { return MicroOf(y.Time) }
func (y Year) AsChron() Chron { return TimeOf(y.Time) }
func (y Year) AsTime() time.Time  { return y.Time }

func (y Year) Increment(i dura.Time) Chron {
	return Chron{y.AddDate(i.Years(), i.Months(), i.Days()).Add(i.Duration())}
}

func (y Year) Decrement(i dura.Time) Chron {
	return Chron{y.AddDate(-1*i.Years(), -1*i.Months(), -1*i.Days()).Add(-1 * i.Duration())}
}

func (y Year) AddN(n int) Year {
	return Year{y.AddDate(n, 0, 0)}
}

// span.Time implementation
func (y Year) Start() Chron {
	return y.AsChron()
}

func (y Year) End() Chron {
	return y.AddN(1).Decrement(dura.Nano)
}

func (y Year) Contains(t Span) bool {
	return !y.Before(t) && !y.After(t)
}

func (y Year) Before(t Span) bool {
	return y.End().AsTime().Before(t.Start().AsTime())
}

func (y Year) After(t Span) bool {
	return y.Start().AsTime().After(t.End().AsTime())
}

func (y Year) Duration() dura.Time {
	return dura.Year
}

func (y Year) AddYears(ys int) Year {
	return y.AddN(ys)
}

func (y Year) AddMonths(m int) Month {
	return y.AsMonth().AddN(m)
}

func (y Year) AddDays(d int) Day {
	return y.AsDay().AddN(d)
}

func (y Year) AddHours(h int) Hour {
	return y.AsHour().AddN(h)
}

func (y Year) AddMinutes(m int) Minute {
	return y.AsMinute().AddN(m)
}

func (y Year) AddSeconds(s int) Second {
	return y.AsSecond().AddN(s)
}

func (y Year) AddMillis(m int) Milli {
	return y.AsMilli().AddN(m)
}

func (y Year) AddMicros(m int) Micro {
	return y.AsMicro().AddN(m)
}

func (y Year) AddNanos(n int) Chron {
	return y.AsChron().AddN(n)
}

func (y *Year) Scan(value interface{}) error {
	if value == nil {
		*y = ZeroValue().AsYear()
		return nil
	}
	if t, ok := value.(time.Time); ok {
		*y = YearOf(t)
		return nil
	}
	return fmt.Errorf("unsupported Scan, storing %s into type *chron.Day", reflect.TypeOf(value))
}

func (y Year) Value() (driver.Value, error) {
	// todo: error check the range.
	return y.Time, nil
}

func (y *Year) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	s := strings.Trim(string(data), `"`)
	t, err := Parse(s)
	*y = YearOf(t)
	return err
}