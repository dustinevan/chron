package chron

import (
	"github.com/stretchr/testify/assert"
	"github.com/dustinevan/chron/dura"
	"time"
	"database/sql/driver"
	"testing"
)

var thour = time.Date(2018, time.June, 5, 12, 0, 0, 0, time.UTC)
var hour = HourOf(thour)

func TestNewHour(t *testing.T) {
	assert.Equal(t, thour, NewHour(2018, time.June, 5, 12).Time)
	assert.Equal(t, thour, HourOf(thour).Time)
}

func TestThisHour(t *testing.T) {
	now := HourOf(time.Now())
	ty := ThisHour()
	assert.Equal(t, now, ty)
}

func TestHour_Transfers(t *testing.T) {
	ch := ThisHour()
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

func TestHour_Increment(t *testing.T) {
	y := hour.Increment(dura.Duration{Year: 1, Month: 2, Day: 30, Dur: time.Second * 500})
	td := thour.AddDate(1, 2, 30).Add(time.Second * 500)
	assert.Exactly(t, td, y.Time)
}

func TestHour_AsTime(t *testing.T) {
	assert.Exactly(t, thour, hour.AsTime())
}

func TestHour_Decrement(t *testing.T) {
	d := hour.Decrement(dura.Duration{Year: 1, Month: 2, Day: 30, Dur: time.Second * 500})
	td := thour.AddDate(-1, -2, -30).Add(time.Second * -500)
	assert.Exactly(t, td, d.Time)

}

func TestHour_AddN(t *testing.T) {
	assert.Exactly(t, NewHour(2018, time.June, 5, 14), hour.AddN(2))
}

func TestHour_Start(t *testing.T) {
	assert.Exactly(t, hour.Time, hour.Start().Time)
}

func TestHour_End(t *testing.T) {
	assert.Exactly(t, hour.Time.Add(time.Hour + (-1 * time.Nanosecond)), hour.End().Time)
}

func TestHour_Contains(t *testing.T) {
	assert.True(t, hour.Contains(NewMinute(2018, time.June, 5, 12, 45)))
	assert.False(t, hour.Contains(hour.AddN(1)))
}

func TestHour_Duration(t *testing.T) {
	assert.Equal(t, hour.Duration().Duration(), dura.Hour.Duration())
}

func TestHour_AddFns(t *testing.T) {
	assert.Exactly(t, hour.AddYears(2), NewHour(2020, time.June, 5, 12))
	assert.Exactly(t, hour.AddMonths(25), NewHour(2020, time.July, 5, 12))
	assert.Exactly(t, hour.AddDays(2), NewHour(2018, time.June, 7, 12))
	assert.Exactly(t, hour.AddHours(25), NewHour(2018, time.June, 6, 13))
	assert.Exactly(t, hour.AddMinutes(72), NewMinute(2018, time.June, 5, 13, 12))
	assert.Exactly(t, hour.AddSeconds(3672), NewSecond(2018, time.June, 5, 13, 1, 12))
	assert.Exactly(t, hour.AddMillis(3672001), NewMilli(2018, time.June, 5, 13, 1, 12, 1))
	assert.Exactly(t, hour.AddMicros(3672000001), NewMicro(2018, time.June, 5, 13, 1, 12, 1))
	assert.Exactly(t, hour.AddNanos(3672000000001), NewTime(2018, time.June, 5, 13, 1, 12, 1))
}

func TestHour_Scan(t *testing.T) {
	var h Hour
	assert.Nil(t, h.Scan(thour))
	assert.Exactly(t, hour, h)
	assert.Nil(t, h.Scan(nil))
	assert.Exactly(t, ZeroValue().AsHour(), h)
	assert.Error(t, h.Scan("wrong value"))
}

func TestHour_Value(t *testing.T) {
	v, err := hour.Value()
	assert.Nil(t, err)
	assert.Exactly(t, v, driver.Value(thour))
}

func TestHour_UnmarshalJSON(t *testing.T) {
	var h Hour
	assert.Nil(t, h.UnmarshalJSON([]byte("null")))
	assert.Exactly(t, ZeroValue().AsHour(), h)
	assert.Error(t, h.UnmarshalJSON([]byte("as;dlkjfd")))
	assert.Exactly(t, ZeroValue().AsHour(), h)
	assert.Nil(t, h.UnmarshalJSON([]byte("\"2018-06-05T12:00:00Z\"")))
	assert.Exactly(t, hour, h)
}

