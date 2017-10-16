package dura

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var tdur = Duration{1, 3, 8, time.Second * 5040}

func TestNewDuration(t *testing.T) {
	assert.Exactly(t, tdur, NewDuration(1, 3, 8, time.Second*5040))
}

func TestYears(t *testing.T) {
	assert.Exactly(t, Duration{Yrs: 3}, Years(3))
}

func TestMonths(t *testing.T) {
	assert.Exactly(t, Duration{Mons: 3}, Months(3))
}

func TestDays(t *testing.T) {
	assert.Exactly(t, Duration{Dys: 5}, Days(5))
}

func TestHours(t *testing.T) {
	assert.Exactly(t, Duration{Dur: time.Hour * 12}, Hours(12))
}

func TestMins(t *testing.T) {
	assert.Exactly(t, Duration{Dur: time.Hour + time.Minute*30}, Mins(90))
}

func TestSecs(t *testing.T) {
	assert.Exactly(t, Duration{Dur: time.Minute + time.Second*30}, Secs(90))
}

func TestMillis(t *testing.T) {
	assert.Exactly(t, Duration{Dur: time.Second + time.Millisecond*275}, Millis(1275))
}

func TestMicros(t *testing.T) {
	assert.Exactly(t, Duration{Dur: time.Second + time.Microsecond*5}, Micros(1000005))
}

func TestNanos(t *testing.T) {
	assert.Exactly(t, Duration{Dur: time.Millisecond*5 + time.Nanosecond*73}, Nanos(5000073))
}

func TestDuration_Mult(t *testing.T) {
	assert.Exactly(t, NewDuration(3, 9, 24, time.Second*15120), tdur.Mult(3))
}

func TestSum(t *testing.T) {
	assert.Exactly(t, NewDuration(3, 9, 24, time.Second*15120), Sum(tdur, tdur, tdur))
}

func TestDuration_Years(t *testing.T) {
	assert.Exactly(t, 1, tdur.Years())
}

func TestDuration_Months(t *testing.T) {
	assert.Exactly(t, 3, tdur.Months())
}

func TestDuration_Days(t *testing.T) {
	assert.Exactly(t, 8, tdur.Days())
}

func TestDuration_Duration(t *testing.T) {
	assert.Exactly(t, time.Second*5040, tdur.Duration())
}

func TestDuration_String(t *testing.T) {
	assert.Exactly(t, "1y3m8d1h24m0s", tdur.String())
}

func TestUnit_Years(t *testing.T) {
	assert.Exactly(t, 0, Zero.Years())
	assert.Exactly(t, 100, Century.Years())
	assert.Exactly(t, 10, Decade.Years())
	assert.Exactly(t, 1, Year.Years())
	assert.Exactly(t, 0, Quarter.Years())
	assert.Exactly(t, 0, Month.Years())
	assert.Exactly(t, 0, Week.Years())
	assert.Exactly(t, 0, Day.Years())
	assert.Exactly(t, 0, Hour.Years())
	assert.Exactly(t, 0, Minute.Years())
	assert.Exactly(t, 0, Second.Years())
	assert.Exactly(t, 0, Milli.Years())
	assert.Exactly(t, 0, Micro.Years())
	assert.Exactly(t, 0, Nano.Years())
}

func TestUnit_Months(t *testing.T) {
	assert.Exactly(t, 0, Zero.Months())
	assert.Exactly(t, 0, Century.Months())
	assert.Exactly(t, 0, Decade.Months())
	assert.Exactly(t, 0, Year.Months())
	assert.Exactly(t, 3, Quarter.Months())
	assert.Exactly(t, 1, Month.Months())
	assert.Exactly(t, 0, Week.Months())
	assert.Exactly(t, 0, Day.Months())
	assert.Exactly(t, 0, Hour.Months())
	assert.Exactly(t, 0, Minute.Months())
	assert.Exactly(t, 0, Second.Months())
	assert.Exactly(t, 0, Milli.Months())
	assert.Exactly(t, 0, Micro.Months())
	assert.Exactly(t, 0, Nano.Months())
}

func TestUnit_Days(t *testing.T) {
	assert.Exactly(t, 0, Zero.Days())
	assert.Exactly(t, 0, Century.Days())
	assert.Exactly(t, 0, Decade.Days())
	assert.Exactly(t, 0, Year.Days())
	assert.Exactly(t, 0, Quarter.Days())
	assert.Exactly(t, 0, Month.Days())
	assert.Exactly(t, 7, Week.Days())
	assert.Exactly(t, 1, Day.Days())
	assert.Exactly(t, 0, Hour.Days())
	assert.Exactly(t, 0, Minute.Days())
	assert.Exactly(t, 0, Second.Days())
	assert.Exactly(t, 0, Milli.Days())
	assert.Exactly(t, 0, Micro.Days())
	assert.Exactly(t, 0, Nano.Days())
}

func TestUnit_Duration(t *testing.T) {
	assert.Exactly(t, time.Duration(0), Zero.Duration())
	assert.Exactly(t, time.Duration(0), Century.Duration())
	assert.Exactly(t, time.Duration(0), Decade.Duration())
	assert.Exactly(t, time.Duration(0), Year.Duration())
	assert.Exactly(t, time.Duration(0), Quarter.Duration())
	assert.Exactly(t, time.Duration(0), Month.Duration())
	assert.Exactly(t, time.Duration(0), Week.Duration())
	assert.Exactly(t, time.Duration(0), Day.Duration())
	assert.Exactly(t, time.Hour, Hour.Duration())
	assert.Exactly(t, time.Minute, Minute.Duration())
	assert.Exactly(t, time.Second, Second.Duration())
	assert.Exactly(t, time.Millisecond, Milli.Duration())
	assert.Exactly(t, time.Microsecond, Micro.Duration())
	assert.Exactly(t, time.Nanosecond, Nano.Duration())
}

func TestUnit_String(t *testing.T) {
	assert.Exactly(t, "Zero Unit", Zero.String())
	assert.Exactly(t, "Century", Century.String())
	assert.Exactly(t, "Decade", Decade.String())
	assert.Exactly(t, "Year", Year.String())
	assert.Exactly(t, "Quarter", Quarter.String())
	assert.Exactly(t, "Month", Month.String())
	assert.Exactly(t, "Week", Week.String())
	assert.Exactly(t, "Day", Day.String())
	assert.Exactly(t, "Hour", Hour.String())
	assert.Exactly(t, "Minute", Minute.String())
	assert.Exactly(t, "Second", Second.String())
	assert.Exactly(t, "Milli", Milli.String())
	assert.Exactly(t, "Micro", Micro.String())
	assert.Exactly(t, "Nano", Nano.String())
}
