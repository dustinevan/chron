package seq

import (
	"context"

	"time"

	"github.com/dustinevan/chron"
	"github.com/dustinevan/chron/length"
)

type sequence struct {

	// -- Boundary and Flow --

	// Inclusive--this is the first time in the sequence. If begin does not
	// match the precision of the output time, the first output will be the
	// next occurrence of that time precision. e.g. a positive time flow
	// MinuteChan() seq, with begin := 2017-10-16 16:40:04.049121656 will
	// begin at 2017-10-16 16:41:00.0, neg time flow would begin at 16:40
	begin chron.TimeExact

	// Exclusive--the timeExact of end is not included in the sequence.
	// If not set this defaults to date.MaxValue() or date.MinValue()
	// depending on inverseTime. A
	end chron.TimeExact

	// true = the sequence goes back in time
	negativeTime bool

	// -- Increment Fields -- Note: Panics if one of these is not set.

	// constantIncrement is the length added to get the next seq time.
	increment chron.Length
	incn      int
	repeats   int

	// variableIncrementFn takes precedence over constantIncrement; The
	// sequence is incremented by the returned Length
	incrementFn func(exact chron.TimeExact) chron.TimeExact

	// -- Offset Fields --

	// if set, this length is added to the next seq time. Offset is not
	// used to calculate the next seq time. e.g. seq bases the next
	// date off curr not curr + offset.
	offset chron.Length
	offn   int
	// variableOffset takes precedence over constantOffset. A good example
	// use case is random offsets the spread the sequence times across an
	// hour.
	offsetFn func(chron.Time) chron.Length

	// -- Channel Buffering --
	// Defaults to 8; set to 0 for unbuffered channels
	bsize int

	// -- Synchronization --
	// if true the sequence goroutine waits until the time of the next date
	// to pass it to the channel
	realtime bool
}

func Sequence(begin chron.Time) *sequence {
	return &sequence{
		begin: begin.AsTimeExact(),
		bsize: 8,
		end:   chron.MaxValue(),
	}
}

func (s *sequence) End(end chron.Time) *sequence {
	s.end = end.AsTimeExact()
	return s
}

func (s *sequence) EndIncl(end chron.Time) *sequence {
	s.end = end.Increment(length.Nano).AsTimeExact()
	return s
}

func (s *sequence) Length(len chron.Length) *sequence {
	s.end = s.begin.Increment(len)
	return s
}

func (s *sequence) Increment(len chron.Length) *sequence {
	s.increment = len
	return s
}

func (s *sequence) IncrementFn(f func(chron.TimeExact) chron.TimeExact) *sequence {
	s.incrementFn = f
	return s
}

// Repeats the same time n times before moving to the next
func (s *sequence) Repeats(n int) {
	s.repeats = n
}

func (s *sequence) Offset(len chron.Length) *sequence {
	s.offset = len
	return s
}

func (s *sequence) OffsetFn(f func(chron.Time) chron.Length) *sequence {
	s.offsetFn = f
	return s
}

func (s *sequence) RealTime() *sequence {
	s.realtime = true
	return s
}

func (s *sequence) ChanSize(i int) *sequence {
	s.bsize = i
	return s
}

type stop func()

func (s *sequence) TimeChan() (<-chan time.Time, stop) {
	in, canc := s.start()
	out := make(chan time.Time, s.bsize)
	stop := func() {
		canc()
		for t := range out {
			seeya(t)
		}
	}
	go func() {
		defer close(out)
		for t := range in {
			out <- t.AsTime()
		}
	}()
	return out, stop
}

func (s *sequence) ExactTimeChan() (<-chan chron.TimeExact, stop) {
	in, canc := s.start()
	out := make(chan chron.TimeExact, s.bsize)
	stop := func() {
		canc()
		for t := range out {
			goodbye(t)
		}
	}
	go func() {
		defer close(out)
		for d := range in {
			out <- d
		}
	}()
	return out, stop
}

