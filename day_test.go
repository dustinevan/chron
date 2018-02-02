package chron

import (
	"testing"
	"time"

	"github.com/dustinevan/chron/dura"
	"github.com/stretchr/testify/assert"
	"database/sql/driver"
)

var tday = time.Date(2018, time.February, 1, 0, 0, 0, 0, time.UTC)
var day = DayOf(tday)

func TestNewDay(t *testing.T) {
	assert.Equal(t, tday, NewDay(2018, time.February, 1).Time)
	assert.Equal(t, tday, DayOf(tday).Time)
}

func TestToday(t *testing.T) {
	now := DayOf(time.Now())
	today := Today()
	assert.Equal(t, now, today)
}

func TestDay_Transfers(t *testing.T) {
	ch := Today()
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

func TestDay_Increment(t *testing.T) {
	d := day.Increment(dura.Duration{Yrs: 1, Mons: 2, Dys: 30, Dur: time.Second * 500})
	td := tday.AddDate(1, 2, 30).Add(time.Second * 500)
	assert.Exactly(t, td, d.Time)
}

func TestDay_AsTime(t *testing.T) {
	assert.Exactly(t, tday, day.AsTime())
}

func TestDay_Decrement(t *testing.T) {
	d := day.Decrement(dura.Duration{Yrs: 1, Mons: 2, Dys: 30, Dur: time.Second * 500})
	td := tday.AddDate(-1, -2, -30).Add(time.Second * -500)
	assert.Exactly(t, td, d.Time)

}

func TestDay_AddN(t *testing.T) {
	assert.Exactly(t, NewDay(2018, time.February, 3), day.AddN(2))
}

func TestDay_Start(t *testing.T) {
	assert.Exactly(t, day.Time, day.Start().Time)
}

func TestDay_End(t *testing.T) {
	assert.Exactly(t, day.Time.Add((time.Hour * 24) - time.Nanosecond), day.End().Time)
}

func TestDay_Contains(t *testing.T) {
	assert.True(t, day.Contains(NewMinute(2018, time.February, 1, 12, 45)))
	assert.False(t, day.Contains(day.AddN(1)))
}

func TestDay_Duration(t *testing.T) {
	assert.Equal(t, day.Duration().Days(), dura.Day.Days())
}

func TestDay_AddFns(t *testing.T) {
	assert.Exactly(t, day.AddYears(2), NewDay(2020, time.February, 1))
	assert.Exactly(t, day.AddMonths(24), NewDay(2020, time.February, 1))
	assert.Exactly(t, day.AddDays(2), day.AddN(2))
	assert.Exactly(t, day.AddHours(25), NewHour(2018, time.February, 2, 1))
	assert.Exactly(t, day.AddMinutes(72), NewMinute(2018, time.February, 1, 1, 12))
	assert.Exactly(t, day.AddSeconds(3672), NewSecond(2018, time.February, 1, 1, 1, 12))
	assert.Exactly(t, day.AddMillis(3672001), NewMilli(2018, time.February, 1, 1, 1, 12, 1))
	assert.Exactly(t, day.AddMicros(3672000001), NewMicro(2018, time.February, 1, 1, 1, 12, 1))
	assert.Exactly(t, day.AddNanos(3672000000001), NewTime(2018, time.February, 1, 1, 1, 12, 1))
}

func TestDay_Scan(t *testing.T) {
	var d Day
	assert.Nil(t, d.Scan(tday))
	assert.Exactly(t, day, d)
	assert.Nil(t, d.Scan(nil))
	assert.Exactly(t, ZeroValue().AsDay(), d)
	assert.Error(t, d.Scan("wrong value"))
}

func TestDay_Value(t *testing.T) {
	v, err := day.Value()
	assert.Nil(t, err)
	assert.Exactly(t, v, driver.Value(tday))
}

func TestDay_UnmarshalJSON(t *testing.T) {
	var d Day
	assert.Nil(t, d.UnmarshalJSON([]byte("null")))
	assert.Exactly(t, ZeroValue().AsDay(), d)
	assert.Error(t, d.UnmarshalJSON([]byte("as;dlkjfd")))
	assert.Exactly(t, ZeroValue().AsDay(), d)
	assert.Nil(t, d.UnmarshalJSON([]byte("\"2018-02-01T00:00:00Z\"")))
	assert.Exactly(t, day, d)
}