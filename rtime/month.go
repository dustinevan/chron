package rtime

import (
	"github.com/dustinevan/chron"
)

type Month int

const (
	January Month = iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func (m Month) Iterator(seed chron.Time) Iterator {
	return MonthIterator{month: m, curr: seed}
}

type MonthIterator struct {
	month Month
	curr  chron.Time
}

func (m MonthIterator) Nth(n int) chron.Span {
	return m.curr.AsMonth().AddN(n)
}

func (m MonthIterator) Prev(t chron.Time) chron.Span {
	return m.Nth(-1)
}

func (m MonthIterator) Next(t chron.Time) chron.Span {
	return m.Nth(1)
}

func (m MonthIterator) Since()

func (m Month) LastDay(y chron.Year) chron.Day {

}
