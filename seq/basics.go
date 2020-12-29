package seq

//func DateRange(begin, end chron.Day) (<-chan chron.Day, cancel) {
//	return NewSequence(begin, End(end), IncrementBy(dura.Day)).DayChan()
//}
//
//func DailyAt(timeofday time.Duration) (<-chan time.Time, cancel) {
//	return NewSequence(chron.DayOf(time.Now()),
//		IncrementBy(dura.Day),
//		OffsetBy(dura.Time{Dur: timeofday}),
//		RealTime()).TimeChan()
//}
//
//func Hourly() (<-chan chron.Hour, cancel) {
//	return NewSequence(
//		chron.HourOf(time.Now()),
//		IncrementBy(dura.Hour),
//		RealTime()).HourChan()
//}
//
//func RandHourly() (<-chan time.Time, cancel) {
//	return NewSequence(chron.HourOf(time.Now()),
//		IncrementBy(dura.Hour),
//		OffsetFn(func(t chron.Time) dura.Time {
//			rnd := rand.New(rand.NewSource(t.AsTime().UnixNano())).Int()
//			offset := int(time.Millisecond) * rnd % int(time.Hour)
//			return dura.Nanos(offset)}
//		}),
//		RealTime()).TimeChan()
//}
