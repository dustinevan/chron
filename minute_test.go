package chron

import (
	"github.com/stretchr/testify/assert"
	"github.com/dustinevan/chron/dura"
	"time"
	"database/sql/driver"
	"testing"
)

var tmin = time.Date(2018, time.June, 5, 12, 10, 0, 0, time.UTC)
var min = MinuteOf(tmin)

func TestNewMinute(t *testing.T) {
	assert.Equal(t, tmin, NewMinute(2018, time.June, 5, 12, 10).Time)
	assert.Equal(t, tmin, MinuteOf(tmin).Time)
}

func TestThisMinute(t *testing.T) {
	now := MinuteOf(time.Now())
	ty := ThisMinute()
	assert.Equal(t, now, ty)
}

func TestMinute_Transfers(t *testing.T) {
	ch := ThisMinute()
	assert.IsType(t, Year{}, ch.AsYear())
	assert.IsType(t, Month{}, ch.AsMonth())
	assert.IsType(t, Day{}, ch.AsDay())
	assert.IsType(t, Hour{}, ch.AsHour())
	assert.IsType(t, Minute{}, ch.AsMinute())
	assert.IsType(t, Second{}, ch.AsSecond())
	assert.IsType(t, Milli{}, ch.AsMilli())
	assert.IsType(t, Micro{}, ch.AsMicro())
	assert.IsType(t, TimeExact{}, ch.AsTimeExact())
}

func TestMinute_Increment(t *testing.T) {
	y := min.Increment(dura.Duration{Year: 1, Month: 2, Day: 30, Dur: time.Second * 500})
	td := tmin.AddDate(1, 2, 30).Add(time.Second * 500)
	assert.Exactly(t, td, y.Time)
}

func TestMinute_AsTime(t *testing.T) {
	assert.Exactly(t, tmin, min.AsTime())
}

func TestMinute_Decrement(t *testing.T) {
	d := min.Decrement(dura.Duration{Year: 1, Month: 2, Day: 30, Dur: time.Second * 500})
	td := tmin.AddDate(-1, -2, -30).Add(time.Second * -500)
	assert.Exactly(t, td, d.Time)

}

func TestMinute_AddN(t *testing.T) {
	assert.Exactly(t, NewMinute(2018, time.June, 5, 12, 12), min.AddN(2))
}

func TestMinute_Start(t *testing.T) {
	assert.Exactly(t, min.Time, min.Start().Time)
}

func TestMinute_End(t *testing.T) {
	assert.Exactly(t, min.Time.Add(time.Minute + (-1 * time.Nanosecond)), min.End().Time)
}

func TestMinute_Contains(t *testing.T) {
	assert.True(t, min.Contains(NewSecond(2018, time.June, 5, 12, 10, 10)))
	assert.False(t, min.Contains(min.AddN(1)))
}

func TestMinute_Duration(t *testing.T) {
	assert.Equal(t, min.Duration().Duration(), dura.Minute.Duration())
}

func TestMinute_AddFns(t *testing.T) {
	assert.Exactly(t, min.AddYears(2), NewMinute(2020, time.June, 5, 12, 10))
	assert.Exactly(t, min.AddMonths(25), NewMinute(2020, time.July, 5, 12, 10))
	assert.Exactly(t, min.AddDays(2), NewMinute(2018, time.June, 7, 12, 10))
	assert.Exactly(t, min.AddHours(25), NewMinute(2018, time.June, 6, 13, 10))
	assert.Exactly(t, min.AddMinutes(72), NewMinute(2018, time.June, 5, 13, 22))
	assert.Exactly(t, min.AddSeconds(3672), NewSecond(2018, time.June, 5, 13, 11, 12))
	assert.Exactly(t, min.AddMillis(3672001), NewMilli(2018, time.June, 5, 13, 11, 12, 1))
	assert.Exactly(t, min.AddMicros(3672000001), NewMicro(2018, time.June, 5, 13, 11, 12, 1))
	assert.Exactly(t, min.AddNanos(3672000000001), NewTime(2018, time.June, 5, 13, 11, 12, 1))
}

func TestMinute_Scan(t *testing.T) {
	var m Minute
	assert.Nil(t, m.Scan(tmin))
	assert.Exactly(t, min, m)
	assert.Nil(t, m.Scan(nil))
	assert.Exactly(t, ZeroValue().AsMinute(), m)
	assert.Error(t, m.Scan("wrong value"))
}

func TestMinute_Value(t *testing.T) {
	v, err := min.Value()
	assert.Nil(t, err)
	assert.Exactly(t, v, driver.Value(tmin))
}

func TestMinute_UnmarshalJSON(t *testing.T) {
	var m Minute
	assert.Nil(t, m.UnmarshalJSON([]byte("null")))
	assert.Exactly(t, ZeroValue().AsMinute(), m)
	assert.Error(t, m.UnmarshalJSON([]byte("as;dlkjfd")))
	assert.Exactly(t, ZeroValue().AsMinute(), m)
	assert.Nil(t, m.UnmarshalJSON([]byte("\"2018-06-05T12:10:00Z\"")))
	assert.Exactly(t, min, m)
}

