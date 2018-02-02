package chron

import (
	"github.com/stretchr/testify/assert"
	"github.com/dustinevan/chron/dura"
	"time"
	"database/sql/driver"
	"testing"
)

var tmicro = time.Date(2018, time.June, 5, 12, 10, 6, 55000, time.UTC)
var micro = MicroOf(tmicro)

func TestNewMicro(t *testing.T) {
	assert.Equal(t, tmicro, NewMicro(2018, time.June, 5, 12, 10, 6, 55).Time)
	assert.Equal(t, tmicro, MicroOf(tmicro).Time)
}

func TestMicro_Transfers(t *testing.T) {
	ch := ThisMicro()
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

func TestMicro_Increment(t *testing.T) {
	y := micro.Increment(dura.NewDuration(1, 2, 30, time.Microsecond * 500))
	td := tmicro.AddDate(1, 2, 30).Add(time.Microsecond * 500)
	assert.Exactly(t, td, y.Time)
}

func TestMicro_AsTime(t *testing.T) {
	assert.Exactly(t, tmicro, micro.AsTime())
}

func TestMicro_Decrement(t *testing.T) {
	d := micro.Decrement(dura.NewDuration(1, 2, 30, time.Microsecond * 500))
	td := tmicro.AddDate(-1, -2, -30).Add(time.Microsecond * -500)
	assert.Exactly(t, td, d.Time)

}

func TestMicro_AddN(t *testing.T) {
	assert.Exactly(t, NewMicro(2018, time.June, 5, 12, 10, 6, 58), micro.AddN(3))
}

func TestMicro_Start(t *testing.T) {
	assert.Exactly(t, micro.Time, micro.Start().Time)
}

func TestMicro_End(t *testing.T) {
	assert.Exactly(t, micro.Time.Add(time.Microsecond + (-1 * time.Nanosecond)), micro.End().Time)
}

func TestMicro_Contains(t *testing.T) {
	assert.True(t, micro.Contains(NewTime(2018, time.June, 5, 12, 10, 6, 55456)))
	assert.False(t, micro.Contains(micro.AddN(1)))
}

func TestMicro_Duration(t *testing.T) {
	assert.Equal(t, micro.Duration().Duration(), dura.Micro.Duration())
}

func TestMicro_AddFns(t *testing.T) {
	assert.Exactly(t, micro.AddYears(2), NewMicro(2020, time.June, 5, 12, 10, 6, 55))
	assert.Exactly(t, micro.AddMonths(25), NewMicro(2020, time.July, 5, 12, 10,6, 55))
	assert.Exactly(t, micro.AddDays(2), NewMicro(2018, time.June, 7, 12, 10, 6, 55))
	assert.Exactly(t, micro.AddHours(25), NewMicro(2018, time.June, 6, 13, 10, 6, 55))
	assert.Exactly(t, micro.AddMinutes(72), NewMicro(2018, time.June, 5, 13, 22, 6, 55))
	assert.Exactly(t, micro.AddSeconds(3672), NewMicro(2018, time.June, 5, 13, 11, 18, 55))
	assert.Exactly(t, micro.AddMillis(3672001), NewMicro(2018, time.June, 5, 13, 11, 18, 1055))
	assert.Exactly(t, micro.AddMicros(3672000001), NewMicro(2018, time.June, 5, 13, 11, 18, 56))
	assert.Exactly(t, micro.AddNanos(3672000000001), NewTime(2018, time.June, 5, 13, 11, 18, 55001))
}

func TestMicro_Scan(t *testing.T) {
	var m Micro
	assert.Nil(t, m.Scan(tmicro))
	assert.Exactly(t, micro, m)
	assert.Nil(t, m.Scan(nil))
	assert.Exactly(t, ZeroValue().AsMicro(), m)
	assert.Error(t, m.Scan("wrong value"))
}

func TestMicro_Value(t *testing.T) {
	v, err := micro.Value()
	assert.Nil(t, err)
	assert.Exactly(t, v, driver.Value(tmicro))
}

func TestMicro_UnmarshalJSON(t *testing.T) {
	var m Micro
	assert.Nil(t, m.UnmarshalJSON([]byte("null")))
	assert.Exactly(t, ZeroValue().AsMicro(), m)
	assert.Error(t, m.UnmarshalJSON([]byte("as;dlkjfd")))
	assert.Exactly(t, ZeroValue().AsMicro(), m)
	assert.Nil(t, m.UnmarshalJSON([]byte("\"2018-06-05T12:10:06.000055Z\"")))
	assert.Exactly(t, micro, m)
}