func (s *sequence) YearChan() (<-chan chron.Year, stop) {
	in, canc := s.start()
	out := make(chan chron.Year, s.bsize)
	stop := func() {
		canc()
		for t := range out {
			goodbye(t)
		}
	}
	go func() {
		defer close(out)
		for d := range in {
			out <- d.AsYear()
		}
	}()
	return out, stop
}

func (s *sequence) MonthChan() (<-chan chron.Month, stop) {
	in, canc := s.start()
	out := make(chan chron.Month, s.bsize)
	stop := func() {
		canc()
		for t := range out {
			goodbye(t)
		}
	}
	go func() {
		defer close(out)
		for d := range in {
			out <- d.AsMonth()
		}
	}()
	return out, stop
}

func (s *sequence) DayChan() (<-chan chron.Day, stop) {
	in, canc := s.start()
	out := make(chan chron.Day, s.bsize)
	stop := func() {
		canc()
		for t := range out {
			goodbye(t)
		}
	}
	go func() {
		defer close(out)
		for d := range in {
			out <- d.AsDay()
		}
	}()
	return out, stop
}

func (s *sequence) HourChan() (<-chan chron.Hour, stop) {
	in, canc := s.start()
	out := make(chan chron.Hour, s.bsize)
	stop := func() {
		canc()
		for t := range out {
			goodbye(t)
		}
	}
	go func() {
		defer close(out)
		for d := range in {
			out <- d.AsHour()
		}
	}()
	return out, stop
}

func (s *sequence) MinuteChan() (<-chan chron.Minute, stop) {
	in, canc := s.start()
	out := make(chan chron.Minute, s.bsize)
	stop := func() {
		canc()
		for t := range out {
			goodbye(t)
		}
	}
	go func() {
		defer close(out)
		for d := range in {
			out <- d.AsMinute()
		}
	}()
	return out, stop
}

func (s *sequence) SecondChan() (<-chan chron.Second, stop) {
	in, canc := s.start()
	out := make(chan chron.Second, s.bsize)
	stop := func() {
		canc()
		for t := range out {
			goodbye(t)
		}
	}
	go func() {
		defer close(out)
		for d := range in {
			out <- d.AsSecond()
		}
	}()
	return out, stop
}

func goodbye(t chron.Time) {}

func seeya(t time.Time) {}

func (s *sequence) start() (<-chan chron.TimeExact, context.CancelFunc) {

	// setup internal context and cancel
	ctx, canc := context.WithCancel(context.Background())

	var out <-chan chron.TimeExact
	// pick the right goroutine to start up based sequence fields
	if s.incrementFn != nil {
		out = s.variableIncs(ctx)
	} else if s.increment != nil {
		out = s.fixedIncs(ctx)
	} else {
		panic("invalid sequence initialization. an increment method must be supplied")
	}

	if s.realtime {
		rtchan := make(chan chron.TimeExact, s.bsize)
		go func() {
			defer close(rtchan)
			for t := range out {
				until := time.Until(t.Time)
				if until > 0 {
					time.Sleep(until)
				}
				rtchan <- t
			}
		}()
		return rtchan, canc
	}
	return out, canc
}

