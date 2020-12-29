package rtime

import (
	"github.com/dustinevan/chron"
)

// rtime.RTime represents a relative time, or a time that isn't know until it is
// provided a definite reference time. e.g. 'Thursday', 'Thanksgiving', or '12:15pm'
// when given a seed time, relative time's produce an iterator that can return
// time spans based on an nth occurrence, like the forth wednesday in September 2018
type RTime interface {
	Iterator(seed chron.Time) Iterator
}

// Iterator implementations have access to an underlying definite seed time.
// As Next() is called the iterator returns the next sequential timespan
// related to the underlying implementation details. For example, an iterator
// that represents Tuesdays would return a chron.Span of each chron.Day where
// the weekday is Tuesday.
type Iterator interface {
	// Returns the nth occurrence of the span without advancing
	Nth(n int) chron.Span

	// Returns the next occurrence and advances the internal time
	Next() chron.Span

	// Returns the previous occurrence and decrements the internal time
	Prev() chron.Span

	// Returns the span between the current time and the proceeding nth occurrence
	Since(n int) chron.Span

	// Returns the span between the current and the nth occurrence
	Until(n int) chron.Span
}

//Oldstuff
//
//type RMonth func(y chron.Year) chron.Month
//type RDate func(y chron.Year) chron.Day
//type RDay func(m chron.Month) chron.Day
//type RHour func(d chron.Day) chron.Hour
//type RMinute func(h chron.Hour) chron.Minute
//type RSecond func(m chron.Minute) chron.Second
//type RMilli func(s chron.Second) chron.Milli
//type RMicro func(m chron.Milli) chron.Micro
//type RTime func(t chron.Time) chron.Chron
//
//type TimeFilter func(t chron.Time) bool
//
//type Series struct {
//	timeFunc RTime
//	start chron.Time
//	stop TimeFilter
//}
//
//func NewSeries(timeFunc RTime, start chron.Time, stop TimeFilter) Series {
//	return Series{timeFunc:timeFunc, start: start, stop: stop}
//}
//
//func (s *Series) NextN(n int) []chron.Chron {
//	result := make([]chron.Chron, n)
//	for i := 0; i < n; i++ {
//		s.start = s.timeFunc(s.start)
//		if s.stop(s.start) {
//			return result
//		}
//		result[i] = s.start.AsChron()
//	}
//	return result
//}
