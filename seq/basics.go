package seq

import (
	"time"

	"math/rand"

	"github.com/dustinevan/chron"
	"github.com/dustinevan/chron/length"
)

func DateRange(begin, end chron.Day) (<-chan chron.Day, stop) {
	return NewSequence(begin, End(end), IncrementBy(length.Hour)).DayChan()
}

func DailyAt(timeofday time.Duration) (<-chan time.Time, stop) {
	return NewSequence(chron.DayOf(time.Now()),
		IncrementBy(length.Day),
		OffsetBy(length.Duration{Dur: timeofday}),
		RealTime()).TimeChan()
}

func Hourly() (<-chan chron.Hour, stop) {
	return NewSequence(
		chron.HourOf(time.Now()),
		IncrementBy(length.Hour),
		RealTime()).HourChan()
}

func RandHourly() (<-chan time.Time, stop) {
	return NewSequence(chron.HourOf(time.Now()),
		IncrementBy(length.Hour),
		OffsetFn(func(t chron.Time) chron.Length {
			rnd := rand.New(rand.NewSource(t.AsTime().UnixNano())).Int()
			offset := int(time.Nanosecond) * rnd % int(time.Hour)
			return length.Duration{Dur: time.Duration(offset)}
		}),
		RealTime()).TimeChan()
}
