package chron

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/dustinevan/chron/dura"
	"github.com/stretchr/testify/assert"
)

var tnano = time.Date(2018, time.June, 5, 12, 10, 6, 55, time.UTC)
var nano = TimeOf(tnano)

func TestNewNano(t *testing.T) {
	assert.Equal(t, tnano, NewTime(2018, time.June, 5, 12, 10, 6, 55).Time)
	assert.Equal(t, tnano, TimeOf(tnano).Time)
}

func TestNano_Transfers(t *testing.T) {
	ch := Now()
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

func TestNano_Increment(t *testing.T) {
	y := nano.Increment(dura.Duration{Year: 1, Month: 2, Day: 30, Dur: time.Nanosecond * 500})
	td := tnano.AddDate(1, 2, 30).Add(time.Nanosecond * 500)
	assert.Exactly(t, td, y.Time)
}

func TestNano_AsTime(t *testing.T) {
	assert.Exactly(t, tnano, nano.AsTime())
}

func TestNano_Decrement(t *testing.T) {
	d := nano.Decrement(dura.Duration{Year: 1, Month: 2, Day: 30, Dur: time.Nanosecond * 500})
	td := tnano.AddDate(-1, -2, -30).Add(time.Nanosecond * -500)
	assert.Exactly(t, td, d.Time)

}

func TestNano_AddN(t *testing.T) {
	assert.Exactly(t, NewTime(2018, time.June, 5, 12, 10, 6, 58), nano.AddN(3))
}

func TestNano_Start(t *testing.T) {
	assert.Exactly(t, nano.Time, nano.Start().Time)
}

func TestNano_End(t *testing.T) {
	assert.Exactly(t, nano.Time.Add(time.Nanosecond+(-1*time.Nanosecond)), nano.End().Time)
}

func TestNano_Contains(t *testing.T) {
	assert.True(t, nano.Contains(NewTime(2018, time.June, 5, 12, 10, 6, 55)))
	assert.False(t, nano.Contains(nano.AddN(1)))
}

func TestNano_Duration(t *testing.T) {
	assert.Equal(t, nano.Duration().Duration(), dura.Nano.Duration())
}

func TestNano_AddFns(t *testing.T) {
	assert.Exactly(t, nano.AddYears(2), NewTime(2020, time.June, 5, 12, 10, 6, 55))
	assert.Exactly(t, nano.AddMonths(25), NewTime(2020, time.July, 5, 12, 10, 6, 55))
	assert.Exactly(t, nano.AddDays(2), NewTime(2018, time.June, 7, 12, 10, 6, 55))
	assert.Exactly(t, nano.AddHours(25), NewTime(2018, time.June, 6, 13, 10, 6, 55))
	assert.Exactly(t, nano.AddMinutes(72), NewTime(2018, time.June, 5, 13, 22, 6, 55))
	assert.Exactly(t, nano.AddSeconds(3672), NewTime(2018, time.June, 5, 13, 11, 18, 55))
	assert.Exactly(t, nano.AddMillis(3672001), NewTime(2018, time.June, 5, 13, 11, 18, 1000055))
	assert.Exactly(t, nano.AddMicros(3672000001), NewTime(2018, time.June, 5, 13, 11, 18, 1055))
	assert.Exactly(t, nano.AddNanos(3672000000001), NewTime(2018, time.June, 5, 13, 11, 18, 56))
}

func TestNano_Scan(t *testing.T) {
	var m Chron
	assert.Nil(t, m.Scan(tnano))
	assert.Exactly(t, nano, m)
	assert.Nil(t, m.Scan(nil))
	assert.Exactly(t, ZeroValue().AsChron(), m)
	assert.Error(t, m.Scan("wrong value"))
}

func TestNano_Value(t *testing.T) {
	v, err := nano.Value()
	assert.Nil(t, err)
	assert.Exactly(t, v, driver.Value(tnano))
}

func TestNano_UnmarshalJSON(t *testing.T) {
	var m Chron
	assert.Nil(t, m.UnmarshalJSON([]byte("null")))
	assert.Exactly(t, ZeroValue().AsChron(), m)
	assert.Error(t, m.UnmarshalJSON([]byte("as;dlkjfd")))
	assert.Exactly(t, ZeroValue().AsChron(), m)
	assert.Nil(t, m.UnmarshalJSON([]byte("\"2018-06-05T12:10:06.000000055Z\"")))
	assert.Exactly(t, nano, m)
}

func TestZeroYear(t *testing.T) {
	assert.Exactly(t, 0, ZeroYear().Year())
	assert.Exactly(t, time.January, ZeroYear().Month())
	assert.Exactly(t, 1, ZeroYear().Day())
	assert.Exactly(t, 0, ZeroYear().Hour())
	assert.Exactly(t, 0, ZeroYear().Minute())
	assert.Exactly(t, 0, ZeroYear().Second())
	assert.Exactly(t, 0, ZeroYear().Nanosecond())
}

func TestZeroUnix(t *testing.T) {
	assert.Exactly(t, TimeOf(time.Unix(0, 0)), ZeroUnix())
}
