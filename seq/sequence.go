package seq

import (
	"context"

	"time"

	"github.com/dustinevan/time/chron"
	"github.com/dustinevan/time/date"
)

type DynamicIncrementFunc func() chron.Length
type DynamicOffsetFunc func(from chron.Time) chron.Time

type sequence struct {

	// -- Boundary and Flow --
	// the start time. if it is more precise than the
	// output time unit, the first output will be the next
	// occurrence after begin
	begin chron.Time

	// exclusive. If not set this defaults to date.MaxValue() or date.MinValue()
	// depending on inverseTime
	end chron.Time

	// true = the sequence goes back in time
	inverseTime bool

	// -- Increment Fields -- Note: Panics if one of these is not set.
	// incUnit is the simple case, it increments by a single time unit
	incLength chron.Length
	incN      int

	// incDynamic takes precedence over incUnit; The sequence is
	// incremented by the returned Interval
	incDynamic DynamicIncrementFunc

	// -- Offset Fields --
	// this interval is added to the outputted date before it is sent.
	// the sequence does not base the next date off this value e.g.
	// the current date does not include offset
	offset chron.Length
	// offsetDynamic takes precedence over offset. A good example use
	// case is random offsets the spread the sequence times across an
	// hour.
	offsetDynamic DynamicOffsetFunc

	// -- Cancellation
	// the parent context, when cancelled the sequencing goroutine will
	// close its channel and return
	ctx context.Context

	// if ctx is not nil, this is a child of ctx, cancels to ctx will also
	// cancel internalCtx. If ctx is nil, this is a child of the background
	// context. calls to Stop will cancel this context, the sequencing
	// goroutine will close it's channel and return
	internalCtx context.Context
	// cancel func used by Stop, clients are expected to keep reading
	// the outgoing channel is closed
	internalCanc context.CancelFunc

	// -- Channel Buffering --
	// defaults to 8; set to 0 for unbuffered channels
	bsize int

	// making this available to the struct. I'm doing so for Halt(), not sure if I should
	outgoing <-chan chron.Time

	// -- Waiting --
	// if true the sequence goroutine waits until the time of the next date
	// to pass it to the channel
	realtime bool
}

func Sequence(begin chron.Time) *sequence {
	return &sequence{
		begin: begin,
		bsize: 8,
	}
}

func (s *sequence) End(end chron.Time) *sequence {
	s.end = &end
	return s
}

func (s *sequence) InvertTime() *sequence {
	s.inverseTime = true
	return s
}

func (s *sequence) FixedIncrement(len chron.Length, n int) *sequence {
	if n < 1 {
		panic("sequence.FixedIncrement passed an out of range n, positive non-zero ints only.")
	}
	s.incLength = len
	s.incN = n
	return s
}

func (s *sequence) DynamicIncrement(f DynamicIncrementFunc) *sequence {
	s.incDynamic = f
	return s
}

func (s *sequence) FixedOffset(len chron.Length) *sequence {
	s.offset = &len
	return s
}

func (s *sequence) DynamicOffset(f DynamicOffsetFunc) *sequence {
	s.offsetDynamic = f
	return s
}

