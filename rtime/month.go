package rtime

import (
	"github.com/dustinevan/chron"
)

type IterationStrategy int

const (
	Standard IterationStrategy = iota
	Human
)

var DefaultIterationStrategy = Standard

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
	return MonthIterator{
		month:    m,
		curr:     seed,
		strategy: DefaultIterationStrategy,
		seedDay:  seed.AsTime().Day(),
	}
}

type MonthIterator struct {
	month    Month
	curr     chron.Time
	strategy IterationStrategy
	seedDay  int // [1- 31]
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

func (m MonthIterator) Since(t chron.Time) chron.Span {
	return chron.ZeroValue()
}

