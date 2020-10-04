package main

import (
	"fmt"

	"time"

	"github.com/dustinevan/chron"
	"github.com/dustinevan/chron/seq"
)

func main() {
	DateRangeExample()
}

func DateRangeExample() {
	daterange, stop := seq.DateRange(chron.NewDay(2017, time.December, 31), chron.NewDay(2018, time.February, 1))
	i := 0
	for t := range daterange {
		if i == 2 {
			stop()
		}
		fmt.Println(t)
		i++
	}
}

func DailyAt() {
	daterange, stop := seq.DailyAt(chron.NewDay())
	i := 0
	for t := range daterange {
		if i == 2 {
			stop()
		}
		fmt.Println(t)
		i++
	}
}