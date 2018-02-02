package chron

import (
	"github.com/stretchr/testify/assert"
	"github.com/dustinevan/chron/dura"
	"time"
	"database/sql/driver"
	"testing"
)

var tmonth = time.Date(2018, time.April, 1, 0, 0, 0, 0, time.UTC)
var month = MonthOf(tmonth)

func TestNewMonth(t *testing.T) {
	assert.Equal(t, tmonth, NewMonth(2018, 4).Time)
	assert.Equal(t, tmonth, MonthOf(tmonth).Time)
}

func TestThisMonth(t *testing.T) {
	now := MonthOf(time.Now())
	ty := ThisMonth()
	assert.Equal(t, now, ty)
}

func TestMonth_Transfers(t *testing.T) {
	ch := ThisMonth()
	assert.IsType(t, Year{}, ch.AsYear())
	assert.IsType(t, Month{}, ch.AsMonth())
	assert.IsType(t, Day{}, ch.AsDay())
	assert.IsType(t, Hour{}, ch.AsHour())
	assert.IsType(t, Minute{}, ch.AsMinute())
	assert.IsType(t, Second{}, ch.AsSecond())
	assert.IsType(t, Milli{}, ch.AsMilli())
	assert.IsType(t, Micro{}, ch.AsMicro())
	assert.IsType(t, Chron{}, ch.AsChron())
}

func TestMonth_Increment(t *testing.T) {
	y := month.Increment(dura.NewDuration(1, 2, 30, time.Second * 500))
	td := tmonth.AddDate(1, 2, 30).Add(time.Second * 500)
	assert.Exactly(t, td, y.Time)
}

func TestMonth_AsTime(t *testing.T) {
	assert.Exactly(t, tmonth, month.AsTime())
}

func TestMonth_Decrement(t *testing.T) {
	d := month.Decrement(dura.NewDuration(1, 2, 30, time.Second * 500))
	td := tmonth.AddDate(-1, -2, -30).Add(time.Second * -500)
	assert.Exactly(t, td, d.Time)

}

func TestMonth_AddN(t *testing.T) {
	assert.Exactly(t, NewMonth(2018, time.June), month.AddN(2))
}

func TestMonth_Start(t *testing.T) {
	assert.Exactly(t, month.Time, month.Start().Time)
}

func TestMonth_End(t *testing.T) {
	assert.Exactly(t, month.Time.AddDate(0,1, 0).Add(-1 * time.Nanosecond), month.End().Time)
}

func TestMonth_Contains(t *testing.T) {
	assert.True(t, month.Contains(NewMinute(2018, time.April, 5, 12, 45)))
	assert.False(t, month.Contains(month.AddN(1)))
}

func TestMonth_Duration(t *testing.T) {
	assert.Equal(t, month.Duration().Months(), dura.Month.Months())
}

func TestMonth_AddFns(t *testing.T) {
	assert.Exactly(t, month.AddYears(2), NewMonth(2020, time.April))
	assert.Exactly(t, month.AddMonths(25), NewMonth(2020, time.May))
	assert.Exactly(t, month.AddDays(2), NewDay(2018, time.April, 3))
	assert.Exactly(t, month.AddHours(25), NewHour(2018, time.April, 2, 1))
	assert.Exactly(t, month.AddMinutes(72), NewMinute(2018, time.April, 1, 1, 12))
	assert.Exactly(t, month.AddSeconds(3672), NewSecond(2018, time.April, 1, 1, 1, 12))
	assert.Exactly(t, month.AddMillis(3672001), NewMilli(2018, time.April, 1, 1, 1, 12, 1))
	assert.Exactly(t, month.AddMicros(3672000001), NewMicro(2018, time.April, 1, 1, 1, 12, 1))
	assert.Exactly(t, month.AddNanos(3672000000001), NewTime(2018, time.April, 1, 1, 1, 12, 1))
}

func TestMonth_Scan(t *testing.T) {
	var m Month
	assert.Nil(t, m.Scan(tmonth))
	assert.Exactly(t, month, m)
	assert.Nil(t, m.Scan(nil))
	assert.Exactly(t, ZeroValue().AsMonth(), m)
	assert.Error(t, m.Scan("wrong value"))
}

func TestMonth_Value(t *testing.T) {
	v, err := month.Value()
	assert.Nil(t, err)
	assert.Exactly(t, v, driver.Value(tmonth))
}

func TestMonth_UnmarshalJSON(t *testing.T) {
	var m Month
	assert.Nil(t, m.UnmarshalJSON([]byte("null")))
	assert.Exactly(t, ZeroValue().AsMonth(), m)
	assert.Error(t, m.UnmarshalJSON([]byte("as;dlkjfd")))
	assert.Exactly(t, ZeroValue().AsMonth(), m)
	assert.Nil(t, m.UnmarshalJSON([]byte("\"2018-04-01T00:00:00Z\"")))
	assert.Exactly(t, month, m)
}
