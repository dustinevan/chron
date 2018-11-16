package rtime

import (
	"time"

	"github.com/dustinevan/chron"
)

type RMonth func(year chron.Year) chron.Month
type RDate func(year chron.Year) chron.Day
type RDay func(month chron.Month) chron.Day
type RHour func(day chron.Day) chron.Hour
type RMinute func(hour chron.Hour) chron.Minute
type RSecond func(minute chron.Minute) chron.Second
type RMilli func(second chron.Second) chron.Milli
type RMicro func(milli chron.Micro) chron.Micro
type RTime func(chron.Time) chron.TimeExact

func RMonthOf(month time.Month) RMonth {
	return func(year chron.Year) chron.Month {
		return chron.NewMonth(year.Year(), month)
	}
}