func (s *sequence) WithContext(ctx context.Context) *sequence {
	s.ctx = ctx
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

func (s *sequence) TimeChan() <-chan chron.Time {
	return s.start()
}

func (s *sequence) YearChan() <-chan chron.Year {
	datech := s.start()
	yearch := make(chan date.Year, s.bsize)
	go func() {
		defer close(yearch)
		for d := range datech {
			yearch <- d.ToYear()
		}
	}()
	return yearch
}

func (s *sequence) MonthChan() <-chan chron.Month {
	datech := s.start()
	ch := make(chan date.Month, s.bsize)
	go func() {
		defer close(ch)
		for d := range datech {
			ch <- d.ToMonth()
		}
	}()
	return ch
}

func (s *sequence) DayChan() <-chan chron.Day {
	datech := s.start()
	ch := make(chan date.Day, s.bsize)
	go func() {
		defer close(ch)
		for d := range datech {
			ch <- d.ToDay()
		}
	}()
	return ch
}

func (s *sequence) HourChan() <-chan chron.Hour {
	datech := s.start()
	ch := make(chan date.Hour, s.bsize)
	go func() {
		defer close(ch)
		for d := range datech {
			ch <- d.ToHour()
		}
	}()
	return ch
}

func (s *sequence) MinuteChan() <-chan chron.Minute {
	datech := s.start()
	ch := make(chan date.Minute, s.bsize)
	go func() {
		defer close(ch)
		for d := range datech {
			ch <- d.ToMinute()
		}
	}()
	return ch
}

func (s *sequence) SecondChan() <-chan chron.Second {
	datech := s.start()
	ch := make(chan date.Second, s.bsize)
	go func() {
		defer close(ch)
		for d := range datech {
			ch <- d.ToSecond()
		}
	}()
	return ch
}

// Asynchronous. It is the clients responsibility to
// read from the channel until close.
func (s *sequence) Stop() {
	go func() { s.internalCanc() }()
}

// I'm preeetty sure this is an anti-pattern. I sort of see use cases for it though.
// Like when someone wants a buffered channel that closes immediately when stop is called
func (s *sequence) Halt() {
	s.Stop()
	for d := range s.outgoing {
		goodbye(d)
	}
}

func goodbye(d chron.Time) {}

func (s *sequence) start() <-chan chron.Time {
	// check that an increment method was setup
	if s.incUnit == 0 && s.incInterval == nil && s.incDynamic == nil {
		panic("invalid sequence initialization. an increment method must be supplied")
	}

	// setup internal context and cancel
	if s.ctx == nil {
		s.internalCtx, s.internalCanc = context.WithCancel(context.Background())
		s.ctx = s.internalCtx
	} else {
		s.internalCtx, s.internalCanc = context.WithCancel(s.ctx)
	}

	// set end
	if s.end == nil {
		if s.inverseTime {
			min := date.MinValue()
			s.end = &min
		} else {
			max := date.MaxValue()
			s.end = &max
		}
	}

	// pick the right goroutine to start up based sequence fields
	if s.incDynamic != nil {
		s.outgoing = s.dynamicallyIncremented()
	} else if s.incInterval != nil {
		s.outgoing = s.intervalIncremented()
	} else {
		s.outgoing = s.unitIncremented()
	}

	if s.realtime {
		rtchan := make(chan date.Date, s.bsize)
		go func() {
			defer close(rtchan)
			for d := range s.outgoing {
				until := time.Until(d.Time())
				if until > 0 {
					time.Sleep(until)
				}
				rtchan <- d
			}
		}()
		return rtchan
	}
	return s.outgoing
}

// One of the six possible goroutines starts
func (s *sequence) unitIncremented() <-chan date.Date {
	outgoing := make(chan date.Date, s.bsize)
	if s.offsetDynamic != nil {
		if s.inverseTime {
			// back in time, unit incremented, with dynamic offset fn
			go func() {
				defer close(outgoing)
				curr := s.begin
				for curr.After(*s.end) {
					if s.ctx.Err() != nil || s.internalCtx.Err() != nil {
						return
					}
					outgoing <- curr.Increment(s.offsetDynamic(curr))
					curr = curr.Sub(s.incUnit, s.incN)
				}
			}()
			return outgoing
		} else {
			// forward in time, unit incremented, with dynamic offset fn
			go func() {
				defer close(outgoing)
				curr := s.begin
				for curr.Before(*s.end) {
					if s.ctx.Err() != nil || s.internalCtx.Err() != nil {
						return
					}
					outgoing <- curr.Increment(s.offsetDynamic(curr))
					curr = curr.Add(s.incUnit, s.incN)
				}
			}()
			return outgoing
		}
	}
	if s.offset != nil {
		if s.inverseTime {
			// back in time, unit incremented, with a fixed offset
			go func() {
				defer close(outgoing)
				curr := s.begin
				for curr.After(*s.end) {
					if s.ctx.Err() != nil || s.internalCtx.Err() != nil {
						return
					}
					outgoing <- curr.Increment(*s.offset)
					curr = curr.Sub(s.incUnit, s.incN)
				}
			}()
			return outgoing
		} else {
			// forward in time, unit incremented, with a fixed offset
			go func() {
				defer close(outgoing)
				curr := s.begin
				for curr.Before(*s.end) {
					if s.ctx.Err() != nil || s.internalCtx.Err() != nil {
						return
					}
					outgoing <- curr.Increment(*s.offset)
					curr = curr.Add(s.incUnit, s.incN)
				}
			}()
			return outgoing
		}
	}

	if s.inverseTime {
		// back in time, unit incremented, no offset
		go func() {
			defer close(outgoing)
			curr := s.begin
			for curr.After(*s.end) {
				if s.ctx.Err() != nil || s.internalCtx.Err() != nil {
					return
				}
				outgoing <- curr
				curr = curr.Sub(s.incUnit, s.incN)
			}
		}()
		return outgoing
	} else {
		// forward in time, unit incremented, no offset
		go func() {
			defer close(outgoing)
			curr := s.begin
			for curr.Before(*s.end) {
				if s.ctx.Err() != nil || s.internalCtx.Err() != nil {
					return
				}
				outgoing <- curr
				curr = curr.Add(s.incUnit, s.incN)
			}
		}()
		return outgoing
	}
}

// One of the six possible goroutines starts
func (s *sequence) intervalIncremented() <-chan date.Date {
	outgoing := make(chan date.Date, s.bsize)
	if s.offsetDynamic != nil {
		if s.inverseTime {
			// back in time, interval incremented, with dynamic offset fn
			go func() {
				defer close(outgoing)
				curr := s.begin
				for curr.After(*s.end) {
					if s.ctx.Err() != nil || s.internalCtx.Err() != nil {
						return
					}
					outgoing <- curr.Increment(s.offsetDynamic(curr))
					curr = curr.Decrement(*s.incInterval)
				}
			}()
			return outgoing
		} else {
			// forward in time, interval incremented, with dynamic offset fn
			go func() {
				defer close(outgoing)
				curr := s.begin
				for curr.Before(*s.end) {
					if s.ctx.Err() != nil || s.internalCtx.Err() != nil {
						return
					}
					outgoing <- curr.Increment(s.offsetDynamic(curr))
					curr = curr.Increment(*s.incInterval)
				}
			}()
			return outgoing
		}
	}
	if s.offset != nil {
		if s.inverseTime {
			// back in time, interval incremented, with a fixed offset
			go func() {
				defer close(outgoing)
				curr := s.begin
				for curr.After(*s.end) {
					if s.ctx.Err() != nil || s.internalCtx.Err() != nil {
						return
					}
					outgoing <- curr.Increment(*s.offset)
					curr = curr.Decrement(*s.incInterval)
				}
			}()
			return outgoing
		} else {
			// forward in time, interval incremented, with a fixed offset
			go func() {
				defer close(outgoing)
				curr := s.begin
				for curr.Before(*s.end) {
					if s.ctx.Err() != nil || s.internalCtx.Err() != nil {
						return
					}
					outgoing <- curr.Increment(*s.offset)
					curr = curr.Increment(*s.incInterval)
				}
			}()
			return outgoing
		}
	}

	if s.inverseTime {
		// back in time, interval incremented, no offset
		go func() {
			defer close(outgoing)
			curr := s.begin
			for curr.After(*s.end) {
				if s.ctx.Err() != nil || s.internalCtx.Err() != nil {
					return
				}
				outgoing <- curr
				curr = curr.Decrement(*s.incInterval)
			}
		}()
		return outgoing
	} else {
		// forward in time, interval incremented, no offset
		go func() {
			defer close(outgoing)
			curr := s.begin
			for curr.Before(*s.end) {
				if s.ctx.Err() != nil || s.internalCtx.Err() != nil {
					return
				}
				outgoing <- curr
				curr = curr.Increment(*s.incInterval)
			}
		}()
		return outgoing
	}
}

// One of the six possible goroutines starts
func (s *sequence) dynamicallyIncremented() <-chan date.Date {
	outgoing := make(chan date.Date, s.bsize)
	if s.offsetDynamic != nil {
		if s.inverseTime {
			// back in time, interval incremented, with dynamic offset fn
			go func() {
				defer close(outgoing)
				curr := s.begin
				for curr.After(*s.end) {
					if s.ctx.Err() != nil || s.internalCtx.Err() != nil {
						return
					}
					outgoing <- curr.Increment(s.offsetDynamic(curr))
					curr = curr.Decrement(s.incDynamic())
				}
			}()
			return outgoing
		} else {
			// forward in time, interval incremented, with dynamic offset fn
			go func() {
				defer close(outgoing)
				curr := s.begin
				for curr.Before(*s.end) {
					if s.ctx.Err() != nil || s.internalCtx.Err() != nil {
						return
					}
					outgoing <- curr.Increment(s.offsetDynamic(curr))
					curr = curr.Increment(s.incDynamic())
				}
			}()
			return outgoing
		}
	}
	if s.offset != nil {
		if s.inverseTime {
			// back in time, interval incremented, with a fixed offset
			go func() {
				defer close(outgoing)
				curr := s.begin
				for curr.After(*s.end) {
					if s.ctx.Err() != nil || s.internalCtx.Err() != nil {
						return
					}
					outgoing <- curr.Increment(*s.offset)
					curr = curr.Decrement(s.incDynamic())
				}
			}()
			return outgoing
		} else {
			// forward in time, interval incremented, with a fixed offset
			go func() {
				defer close(outgoing)
				curr := s.begin
				for curr.Before(*s.end) {
					if s.ctx.Err() != nil || s.internalCtx.Err() != nil {
						return
					}
					outgoing <- curr.Increment(*s.offset)
					curr = curr.Increment(s.incDynamic())
				}
			}()
			return outgoing
		}
	}

	if s.inverseTime {
		// back in time, interval incremented, no offset
		go func() {
			defer close(outgoing)
			curr := s.begin
			for curr.After(*s.end) {
				if s.ctx.Err() != nil || s.internalCtx.Err() != nil {
					return
				}
				outgoing <- curr
				curr = curr.Decrement(s.incDynamic())
			}
		}()
		return outgoing
	} else {
		// forward in time, interval incremented, no offset
		go func() {
			defer close(outgoing)
			curr := s.begin
			for curr.Before(*s.end) {
				if s.ctx.Err() != nil || s.internalCtx.Err() != nil {
					return
				}
				outgoing <- curr
				curr = curr.Increment(s.incDynamic())
			}
		}()
		return outgoing
	}
}
