package main

import (
	"fmt"

	"github.com/dustinevan/chron"
	"github.com/dustinevan/chron/length"
	"github.com/dustinevan/chron/seq"
)

func main() {
	mins, _ := seq.Sequence(chron.ThisHour().AddN(1)).
		Length(length.Hours(12)).
		Increment(length.Hours(2)).
		Offset(length.Sum(length.Mins(30), length.Secs(7))).
		MonthChan()

	for m := range mins {
		fmt.Println(m)
	}

}
