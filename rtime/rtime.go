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
	// Returns the nth occurrence and advances the internal time (if n is negative it moves in reverse)
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
