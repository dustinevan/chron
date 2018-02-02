package chron

import (
	"time"

	"github.com/dustinevan/chron/dura"
	"fmt"
	"reflect"
	"database/sql/driver"
	"strings"
)

// Time implementations are instants in time that are transferable to
// other instants with a different precision--year, month, day, hour,
// minute, second, milli, micro, nano (which is called TimeExact).
// Implementations are also transferable to an underlying time.Time
// via AsTime().
type Time interface {

	// Implementations of Time have methods that transfer the data to
	// structs with different precision. For example: 2017-01-05 12:45:06
	// is a has second precision, if this data were represented as a chron.Second
	// sec.AsDay() would truncate the time to 2017-01-05 00:00:00.
	AsYear() Year
	AsMonth() Month
	AsDay() Day
	AsHour() Hour
	AsMinute() Minute
	AsSecond() Second
	AsMilli() Milli
	AsMicro() Micro
	AsChron() Chron
	AsTime() time.Time
	Incrementer
}

type Incrementer interface {
	Increment(dura.Time) Chron
	Decrement(dura.Time) Chron
}

type Chron struct {
	time.Time
}

func Now() Chron {
	return TimeOf(time.Now().In(time.UTC))
}

func NewTime(year int, month time.Month, day, hour, min, sec, nano int) Chron {
	return Chron{time.Date(year, time.Month(month), day, hour, min, sec, nano, time.UTC)}
}

func TimeOf(t time.Time) Chron {
	t = t.UTC()
	return Chron{t}
}

func (t Chron) AsYear() Year       { return YearOf(t.Time) }
func (t Chron) AsMonth() Month     { return MonthOf(t.Time) }
func (t Chron) AsDay() Day         { return DayOf(t.Time) }
func (t Chron) AsHour() Hour       { return HourOf(t.Time) }
func (t Chron) AsMinute() Minute   { return MinuteOf(t.Time) }
func (t Chron) AsSecond() Second   { return SecondOf(t.Time) }
func (t Chron) AsMilli() Milli     { return MilliOf(t.Time) }
func (t Chron) AsMicro() Micro     { return MicroOf(t.Time) }
func (t Chron) AsChron() Chron { return t }
func (t Chron) AsTime() time.Time  { return t.Time }

func (t Chron) Increment(d dura.Time) Chron {
	return Chron{t.AddDate(d.Years(), d.Months(), d.Days()).Add(d.Duration())}
}

func (t Chron) Decrement(d dura.Time) Chron {
	return Chron{t.AddDate(-1*d.Years(), -1*d.Months(), -1*d.Days()).Add(-1 * d.Duration())}
}

// AddN adds n Nanoseconds to the TimeExact
func (t Chron) AddN(n int) Chron {
	return TimeOf(t.AsTime().Add(time.Duration(n)))
}

// span.Time implementation
func (t Chron) Start() Chron {
	return t
}

func (t Chron) End() Chron {
	return t
}

func (t Chron) Contains(s Span) bool {
	return !t.Before(s) && !t.After(s)
}

func (t Chron) Before(s Span) bool {
	return t.End().AsTime().Before(s.Start().AsTime())
}

func (t Chron) After(s Span) bool {
	return t.Start().AsTime().After(s.End().AsTime())
}

func (t Chron) Duration() dura.Time {
	return dura.Nano
}



func (t Chron) AddYears(y int) Chron {
	return t.Increment(dura.Years(y))
}

func (t Chron) AddMonths(m int) Chron {
	return t.Increment(dura.Months(m))
}

func (t Chron) AddDays(d int) Chron {
	return t.Increment(dura.Days(d))
}

func (t Chron) AddHours(h int) Chron {
	return t.Increment(dura.Hours(h))
}

func (t Chron) AddMinutes(m int) Chron {
	return t.Increment(dura.Mins(m))
}

func (t Chron) AddSeconds(s int) Chron {
	return t.Increment(dura.Secs(s))
}

func (t Chron) AddMillis(m int) Chron {
	return t.Increment(dura.Millis(m))
}

func (t Chron) AddMicros(m int) Chron {
	return t.Increment(dura.Micros(m))
}

func (t Chron) AddNanos(n int) Chron {
	return t.AddN(n)
}

func (t *Chron) Scan(value interface{}) error {
	if value == nil {
		*t = ZeroValue().AsChron()
		return nil
	}
	if tt, ok := value.(time.Time); ok {
		*t = TimeOf(tt)
		return nil
	}
	return fmt.Errorf("unsupported Scan, storing %s into type *chron.Day", reflect.TypeOf(value))
}

func (t Chron) Value() (driver.Value, error) {
	// todo: error check the range.
	return t.Time, nil
}


func ZeroValue() Chron {
	return TimeOf(time.Time{})
}

func ZeroYear() Chron {
	return NewYear(0).AsChron()
}

func ZeroUnix() Chron {
	return TimeOf(time.Unix(0, 0))
}

func ZeroTime() time.Time {
	return time.Time{}.UTC()
}

// see: https://stackoverflow.com/questions/25065055/what-is-the-maximum-time-time-in-go
// and time.Unix() implementation
var unixToInternal = int64((1969*365 + 1969/4 - 1969/100 + 1969/400) * 24 * 60 * 60)
var max = time.Unix(1<<63-1-unixToInternal, 999999999).UTC()
var minimum = time.Unix(-1*int64(^uint(0)>>1)-1+unixToInternal, 0).UTC()

func MaxValue() Chron {
	return TimeOf(max)
}

func MinValue() Chron {
	return TimeOf(minimum)
}

func Parse(s string) (time.Time, error) {
	var errs []error
	for _, fn := range ParseFunctions {
		t, err := fn(s)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		return t, nil
	}
	return ZeroTime(), ErrJoin(errs, "; ")
}

func ErrJoin(errs []error, delim string) error {
	s := make([]string, 0)
	for _, e := range errs {
		s = append(s, e.Error())
	}
	return fmt.Errorf("%s", strings.Join(s, delim))
}
