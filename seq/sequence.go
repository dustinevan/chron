package seq

//type SeqOption func(*Sequence) error
//
//type Sequence struct {
//
//	// -- Boundary and Flow --
//
//	// Inclusive--this is the first time in the sequence. If begin does not
//	// match the precision of the output time, the first output will be the
//	// next occurrence of that time precision. e.g. a positive time flow
//	// MinuteChan() seq, with begin := 2017-10-16 16:40:04.049121656 will
//	// begin at 2017-10-16 16:41:00.0, neg time flow would begin at 16:40
//	begin chron.TimeExact
//
//	// Exclusive--the timeExact of end is not included in the sequence.
//	// If not set this defaults to date.MaxValue() or date.MinValue()
//	// depending on inverseTime. A
//	end chron.TimeExact
//
//	// true = the sequence goes back in time
//	negativeTime bool
//
//	// -- Increment Fields -- Note: Panics if one of these is not set.
//
//	// constantIncrement is the length added to get the next seq time.
//	increment dura.Time
//	repeats   int
//
//	// variableIncrementFn takes precedence over constantIncrement; The
//	// sequence is incremented by the returned Length
//	incrementFn func(exact chron.TimeExact) chron.TimeExact
//
//	// -- Offset Fields --
//
//	// if set, this length is added to the next seq time. Offset is not
//	// used to calculate the next seq time. e.g. seq bases the next
//	// date off curr not curr + offset.
//	offset dura.Time
//	// variableOffset takes precedence over constantOffset. A good example
//	// use case is random offsets the spread the sequence times across an
//	// hour.
//	offsetFn func(chron.Time) dura.Time
//
//	// -- Channel Buffering --
//	// Defaults to 1; set to 0 for unbuffered channels
//	bsize int
//
//	// -- Synchronization --
//	// if true the sequence goroutine waits until the time of the next date
//	// to pass it to the channel
//	realtime bool
//}
//
//func NewSequence(begin chron.Time, opts ...SeqOption) *Sequence {
//	s := &Sequence{
//		begin: begin.AsTimeExact(),
//		bsize: 1,
//		end:   chron.MaxValue(),
//	}
//	for _, opt := range opts {
//		err := opt(s)
//		if err != nil {
//			panic(err.Error())
//		}
//	}
//	return s
//}
//
//func End(end chron.Time) SeqOption {
//	return func(s *Sequence) error {
//		if s.negativeTime == false {
//			if s.begin.After(end.AsTime()) {
//				return fmt.Errorf("end %s and begin %s are inconsistent with negativeTime = %s",
//					s.end, s.begin, s.negativeTime)
//			}
//			s.end = end.AsTimeExact()
//			return nil
//
//		}
//		if s.begin.Before(end.AsTime()) {
//			return fmt.Errorf("end %s and begin %s are inconsistent with negativeTime = %s",
//				s.end, s.begin, s.negativeTime)
//		}
//		s.end = end.AsTimeExact()
//		return nil
//	}
//}
//
//func InclusiveEnd(end chron.Time) SeqOption {
//	return End(end.Increment(dura.Nano))
//}
//
//func ForPeriod(len dura.Time) SeqOption {
//	return func(s *Sequence) error {
//		if s.negativeTime == false {
//			s.end = s.begin.Increment(len)
//			return nil
//		}
//		s.end = s.begin.Decrement(len)
//		return nil
//	}
//}
//
//func IncrementBy(len dura.Time) SeqOption {
//	return func(s *Sequence) error {
//		if s.incrementFn != nil {
//			return fmt.Errorf("IncrementBy called when a dynamic increment function is already set")
//		}
//		s.increment = len
//		return nil
//	}
//}
//
//func IncrementFn(f func(chron.TimeExact) chron.TimeExact) SeqOption {
//	return func(s *Sequence) error {
//		s.incrementFn = f
//		return nil
//	}
//}
//
//func OffsetBy(len dura.Time) SeqOption {
//	return func(s *Sequence) error {
//		if s.offsetFn != nil {
//			return fmt.Errorf("OffsetBy called when a dynamic offset function is already set")
//		}
//		s.increment = len
//		return nil
//	}
//}
//
//func OffsetFn(f func(chron.Time) dura.Time) SeqOption {
//	return func(s *Sequence) error {
//		s.offsetFn = f
//		return nil
//	}
//}
//
//func RealTime() SeqOption {
//	return func(s *Sequence) error {
//		s.realtime = true
//		return nil
//	}
//}
//
//func ChanSize(i int) SeqOption {
//	return func(s *Sequence) error {
//		s.bsize = i
//		return nil
//	}
//}
//
//type cancel func()
//
//func (s *Sequence) TimeChan() (<-chan time.Time, cancel) {
//	in, canc := s.start()
//	out := make(chan time.Time, s.bsize)
//	stop := func() {
//		canc()
//		for t := range out {
//			seeya(t)
//		}
//	}
//	go func() {
//		defer close(out)
//		for t := range in {
//			out <- t.AsTime()
//		}
//	}()
//	return out, stop
//}
//
//func (s *Sequence) ExactTimeChan() (<-chan chron.TimeExact, cancel) {
//	in, canc := s.start()
//	out := make(chan chron.TimeExact, s.bsize)
//	stop := func() {
//		canc()
//		for t := range out {
//			goodbye(t)
//		}
//	}
//	go func() {
//		defer close(out)
//		for d := range in {
//			out <- d
//		}
//	}()
//	return out, stop
//}
//
//func (s *Sequence) YearChan() (<-chan chron.Year, cancel) {
//	in, canc := s.start()
//	out := make(chan chron.Year, s.bsize)
//	stop := func() {
//		canc()
//		for t := range out {
//			goodbye(t)
//		}
//	}
//	go func() {
//		defer close(out)
//		for d := range in {
//			out <- d.AsYear()
//		}
//	}()
//	return out, stop
//}
//
//func (s *Sequence) MonthChan() (<-chan chron.Month, cancel) {
//	in, canc := s.start()
//	out := make(chan chron.Month, s.bsize)
//	stop := func() {
//		canc()
//		for t := range out {
//			goodbye(t)
//		}
//	}
//	go func() {
//		defer close(out)
//		for d := range in {
//			out <- d.AsMonth()
//		}
//	}()
//	return out, stop
//}
//
//func (s *Sequence) DayChan() (<-chan chron.Day, cancel) {
//	in, canc := s.start()
//	out := make(chan chron.Day, s.bsize)
//	stop := func() {
//		canc()
//		for t := range out {
//			goodbye(t)
//		}
//	}
//	go func() {
//		defer close(out)
//		for d := range in {
//			out <- d.AsDay()
//		}
//	}()
//	return out, stop
//}
//
//func (s *Sequence) HourChan() (<-chan chron.Hour, cancel) {
//	in, canc := s.start()
//	out := make(chan chron.Hour, s.bsize)
//	stop := func() {
//		canc()
//		for t := range out {
//			goodbye(t)
//		}
//	}
//	go func() {
//		defer close(out)
//		for d := range in {
//			out <- d.AsHour()
//		}
//	}()
//	return out, stop
//}
//
//func (s *Sequence) MinuteChan() (<-chan chron.Minute, cancel) {
//	in, canc := s.start()
//	out := make(chan chron.Minute, s.bsize)
//	stop := func() {
//		canc()
//		for t := range out {
//			goodbye(t)
//		}
//	}
//	go func() {
//		defer close(out)
//		for d := range in {
//			out <- d.AsMinute()
//		}
//	}()
//	return out, stop
//}
//
//func (s *Sequence) SecondChan() (<-chan chron.Second, cancel) {
//	in, canc := s.start()
//	out := make(chan chron.Second, s.bsize)
//	stop := func() {
//		canc()
//		for t := range out {
//			goodbye(t)
//		}
//	}
//	go func() {
//		defer close(out)
//		for d := range in {
//			out <- d.AsSecond()
//		}
//	}()
//	return out, stop
//}
//
//func goodbye(t chron.Time) {}
//
//func seeya(t time.Time) {}
//
//func (s *Sequence) start() (<-chan chron.TimeExact, context.CancelFunc) {
//
//	// setup internal context and cancel
//	ctx, canc := context.WithCancel(context.Background())
//
//	var out <-chan chron.TimeExact
//	// pick the right goroutine to start up based on Sequence fields
//	if s.incrementFn != nil {
//		out = s.variableIncs(ctx)
//	} else if s.increment != nil {
//		out = s.fixedIncs(ctx)
//	} else {
//		panic("invalid sequence initialization. an increment method must be supplied")
//	}
//
//	if s.realtime {
//		rtchan := make(chan chron.TimeExact, s.bsize)
//		go func() {
//			defer close(rtchan)
//			for t := range out {
//				until := time.Until(t.Time)
//				if until > 0 {
//					time.Sleep(until)
//				}
//				rtchan <- t
//			}
//		}()
//		return rtchan, canc
//	}
//	return out, canc
//}
//
//// One of the six possible goroutines starts
//func (s *Sequence) fixedIncs(ctx context.Context) <-chan chron.TimeExact {
//	out := make(chan chron.TimeExact, s.bsize)
//	if s.offsetFn != nil {
//		if s.begin.After(s.end.Time) {
//			// back in time, interval incremented, with dynamic offset fn
//			go func() {
//				defer close(out)
//				curr := s.begin
//				for curr.After(s.end.Time) {
//					if ctx.Err() != nil {
//						return
//					}
//					out <- curr.Increment(s.offsetFn(curr))
//					curr = curr.Decrement(s.increment)
//				}
//			}()
//			return out
//		} else {
//			// forward in time, interval incremented, with dynamic offset fn
//			go func() {
//				defer close(out)
//				curr := s.begin
//				for curr.Before(s.end.Time) {
//					if ctx != nil {
//						return
//					}
//					out <- curr.Increment(s.offsetFn(curr))
//					curr = curr.Increment(s.increment)
//				}
//			}()
//			return out
//		}
//	}
//	if s.offset != nil {
//		if s.begin.After(s.end.Time) {
//			// back in time, interval incremented, with a fixed offset
//			go func() {
//				defer close(out)
//				curr := s.begin
//				for curr.After(s.end.Time) {
//					if ctx.Err() != nil {
//						return
//					}
//					out <- curr.Increment(s.offset)
//					curr = curr.Decrement(s.increment)
//				}
//			}()
//			return out
//		} else {
//			// forward in time, interval incremented, with a fixed offset
//			go func() {
//				defer close(out)
//				curr := s.begin
//				for curr.Before(s.end.Time) {
//					if ctx.Err() != nil {
//						return
//					}
//					out <- curr.Increment(s.offset)
//					curr = curr.Increment(s.increment)
//				}
//			}()
//			return out
//		}
//	}
//
//	if s.begin.After(s.end.Time) {
//		// back in time, interval incremented, no offset
//		go func() {
//			defer close(out)
//			curr := s.begin
//			for curr.After(s.end.Time) {
//				if ctx.Err() != nil {
//					return
//				}
//				out <- curr
//				curr = curr.Decrement(s.increment)
//			}
//		}()
//		return out
//	} else {
//		// forward in time, interval incremented, no offset
//		go func() {
//			defer close(out)
//			curr := s.begin
//			for curr.Before(s.end.Time) {
//				if ctx.Err() != nil {
//					return
//				}
//				out <- curr
//				curr = curr.Increment(s.increment)
//			}
//		}()
//		return out
//	}
//}
//
//// One of the six possible goroutines starts
//func (s *Sequence) variableIncs(ctx context.Context) <-chan chron.TimeExact {
//	out := make(chan chron.TimeExact, s.bsize)
//	if s.offsetFn != nil {
//		if s.begin.After(s.end.Time) {
//			// back in time, interval incremented, with dynamic offset fn
//			go func() {
//				defer close(out)
//				curr := s.begin
//				for curr.After(s.end.Time) {
//					if ctx.Err() != nil {
//						return
//					}
//					out <- curr.Increment(s.offsetFn(curr))
//					curr = s.incrementFn(curr)
//				}
//			}()
//			return out
//		} else {
//			// forward in time, interval incremented, with dynamic offset fn
//			go func() {
//				defer close(out)
//				curr := s.begin
//				for curr.Before(s.end.Time) {
//					if ctx != nil {
//						return
//					}
//					out <- curr.Increment(s.offsetFn(curr))
//					curr = s.incrementFn(curr)
//				}
//			}()
//			return out
//		}
//	}
//	if s.offset != nil {
//		if s.begin.After(s.end.Time) {
//			// back in time, interval incremented, with a fixed offset
//			go func() {
//				defer close(out)
//				curr := s.begin
//				for curr.After(s.end.Time) {
//					if ctx.Err() != nil {
//						return
//					}
//					out <- curr.Increment(s.offset)
//					curr = s.incrementFn(curr)
//				}
//			}()
//			return out
//		} else {
//			// forward in time, interval incremented, with a fixed offset
//			go func() {
//				defer close(out)
//				curr := s.begin
//				for curr.Before(s.end.Time) {
//					if ctx.Err() != nil {
//						return
//					}
//					out <- curr.Increment(s.offset)
//					curr = s.incrementFn(curr)
//				}
//			}()
//			return out
//		}
//	}
//
//	if s.begin.After(s.end.Time) {
//		// back in time, interval incremented, no offset
//		go func() {
//			defer close(out)
//			curr := s.begin
//			for curr.After(s.end.Time) {
//				if ctx.Err() != nil {
//					return
//				}
//				out <- curr
//				curr = s.incrementFn(curr)
//			}
//		}()
//		return out
//	} else {
//		// forward in time, interval incremented, no offset
//		go func() {
//			defer close(out)
//			curr := s.begin
//			for curr.Before(s.end.Time) {
//				if ctx.Err() != nil {
//					return
//				}
//				out <- curr
//				curr = s.incrementFn(curr)
//			}
//		}()
//		return out
//	}
//}
