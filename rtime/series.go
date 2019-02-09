package rtime

import "github.com/dustinevan/chron"

type Series struct {
	timeFunc RTime
	start chron.Time
	stop TimeFilter
}

func NewSeries(timeFunc RTime, start chron.Time, stop TimeFilter) Series {
	return Series{timeFunc:timeFunc, start: start, stop: stop}
}

func (s *Series) NextN(n int) []chron.Chron {
	result := make([]chron.Chron, n)
	for i := 0; i < n; i++ {
		s.start = s.timeFunc(s.start)
		if s.stop(s.start) {
			return result
		}
		result[i] = s.start.AsChron()
	}
	return result
}

