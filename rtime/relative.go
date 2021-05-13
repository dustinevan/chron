package rtime

import (
	"github.com/dustinevan/chron"
)

type Relative func(chron.Chron) (chron.Chron, bool)

// func Months(m ...time.Month) Relative {
// 	return func(exact chron.Chron) (chron.Chron, bool) {
// 		thisM := exact.Month()
// 		nextM := thisM
// 		for mon := range m {
//
// 		}
// 	}
// }

type Month func(y chron.Time) chron.Month
type MonthDay func(y chron.Year) chron.Day
type DayOfMonth func(m chron.Month) chron.Day
type HourOfDay func(d chron.Time) chron.Hour
type TimeOfDay func(d chron.Time) chron.Minute
type MinOfHour func(d chron.Time) chron.Minute
type SecOfMin func(d chron.Time) chron.Second

// func NewTimeOfDay(hr, min int) TimeOfDay {
// 	return
// }

