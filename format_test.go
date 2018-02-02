package chron

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
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
}