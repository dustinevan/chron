package chron

import (
	"time"

	"github.com/dustinevan/chron/dura"
	"fmt"
	"reflect"
	"database/sql/driver"
	"strings"
)

type Minute struct {
	time.Time
}

func NewMinute(year int, month time.Month, day, hour, min int) Minute {
	return Minute{time.Date(year, month, day, hour, min, 0, 0, time.UTC)}
}

func ThisMinute() Minute {
	return Now().AsMinute()
}

func MinuteOf(t time.Time) Minute {
	t = t.UTC()
	return NewMinute(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
}

func (m Minute) AsYear() Year       { return YearOf(m.Time) }
func (m Minute) AsMonth() Month     { return MonthOf(m.Time) }
func (m Minute) AsDay() Day         { return DayOf(m.Time) }
func (m Minute) AsHour() Hour       { return HourOf(m.Time) }
func (m Minute) AsMinute() Minute   { return m }
func (m Minute) AsSecond() Second   { return SecondOf(m.Time) }
func (m Minute) AsMilli() Milli     { return MilliOf(m.Time) }
func (m Minute) AsMicro() Micro     { return MicroOf(m.Time) }
func (m Minute) AsChron() Chron { return TimeOf(m.Time) }
func (m Minute) AsTime() time.Time  { return m.Time }

func (m Minute) Increment(l dura.Time) Chron {
	return Chron{m.AddDate(l.Years(), l.Months(), l.Days()).Add(l.Duration())}
}

func (m Minute) Decrement(l dura.Time) Chron {
	return Chron{m.AddDate(-1*l.Years(), -1*l.Months(), -1*l.Days()).Add(-1 * l.Duration())}
}

func (m Minute) AddN(n int) Minute {
	return Minute{m.Add(time.Duration(int(time.Minute) * n))}
}

// span.Time implementation
func (m Minute) Start() Chron {
	return m.AsChron()
}

func (m Minute) End() Chron {
	return m.AddN(1).Decrement(dura.Nano)
}

func (m Minute) Contains(t Span) bool {
	return !m.Before(t) && !m.After(t)
}

func (m Minute) Before(t Span) bool {
	return m.End().AsTime().Before(t.Start().AsTime())
}

func (m Minute) After(t Span) bool {
	return m.Start().AsTime().After(t.End().AsTime())
}

func (m Minute) Duration() dura.Time {
	return dura.Minute
}

func (m Minute) AddYears(y int) Minute {
	return m.Increment(dura.Years(y)).AsMinute()
}

func (m Minute) AddMonths(ms int) Minute {
	return m.Increment(dura.Months(ms)).AsMinute()
}

func (m Minute) AddDays(d int) Minute {
	return m.Increment(dura.Days(d)).AsMinute()
}

func (m Minute) AddHours(h int) Minute {
	return m.Increment(dura.Hours(h)).AsMinute()
}

func (m Minute) AddMinutes(ms int) Minute {
	return m.AddN(ms)
}

func (m Minute) AddSeconds(s int) Second {
	return m.AsSecond().AddN(s)
}

func (m Minute) AddMillis(ms int) Milli {
	return m.AsMilli().AddN(ms)
}

func (m Minute) AddMicros(ms int) Micro {
	return m.AsMicro().AddN(ms)
}

func (m Minute) AddNanos(n int) Chron {
	return m.AsChron().AddN(n)
}

func (m *Minute) Scan(value interface{}) error {
	if value == nil {
		*m = ZeroValue().AsMinute()
		return nil
	}
	if t, ok := value.(time.Time); ok {
		*m = MinuteOf(t)
		return nil
	}
	return fmt.Errorf("unsupported Scan, storing %s into type *chron.Day", reflect.TypeOf(value))
}

func (m Minute) Value() (driver.Value, error) {
	// todo: error check the range.
	return m.Time, nil
}

func (m *Minute) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	s := strings.Trim(string(data), `"`)
	t, err := Parse(s)
	*m = MinuteOf(t)
	return err
}