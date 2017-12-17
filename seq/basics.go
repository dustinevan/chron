package seq

import (
	"time"

	"math/rand"

	"github.com/dustinevan/chron"
	"github.com/dustinevan/chron/dura"
	"github.com/golang/protobuf/ptypes/duration"
)

func DateRange(begin, end chron.Day) (<-chan chron.Day, stop) {
	return NewSequence(begin, End(end), IncrementBy(dura.Day)).DayChan()
}

func DailyAt(timeofday time.Duration) (<-chan time.Time, stop) {
	return NewSequence(chron.DayOf(time.Now()),
		IncrementBy(dura.Day),
		OffsetBy(dura.Duration{Dur: timeofday}),
		RealTime()).TimeChan()
}

func Hourly() (<-chan chron.Hour, stop) {
	return NewSequence(
		chron.HourOf(time.Now()),
		IncrementBy(dura.Hour),
		RealTime()).HourChan()
}

func RandHourly() (<-chan time.Time, stop) {
	return NewSequence(chron.HourOf(time.Now()),
		IncrementBy(dura.Hour),
		OffsetFn(func(t chron.Time) dura.Duration {
			rnd := rand.New(rand.NewSource(t.AsTime().UnixNano())).Int()
			offset := int(time.Millisecond) * rnd % int(time.Hour)
			return dura.Nanos(offset)}
		}),
		RealTime()).TimeChan()
}
