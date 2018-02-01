package chron

import (
	"github.com/stretchr/testify/assert"
	"time"
	"github.com/dustinevan/chron/dura"
	"database/sql/driver"
	"testing"
)

var tyear = time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)
var year = YearOf(tyear)

func TestNewYear(t *testing.T) {
	assert.Equal(t, tyear, NewYear(2018).Time)
	assert.Equal(t, tyear, YearOf(tyear).Time)
}

func TestThisYear(t *testing.T) {
	now := YearOf(time.Now())
	ty := ThisYear()
	assert.Equal(t, now, ty)
}

func TestYear_Transfers(t *testing.T) {
	ch := ThisYear()
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

func TestYear_Increment(t *testing.T) {
	y := year.Increment(dura.Duration{Year: 1, Month: 2, Day: 30, Dur: time.Second * 500})
	td := tyear.AddDate(1, 2, 30).Add(time.Second * 500)
	assert.Exactly(t, td, y.Time)
}

func TestYear_AsTime(t *testing.T) {
	assert.Exactly(t, tyear, year.AsTime())
}

func TestYear_Decrement(t *testing.T) {
	d := year.Decrement(dura.Duration{Year: 1, Month: 2, Day: 30, Dur: time.Second * 500})
	td := tyear.AddDate(-1, -2, -30).Add(time.Second * -500)
	assert.Exactly(t, td, d.Time)

}

func TestYear_AddN(t *testing.T) {
	assert.Exactly(t, NewYear(2020), year.AddN(2))
}

func TestYear_Start(t *testing.T) {
	assert.Exactly(t, year.Time, year.Start().Time)
}

func TestYear_End(t *testing.T) {
	assert.Exactly(t, year.Time.AddDate(1,0, 0).Add(-1 * time.Nanosecond), year.End().Time)
}

func TestYear_Contains(t *testing.T) {
	assert.True(t, year.Contains(NewMinute(2018, time.February, 1, 12, 45)))
	assert.False(t, year.Contains(year.AddN(1)))
}

func TestYear_Duration(t *testing.T) {
	assert.Equal(t, year.Duration().Years(), dura.Year.Years())
}

func TestYear_AddFns(t *testing.T) {
	assert.Exactly(t, year.AddYears(2), NewYear(2020))
	assert.Exactly(t, year.AddMonths(25), NewMonth(2020, time.February))
	assert.Exactly(t, year.AddDays(2), NewDay(2018, time.January, 3))
	assert.Exactly(t, year.AddHours(25), NewHour(2018, time.January, 2, 1))
	assert.Exactly(t, year.AddMinutes(72), NewMinute(2018, time.January, 1, 1, 12))
	assert.Exactly(t, year.AddSeconds(3672), NewSecond(2018, time.January, 1, 1, 1, 12))
	assert.Exactly(t, year.AddMillis(3672001), NewMilli(2018, time.January, 1, 1, 1, 12, 1))
	assert.Exactly(t, year.AddMicros(3672000001), NewMicro(2018, time.January, 1, 1, 1, 12, 1))
	assert.Exactly(t, year.AddNanos(3672000000001), NewTime(2018, time.January, 1, 1, 1, 12, 1))
}

func TestYear_Scan(t *testing.T) {
	var y Year
	assert.Nil(t, y.Scan(tyear))
	assert.Exactly(t, year, y)
	assert.Nil(t, y.Scan(nil))
	assert.Exactly(t, ZeroValue().AsYear(), y)
	assert.Error(t, y.Scan("wrong value"))
}

func TestYear_Value(t *testing.T) {
	v, err := year.Value()
	assert.Nil(t, err)
	assert.Exactly(t, v, driver.Value(tyear))
}

func TestYear_UnmarshalJSON(t *testing.T) {
	var y Year
	assert.Nil(t, y.UnmarshalJSON([]byte("null")))
	assert.Exactly(t, ZeroValue().AsYear(), y)
	assert.Error(t, y.UnmarshalJSON([]byte("as;dlkjfd")))
	assert.Exactly(t, ZeroValue().AsYear(), y)
	assert.Nil(t, y.UnmarshalJSON([]byte("\"2018-02-01T00:00:00Z\"")))
	assert.Exactly(t, year, y)
}