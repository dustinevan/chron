package chron

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseUnixSeconds(t *testing.T) {
	d := NewDay(2016, 03, 17)
	ti, err := ParseUnixSeconds("1458172800")
	assert.True(t, err == nil)
	assert.Exactly(t, d.Time, ti)
}

func TestParseWithFormats(t *testing.T) {
	tt, err := ParseWithFormats("03-Feb-18")
	day := NewDay(2018, time.February, 3)
	assert.Nil(t, err)
	assert.Exactly(t, day.AsTime(), tt)
	assert.Exactly(t, day, DayOf(tt))

	tt, err = ParseWithFormats("03-Feb-2018")
	assert.Nil(t, err)
	assert.Exactly(t, day.AsTime(), tt)
	assert.Exactly(t, day, DayOf(tt))

	tt, err = ParseWithFormats("02-03-18")
	assert.Nil(t, err)
	assert.Exactly(t, day.AsTime(), tt)
	assert.Exactly(t, day, DayOf(tt))

	tt, err = ParseWithFormats("02-03-2018")
	assert.Nil(t, err)
	assert.Exactly(t, day.AsTime(), tt)
	assert.Exactly(t, day, DayOf(tt))

	tt, err = ParseWithFormats("02/03/18")
	assert.Nil(t, err)
	assert.Exactly(t, day.AsTime(), tt)
	assert.Exactly(t, day, DayOf(tt))

	tt, err = ParseWithFormats("02/03/2018")
	assert.Nil(t, err)
	assert.Exactly(t, day.AsTime(), tt)
	assert.Exactly(t, day, DayOf(tt))

	tt, err = ParseWithFormats("02/03/2018 3:13 PM")
	min := NewMinute(2018, time.February, 3, 15, 13)
	assert.Nil(t, err)
	assert.Exactly(t, min.AsTime(), tt)
	assert.Exactly(t, min, MinuteOf(tt))

	tt, err = ParseWithFormats("02/03/2018 3:13:52 PM")
	sec := NewSecond(2018, time.February, 3, 15, 13, 52)
	assert.Nil(t, err)
	assert.Exactly(t, sec.AsTime(), tt)
	assert.Exactly(t, sec, SecondOf(tt))

	tt, err = ParseWithFormats("02/03/2018 15:13")
	assert.Nil(t, err)
	assert.Exactly(t, min.AsTime(), tt)
	assert.Exactly(t, min, MinuteOf(tt))

	tt, err = ParseWithFormats("02/03/2018 15:13:52")
	assert.Nil(t, err)
	assert.Exactly(t, sec.AsTime(), tt)
	assert.Exactly(t, sec, SecondOf(tt))

	month := NewMonth(2018, time.February)
	tt, err = ParseWithFormats("Feb-2018")
	assert.Nil(t, err)
	assert.Exactly(t, month.AsTime(), tt)
	assert.Exactly(t, month, MonthOf(tt))

	tt, err = ParseWithFormats("Feb-18")
	assert.Nil(t, err)
	assert.Exactly(t, month.AsTime(), tt)
	assert.Exactly(t, month, MonthOf(tt))

	tt, err = ParseWithFormats("02-2018")
	assert.Nil(t, err)
	assert.Exactly(t, month.AsTime(), tt)
	assert.Exactly(t, month, MonthOf(tt))

	tt, err = ParseWithFormats("02-18")
	assert.Nil(t, err)
	assert.Exactly(t, month.AsTime(), tt)
	assert.Exactly(t, month, MonthOf(tt))

	tt, err = ParseWithFormats("02/18")
	assert.Nil(t, err)
	assert.Exactly(t, month.AsTime(), tt)
	assert.Exactly(t, month, MonthOf(tt))

	year := NewYear(2018)
	tt, err = ParseWithFormats("2018")
	assert.Nil(t, err)
	assert.Exactly(t, year.AsTime(), tt)
	assert.Exactly(t, year, YearOf(tt))

	tt, err = ParseWithFormats("2018-03-26 05:10:53.411453356 +0000 UTC")
	assert.Nil(t, err)
	assert.Exactly(t, NewTime(2018, time.March, 26, 5, 10, 53, 411453356), TimeOf(tt))

	tt, err = ParseWithFormats("2018-02-28")
	assert.Nil(t, err)
	assert.Exactly(t, NewDay(2018, time.February, 28).AsTime(), tt)
	assert.Exactly(t, NewDay(2018, time.February, 28), DayOf(tt))

	tt, err = ParseWithFormats("20180228T150219")
	assert.Nil(t, err)
	assert.Exactly(t, NewSecond(2018, time.February, 28, 15, 02, 19).AsTime(), tt)
	assert.Exactly(t, NewSecond(2018, time.February, 28, 15, 02, 19), SecondOf(tt))
}
