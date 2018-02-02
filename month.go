package chron

import (
	"time"

	"github.com/dustinevan/chron/dura"
	"fmt"
	"reflect"
	"database/sql/driver"
	"strings"
)

type Month struct {
	time.Time
}

func NewMonth(year int, month time.Month) Month {
	return Month{time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)}
}

func ThisMonth() Month {
	return Now().AsMonth()
}

func MonthOf(t time.Time) Month {
	t = t.UTC()
	return NewMonth(t.Year(), t.Month())
}

func (m Month) AsYear() Year       { return YearOf(m.Time) }
func (m Month) AsMonth() Month     { return m }
func (m Month) AsDay() Day         { return DayOf(m.Time) }
func (m Month) AsHour() Hour       { return HourOf(m.Time) }
func (m Month) AsMinute() Minute   { return MinuteOf(m.Time) }
func (m Month) AsSecond() Second   { return SecondOf(m.Time) }
func (m Month) AsMilli() Milli     { return MilliOf(m.Time) }
func (m Month) AsMicro() Micro     { return MicroOf(m.Time) }
func (m Month) AsChron() Chron { return TimeOf(m.Time) }
func (m Month) AsTime() time.Time  { return m.Time }

func (m Month) Increment(l dura.Time) Chron {
	return Chron{m.AddDate(l.Years(), l.Months(), l.Days()).Add(l.Duration())}
}

func (m Month) Decrement(l dura.Time) Chron {
	return Chron{m.AddDate(-1*l.Years(), -1*l.Months(), -1*l.Days()).Add(-1 * l.Duration())}
}

func (m Month) AddN(n int) Month {
	return Month{m.AddDate(0, n, 0)}
}

// span.Time implementation
func (m Month) Start() Chron {
	return m.AsChron()
}

func (m Month) End() Chron {
	return m.AddN(1).Decrement(dura.Nano)
}

func (m Month) Contains(t Span) bool {
	return !m.Before(t) && !m.After(t)
}

func (m Month) Before(t Span) bool {
	return m.End().AsTime().Before(t.Start().AsTime())
}

func (m Month) After(t Span) bool {
	return m.Start().AsTime().After(t.End().AsTime())
}

func (m Month) Duration() dura.Time {
	return dura.Month
}

func (m Month) AddYears(y int) Month {
	return m.Increment(dura.Years(y)).AsMonth()
}

func (m Month) AddMonths(ms int) Month {
	return m.AddN(ms)
}

func (m Month) AddDays(d int) Day {
	return m.AsDay().AddN(d)
}

func (m Month) AddHours(h int) Hour {
	return m.AsHour().AddN(h)
}

func (m Month) AddMinutes(mi int) Minute {
	return m.AsMinute().AddN(mi)
}

func (m Month) AddSeconds(s int) Second {
	return m.AsSecond().AddN(s)
}

func (m Month) AddMillis(mi int) Milli {
	return m.AsMilli().AddN(mi)
}

func (m Month) AddMicros(mi int) Micro {
	return m.AsMicro().AddN(mi)
}

func (m Month) AddNanos(n int) Chron {
	return m.AsChron().AddN(n)
}

func (m *Month) Scan(value interface{}) error {
	if value == nil {
		*m = ZeroValue().AsMonth()
		return nil
	}
	if t, ok := value.(time.Time); ok {
		*m = MonthOf(t)
		return nil
	}
	return fmt.Errorf("unsupported Scan, storing %s into type *chron.Day", reflect.TypeOf(value))
}

func (m Month) Value() (driver.Value, error) {
	// todo: error check the range.
	return m.Time, nil
}

func (m *Month) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	s := strings.Trim(string(data), `"`)
	t, err := Parse(s)
	*m = MonthOf(t)
	return err
}