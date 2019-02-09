package rtime

import (
	"github.com/dustinevan/chron"
)

type RMonth func(y chron.Year) chron.Month
type RDate func(y chron.Year) chron.Day
type RDay func(m chron.Month) chron.Day
type RHour func(d chron.Day) chron.Hour
type RMinute func(h chron.Hour) chron.Minute
type RSecond func(m chron.Minute) chron.Second
type RMilli func(s chron.Second) chron.Milli
type RMicro func(m chron.Milli) chron.Micro
type RTime func(t chron.Time) chron.Chron

type TimeFilter func(t chron.Time) bool
