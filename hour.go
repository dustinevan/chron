package chron

import (
	"time"

	"github.com/dustinevan/chron/dura"
	"fmt"
	"reflect"
	"database/sql/driver"
	"strings"
)

type Hour struct {
	time.Time
}

func NewHour(year int, month time.Month, day, hour int) Hour {
	return Hour{time.Date(year, month, day, hour, 0, 0, 0, time.UTC)}
}

func ThisHour() Hour {
	return Now().AsHour()
}

func HourOf(t time.Time) Hour {
	t = t.UTC()
	return NewHour(t.Year(), t.Month(), t.Day(), t.Hour())
}

func (h Hour) AsYear() Year       { return YearOf(h.Time) }
func (h Hour) AsMonth() Month     { return MonthOf(h.Time) }
func (h Hour) AsDay() Day         { return DayOf(h.Time) }
func (h Hour) AsHour() Hour       { return h }
func (h Hour) AsMinute() Minute   { return MinuteOf(h.Time) }
func (h Hour) AsSecond() Second   { return SecondOf(h.Time) }
func (h Hour) AsMilli() Milli     { return MilliOf(h.Time) }
func (h Hour) AsMicro() Micro     { return MicroOf(h.Time) }
func (h Hour) AsChron() Chron { return TimeOf(h.Time) }
func (h Hour) AsTime() time.Time  { return h.Time }

func (h Hour) Increment(l dura.Time) Chron {
	return Chron{h.AddDate(l.Years(), l.Months(), l.Days()).Add(l.Duration())}
}

func (h Hour) Decrement(l dura.Time) Chron {
	return Chron{h.AddDate(-1*l.Years(), -1*l.Months(), -1*l.Days()).Add(-1 * l.Duration())}
}

func (h Hour) AddN(n int) Hour {
	return Hour{h.Add(time.Duration(int(time.Hour) * n))}
}

// span.Time implementation
func (h Hour) Start() Chron {
	return h.AsChron()
}

func (h Hour) End() Chron {
	return h.AddN(1).Decrement(dura.Nano)
}

func (h Hour) Contains(t Span) bool {
	return !h.Before(t) && !h.After(t)
}

func (h Hour) Before(t Span) bool {
	return h.End().AsTime().Before(t.Start().AsTime())
}

func (h Hour) After(t Span) bool {
	return h.Start().AsTime().After(t.End().AsTime())
}

func (h Hour) Duration() dura.Time {
	return dura.Hour
}

func (h Hour) AddYears(y int) Hour {
	return h.Increment(dura.Years(y)).AsHour()
}

func (h Hour) AddMonths(m int) Hour {
	return h.Increment(dura.Months(m)).AsHour()
}

func (h Hour) AddDays(d int) Hour {
	return h.Increment(dura.Days(d)).AsHour()
}

func (h Hour) AddHours(hs int) Hour {
	return h.AddN(hs)
}

func (h Hour) AddMinutes(m int) Minute {
	return h.AsMinute().AddN(m)
}

func (h Hour) AddSeconds(s int) Second {
	return h.AsSecond().AddN(s)
}

func (h Hour) AddMillis(m int) Milli {
	return h.AsMilli().AddN(m)
}

func (h Hour) AddMicros(m int) Micro {
	return h.AsMicro().AddN(m)
}

func (h Hour) AddNanos(n int) Chron {
	return h.AsChron().AddN(n)
}

func (h *Hour) Scan(value interface{}) error {
	if value == nil {
		*h = ZeroValue().AsHour()
		return nil
	}
	if t, ok := value.(time.Time); ok {
		*h = HourOf(t)
		return nil
	}
	return fmt.Errorf("unsupported Scan, storing %s into type *chron.Day", reflect.TypeOf(value))
}

func (h Hour) Value() (driver.Value, error) {
	// todo: error check the range.
	return h.Time, nil
}

func (h *Hour) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	s := strings.Trim(string(data), `"`)
	t, err := Parse(s)
	*h = HourOf(t)
	return err
}