// One of the six possible goroutines starts
func (s *sequence) fixedIncs(ctx context.Context) <-chan chron.TimeExact {
	out := make(chan chron.TimeExact, s.bsize)
	if s.offsetFn != nil {
		if s.begin.After(s.end.Time) {
			// back in time, interval incremented, with dynamic offset fn
			go func() {
				defer close(out)
				curr := s.begin
				for curr.After(s.end.Time) {
					if ctx.Err() != nil {
						return
					}
					out <- curr.Increment(s.offsetFn(curr))
					curr = curr.Decrement(s.increment)
				}
			}()
			return out
		} else {
			// forward in time, interval incremented, with dynamic offset fn
			go func() {
				defer close(out)
				curr := s.begin
				for curr.Before(s.end.Time) {
					if ctx != nil {
						return
					}
					out <- curr.Increment(s.offsetFn(curr))
					curr = curr.Increment(s.increment)
				}
			}()
			return out
		}
	}
	if s.offset != nil {
		if s.begin.After(s.end.Time) {
			// back in time, interval incremented, with a fixed offset
			go func() {
				defer close(out)
				curr := s.begin
				for curr.After(s.end.Time) {
					if ctx.Err() != nil {
						return
					}
					out <- curr.Increment(s.offset)
					curr = curr.Decrement(s.increment)
				}
			}()
			return out
		} else {
			// forward in time, interval incremented, with a fixed offset
			go func() {
				defer close(out)
				curr := s.begin
				for curr.Before(s.end.Time) {
					if ctx.Err() != nil {
						return
					}
					out <- curr.Increment(s.offset)
					curr = curr.Increment(s.increment)
				}
			}()
			return out
		}
	}

	if s.begin.After(s.end.Time) {
		// back in time, interval incremented, no offset
		go func() {
			defer close(out)
			curr := s.begin
			for curr.After(s.end.Time) {
				if ctx.Err() != nil {
					return
				}
				out <- curr
				curr = curr.Decrement(s.increment)
			}
		}()
		return out
	} else {
		// forward in time, interval incremented, no offset
		go func() {
			defer close(out)
			curr := s.begin
			for curr.Before(s.end.Time) {
				if ctx.Err() != nil {
					return
				}
				out <- curr
				curr = curr.Increment(s.increment)
			}
		}()
		return out
	}
}

// One of the six possible goroutines starts
func (s *sequence) variableIncs(ctx context.Context) <-chan chron.TimeExact {
	out := make(chan chron.TimeExact, s.bsize)
	if s.offsetFn != nil {
		if s.begin.After(s.end.Time) {
			// back in time, interval incremented, with dynamic offset fn
			go func() {
				defer close(out)
				curr := s.begin
				for curr.After(s.end.Time) {
					if ctx.Err() != nil {
						return
					}
					out <- curr.Increment(s.offsetFn(curr))
					curr = s.incrementFn(curr)
				}
			}()
			return out
		} else {
			// forward in time, interval incremented, with dynamic offset fn
			go func() {
				defer close(out)
				curr := s.begin
				for curr.Before(s.end.Time) {
					if ctx != nil {
						return
					}
					out <- curr.Increment(s.offsetFn(curr))
					curr = s.incrementFn(curr)
				}
			}()
			return out
		}
	}
	if s.offset != nil {
		if s.begin.After(s.end.Time) {
			// back in time, interval incremented, with a fixed offset
			go func() {
				defer close(out)
				curr := s.begin
				for curr.After(s.end.Time) {
					if ctx.Err() != nil {
						return
					}
					out <- curr.Increment(s.offset)
					curr = s.incrementFn(curr)
				}
			}()
			return out
		} else {
			// forward in time, interval incremented, with a fixed offset
			go func() {
				defer close(out)
				curr := s.begin
				for curr.Before(s.end.Time) {
					if ctx.Err() != nil {
						return
					}
					out <- curr.Increment(s.offset)
					curr = s.incrementFn(curr)
				}
			}()
			return out
		}
	}

	if s.begin.After(s.end.Time) {
		// back in time, interval incremented, no offset
		go func() {
			defer close(out)
			curr := s.begin
			for curr.After(s.end.Time) {
				if ctx.Err() != nil {
					return
				}
				out <- curr
				curr = s.incrementFn(curr)
			}
		}()
		return out
	} else {
		// forward in time, interval incremented, no offset
		go func() {
			defer close(out)
			curr := s.begin
			for curr.Before(s.end.Time) {
				if ctx.Err() != nil {
					return
				}
				out <- curr
				curr = s.incrementFn(curr)
			}
		}()
		return out
	}
}
