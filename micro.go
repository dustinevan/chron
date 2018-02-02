package chron

import (
	"time"

	"github.com/dustinevan/chron/dura"
	"fmt"
	"reflect"
	"database/sql/driver"
	"strings"
)

type Micro struct {
	time.Time
}

func NewMicro(year int, month time.Month, day, hour, min, sec, micro int) Micro {
	return Micro{time.Date(year, month, day, hour, min, sec, micro*1000, time.UTC)}
}

func ThisMicro() Micro {
	return Now().AsMicro()
}

func MicroOf(t time.Time) Micro {
	t = t.UTC()
	return Micro{t.Truncate(time.Microsecond)}
}

func (m Micro) AsYear() Year       { return YearOf(m.Time) }
func (m Micro) AsMonth() Month     { return MonthOf(m.Time) }
func (m Micro) AsDay() Day         { return DayOf(m.Time) }
func (m Micro) AsHour() Hour       { return HourOf(m.Time) }
func (m Micro) AsMinute() Minute   { return MinuteOf(m.Time) }
func (m Micro) AsSecond() Second   { return SecondOf(m.Time) }
func (m Micro) AsMilli() Milli     { return MilliOf(m.Time) }
func (m Micro) AsMicro() Micro     { return m }
func (m Micro) AsChron() Chron { return TimeOf(m.Time) }
func (m Micro) AsTime() time.Time  { return m.Time }

func (m Micro) Increment(l dura.Time) Chron {
	return Chron{m.AddDate(l.Years(), l.Months(), l.Days()).Add(l.Duration())}
}

func (m Micro) Decrement(l dura.Time) Chron {
	return Chron{m.AddDate(-1*l.Years(), -1*l.Months(), -1*l.Days()).Add(-1 * l.Duration())}
}

func (m Micro) AddN(n int) Micro {
	return Micro{m.Add(time.Duration(int(time.Microsecond) * n))}
}

// span.Time implementation
func (m Micro) Start() Chron {
	return m.AsChron()
}

func (m Micro) End() Chron {
	return m.AddN(1).Decrement(dura.Nano)
}

func (m Micro) Contains(t Span) bool {
	return !m.Before(t) && !m.After(t)
}

func (m Micro) Before(t Span) bool {
	return m.End().AsTime().Before(t.Start().AsTime())
}

func (m Micro) After(t Span) bool {
	return m.Start().AsTime().After(t.End().AsTime())
}

func (m Micro) Duration() dura.Time {
	return dura.Micro
}

func (m Micro) AddYears(y int) Micro {
	return m.Increment(dura.Years(y)).AsMicro()
}

func (m Micro) AddMonths(ms int) Micro {
	return m.Increment(dura.Months(ms)).AsMicro()
}

func (m Micro) AddDays(d int) Micro {
	return m.Increment(dura.Days(d)).AsMicro()
}

func (m Micro) AddHours(h int) Micro {
	return m.Increment(dura.Hours(h)).AsMicro()
}

func (m Micro) AddMinutes(ms int) Micro {
	return m.Increment(dura.Mins(ms)).AsMicro()
}

func (m Micro) AddSeconds(s int) Micro {
	return m.Increment(dura.Secs(s)).AsMicro()
}

func (m Micro) AddMillis(ms int) Micro {
	return m.Increment(dura.Millis(ms)).AsMicro()
}

func (m Micro) AddMicros(ms int) Micro {
	return m.AddN(ms)
}

func (m Micro) AddNanos(n int) Chron {
	return m.AsChron().AddN(n)
}

func (m *Micro) Scan(value interface{}) error {
	if value == nil {
		*m = ZeroValue().AsMicro()
		return nil
	}
	if t, ok := value.(time.Time); ok {
		*m = MicroOf(t)
		return nil
	}
	return fmt.Errorf("unsupported Scan, storing %s into type *chron.Day", reflect.TypeOf(value))
}

func (m Micro) Value() (driver.Value, error) {
	// todo: error check the range.
	return m.Time, nil
}

func (m *Micro) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	s := strings.Trim(string(data), `"`)
	t, err := Parse(s)
	*m = MicroOf(t)
	return err
}