package chron

import (
	"github.com/dustinevan/chron/dura"
)

type SeqOption func(*Sequence) error

type Sequence struct {

	// -- Boundary and Flow --

	// Inclusive--this is the first time in the sequence. If begin does not
	// match the precision of the output time, the first output will be the
	// next occurrence of that time precision. e.g. a positive time flow
	// MinuteChan() seq, with begin := 2017-10-16 16:40:04.049121656 will
	// begin at 2017-10-16 16:41:00.0, neg time flow would begin at 16:40
	begin TimeExact

	// Exclusive--the timeExact of end is not included in the sequence.
	// If not set this defaults to date.MaxValue() or date.MinValue()
	// depending on inverseTime. A
	end TimeExact

	// true = the sequence goes back in time
	negativeTime bool

	// -- Increment Fields -- Note: Panics if one of these is not set.

	// constantIncrement is the length added to get the next seq time.
	increment dura.Time
	repeats   int

	// variableIncrementFn takes precedence over constantIncrement; The
	// sequence is incremented by the returned Length
	incrementFn func(exact TimeExact) TimeExact

	// -- Offset Fields --

	// if set, this length is added to the next seq time. Offset is not
	// used to calculate the next seq time. e.g. seq bases the next
	// date off curr not curr + offset.
	offset dura.Time
	// variableOffset takes precedence over constantOffset. A good example
	// use case is random offsets the spread the sequence times across an
	// hour.
	offsetFn func(Time) dura.Time

	// -- Channel Buffering --
	// Defaults to 1; set to 0 for unbuffered channels
	bsize int

	// -- Synchronization --
	// if true the sequence goroutine waits until the time of the next date
	// to pass it to the channel
	realtime bool
}

func NewSequence(begin Time, opts ...SeqOption) *Sequence {
	s := &Sequence{
		begin: begin.AsTimeExact(),
		bsize: 1,
		end:   MaxValue(),
	}
	for _, opt := range opts {
		err := opt(s)
		if err != nil {
			panic(err.Error())
		}
	}
	return s
}
