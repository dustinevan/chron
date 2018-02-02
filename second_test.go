package chron

import (
	"github.com/stretchr/testify/assert"
	"github.com/dustinevan/chron/dura"
	"time"
	"database/sql/driver"
	"testing"
)

var tsec = time.Date(2018, time.June, 5, 12, 10, 6, 0, time.UTC)
var sec = SecondOf(tsec)

func TestNewSecond(t *testing.T) {
	assert.Equal(t, tsec, NewSecond(2018, time.June, 5, 12, 10, 6).Time)
	assert.Equal(t, tsec, SecondOf(tsec).Time)
}

func TestThisSecond(t *testing.T) {
	now := SecondOf(time.Now())
	ty := ThisSecond()
	assert.Equal(t, now, ty)
}

func TestSecond_Transfers(t *testing.T) {
	ch := ThisSecond()
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

func TestSecond_Increment(t *testing.T) {
	y := sec.Increment(dura.Duration{Year: 1, Month: 2, Day: 30, Dur: time.Second * 500})
	td := tsec.AddDate(1, 2, 30).Add(time.Second * 500)
	assert.Exactly(t, td, y.Time)
}

func TestSecond_AsTime(t *testing.T) {
	assert.Exactly(t, tsec, sec.AsTime())
}

func TestSecond_Decrement(t *testing.T) {
	d := sec.Decrement(dura.Duration{Year: 1, Month: 2, Day: 30, Dur: time.Second * 500})
	td := tsec.AddDate(-1, -2, -30).Add(time.Second * -500)
	assert.Exactly(t, td, d.Time)

}

func TestSecond_AddN(t *testing.T) {
	assert.Exactly(t, NewSecond(2018, time.June, 5, 12, 10, 8), sec.AddN(2))
}

func TestSecond_Start(t *testing.T) {
	assert.Exactly(t, sec.Time, sec.Start().Time)
}

func TestSecond_End(t *testing.T) {
	assert.Exactly(t, sec.Time.Add(time.Second + (-1 * time.Nanosecond)), sec.End().Time)
}

func TestSecond_Contains(t *testing.T) {
	assert.True(t, sec.Contains(NewTime(2018, time.June, 5, 12, 10, 6, 234973)))
	assert.False(t, sec.Contains(sec.AddN(1)))
}

func TestSecond_Duration(t *testing.T) {
	assert.Equal(t, sec.Duration().Duration(), dura.Second.Duration())
}

func TestSecond_AddFns(t *testing.T) {
	assert.Exactly(t, sec.AddYears(2), NewSecond(2020, time.June, 5, 12, 10, 6))
	assert.Exactly(t, sec.AddMonths(25), NewSecond(2020, time.July, 5, 12, 10,6))
	assert.Exactly(t, sec.AddDays(2), NewSecond(2018, time.June, 7, 12, 10, 6))
	assert.Exactly(t, sec.AddHours(25), NewSecond(2018, time.June, 6, 13, 10, 6))
	assert.Exactly(t, sec.AddMinutes(72), NewSecond(2018, time.June, 5, 13, 22, 6))
	assert.Exactly(t, sec.AddSeconds(3672), NewSecond(2018, time.June, 5, 13, 11, 18))
	assert.Exactly(t, sec.AddMillis(3672001), NewMilli(2018, time.June, 5, 13, 11, 18, 1))
	assert.Exactly(t, sec.AddMicros(3672000001), NewMicro(2018, time.June, 5, 13, 11, 18, 1))
	assert.Exactly(t, sec.AddNanos(3672000000001), NewTime(2018, time.June, 5, 13, 11, 18, 1))
}

func TestSecond_Scan(t *testing.T) {
	var m Second
	assert.Nil(t, m.Scan(tsec))
	assert.Exactly(t, sec, m)
	assert.Nil(t, m.Scan(nil))
	assert.Exactly(t, ZeroValue().AsSecond(), m)
	assert.Error(t, m.Scan("wrong value"))
}

func TestSecond_Value(t *testing.T) {
	v, err := sec.Value()
	assert.Nil(t, err)
	assert.Exactly(t, v, driver.Value(tsec))
}

func TestSecond_UnmarshalJSON(t *testing.T) {
	var m Second
	assert.Nil(t, m.UnmarshalJSON([]byte("null")))
	assert.Exactly(t, ZeroValue().AsSecond(), m)
	assert.Error(t, m.UnmarshalJSON([]byte("as;dlkjfd")))
	assert.Exactly(t, ZeroValue().AsSecond(), m)
	assert.Nil(t, m.UnmarshalJSON([]byte("\"2018-06-05T12:10:06Z\"")))
	assert.Exactly(t, sec, m)
}

