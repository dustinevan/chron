package chron

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/dustinevan/chron/dura"
	"github.com/stretchr/testify/assert"
)

var tnano = time.Date(2018, time.June, 5, 12, 10, 6, 55, time.UTC)
var chr = TimeOf(tnano)

func TestNewNano(t *testing.T) {
	assert.Equal(t, tnano, NewTime(2018, time.June, 5, 12, 10, 6, 55).Time)
	assert.Equal(t, tnano, TimeOf(tnano).Time)
}

func TestChron_Transfers(t *testing.T) {
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

func TestChron_Increment(t *testing.T) {
	y := chr.Increment(dura.NewDuration(1, 2, 30, time.Nanosecond * 500))
	td := tnano.AddDate(1, 2, 30).Add(time.Nanosecond * 500)
	assert.Exactly(t, td, y.Time)
}

func TestChron_AsTime(t *testing.T) {
	assert.Exactly(t, tnano, chr.AsTime())
}

func TestChron_Decrement(t *testing.T) {
	d := chr.Decrement(dura.NewDuration(1, 2, 30, time.Nanosecond * 500))
	td := tnano.AddDate(-1, -2, -30).Add(time.Nanosecond * -500)
	assert.Exactly(t, td, d.Time)

}

func TestChron_AddN(t *testing.T) {
	assert.Exactly(t, NewTime(2018, time.June, 5, 12, 10, 6, 58), chr.AddN(3))
}

func TestChron_Start(t *testing.T) {
	assert.Exactly(t, chr.Time, chr.Start().Time)
}

func TestChron_End(t *testing.T) {
	assert.Exactly(t, chr.Time.Add(time.Nanosecond+(-1*time.Nanosecond)), chr.End().Time)
}

func TestChron_Contains(t *testing.T) {
	assert.True(t, chr.Contains(NewTime(2018, time.June, 5, 12, 10, 6, 55)))
	assert.False(t, chr.Contains(chr.AddN(1)))
}

func TestChron_Duration(t *testing.T) {
	assert.Equal(t, chr.Duration().Duration(), dura.Nano.Duration())
}

func TestChron_AddFns(t *testing.T) {
	assert.Exactly(t, chr.AddYears(2), NewTime(2020, time.June, 5, 12, 10, 6, 55))
	assert.Exactly(t, chr.AddMonths(25), NewTime(2020, time.July, 5, 12, 10, 6, 55))
	assert.Exactly(t, chr.AddDays(2), NewTime(2018, time.June, 7, 12, 10, 6, 55))
	assert.Exactly(t, chr.AddHours(25), NewTime(2018, time.June, 6, 13, 10, 6, 55))
	assert.Exactly(t, chr.AddMinutes(72), NewTime(2018, time.June, 5, 13, 22, 6, 55))
	assert.Exactly(t, chr.AddSeconds(3672), NewTime(2018, time.June, 5, 13, 11, 18, 55))
	assert.Exactly(t, chr.AddMillis(3672001), NewTime(2018, time.June, 5, 13, 11, 18, 1000055))
	assert.Exactly(t, chr.AddMicros(3672000001), NewTime(2018, time.June, 5, 13, 11, 18, 1055))
	assert.Exactly(t, chr.AddNanos(3672000000001), NewTime(2018, time.June, 5, 13, 11, 18, 56))
}

func TestChron_Scan(t *testing.T) {
	var m Chron
	assert.Nil(t, m.Scan(tnano))
	assert.Exactly(t, chr, m)
	assert.Nil(t, m.Scan(nil))
	assert.Exactly(t, ZeroValue().AsChron(), m)
	assert.Error(t, m.Scan("wrong value"))
}

func TestChron_Value(t *testing.T) {
	v, err := chr.Value()
	assert.Nil(t, err)
	assert.Exactly(t, v, driver.Value(tnano))
}

func TestChron_UnmarshalJSON(t *testing.T) {
	var m Chron
	assert.Nil(t, m.UnmarshalJSON([]byte("null")))
	assert.Exactly(t, ZeroValue().AsChron(), m)
	assert.Error(t, m.UnmarshalJSON([]byte("as;dlkjfd")))
	assert.Exactly(t, ZeroValue().AsChron(), m)
	assert.Nil(t, m.UnmarshalJSON([]byte("\"2018-06-05T12:10:06.000000055Z\"")))
	assert.Exactly(t, chr, m)
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
