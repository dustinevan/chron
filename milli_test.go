package chron

import (
	"github.com/stretchr/testify/assert"
	"github.com/dustinevan/chron/dura"
	"time"
	"database/sql/driver"
	"testing"
)

var tmilli = time.Date(2018, time.June, 5, 12, 10, 6, 55000000, time.UTC)
var milli = MilliOf(tmilli)

func TestNewMilli(t *testing.T) {
	assert.Equal(t, tmilli, NewMilli(2018, time.June, 5, 12, 10, 6, 55).Time)
	assert.Equal(t, tmilli, MilliOf(tmilli).Time)
}

func TestMilli_Transfers(t *testing.T) {
	ch := ThisMilli()
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

func TestMilli_Increment(t *testing.T) {
	y := milli.Increment(dura.Duration{Year: 1, Month: 2, Day: 30, Dur: time.Millisecond * 500})
	td := tmilli.AddDate(1, 2, 30).Add(time.Millisecond * 500)
	assert.Exactly(t, td, y.Time)
}

func TestMilli_AsTime(t *testing.T) {
	assert.Exactly(t, tmilli, milli.AsTime())
}

func TestMilli_Decrement(t *testing.T) {
	d := milli.Decrement(dura.Duration{Year: 1, Month: 2, Day: 30, Dur: time.Millisecond * 500})
	td := tmilli.AddDate(-1, -2, -30).Add(time.Millisecond * -500)
	assert.Exactly(t, td, d.Time)

}

func TestMilli_AddN(t *testing.T) {
	assert.Exactly(t, NewMilli(2018, time.June, 5, 12, 10, 6, 58), milli.AddN(3))
}

func TestMilli_Start(t *testing.T) {
	assert.Exactly(t, milli.Time, milli.Start().Time)
}

func TestMilli_End(t *testing.T) {
	assert.Exactly(t, milli.Time.Add(time.Millisecond + (-1 * time.Nanosecond)), milli.End().Time)
}

func TestMilli_Contains(t *testing.T) {
	assert.True(t, milli.Contains(NewTime(2018, time.June, 5, 12, 10, 6, 55000456)))
	assert.False(t, milli.Contains(milli.AddN(1)))
}

func TestMilli_Duration(t *testing.T) {
	assert.Equal(t, milli.Duration().Duration(), dura.Milli.Duration())
}

func TestMilli_AddFns(t *testing.T) {
	assert.Exactly(t, milli.AddYears(2), NewMilli(2020, time.June, 5, 12, 10, 6, 55))
	assert.Exactly(t, milli.AddMonths(25), NewMilli(2020, time.July, 5, 12, 10,6, 55))
	assert.Exactly(t, milli.AddDays(2), NewMilli(2018, time.June, 7, 12, 10, 6, 55))
	assert.Exactly(t, milli.AddHours(25), NewMilli(2018, time.June, 6, 13, 10, 6, 55))
	assert.Exactly(t, milli.AddMinutes(72), NewMilli(2018, time.June, 5, 13, 22, 6, 55))
	assert.Exactly(t, milli.AddSeconds(3672), NewMilli(2018, time.June, 5, 13, 11, 18, 55))
	assert.Exactly(t, milli.AddMillis(3672001), NewMilli(2018, time.June, 5, 13, 11, 18, 56))
	assert.Exactly(t, milli.AddMicros(3672000001), NewMicro(2018, time.June, 5, 13, 11, 18, 55001))
	assert.Exactly(t, milli.AddNanos(3672000000001), NewTime(2018, time.June, 5, 13, 11, 18, 55000001))
}

func TestMilli_Scan(t *testing.T) {
	var m Milli
	assert.Nil(t, m.Scan(tmilli))
	assert.Exactly(t, milli, m)
	assert.Nil(t, m.Scan(nil))
	assert.Exactly(t, ZeroValue().AsMilli(), m)
	assert.Error(t, m.Scan("wrong value"))
}

func TestMilli_Value(t *testing.T) {
	v, err := milli.Value()
	assert.Nil(t, err)
	assert.Exactly(t, v, driver.Value(tmilli))
}

func TestMilli_UnmarshalJSON(t *testing.T) {
	var m Milli
	assert.Nil(t, m.UnmarshalJSON([]byte("null")))
	assert.Exactly(t, ZeroValue().AsMilli(), m)
	assert.Error(t, m.UnmarshalJSON([]byte("as;dlkjfd")))
	assert.Exactly(t, ZeroValue().AsMilli(), m)
	assert.Nil(t, m.UnmarshalJSON([]byte("\"2018-06-05T12:10:06.055Z\"")))
	assert.Exactly(t, milli, m)
}
