package seq

import (
	"time"

	"github.com/dustinevan/chron"
	"github.com/dustinevan/chron/length"
)

func Daily(begin, end time.Time) (<-chan chron.Day, stop) {
	return Sequence(chron.DayOf(begin)).
		End(chron.DayOf(end)).
		Increment(length.Day).
		DayChan()
}
