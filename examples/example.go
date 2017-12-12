package main

import (
	"fmt"

	"github.com/dustinevan/chron"
	"github.com/dustinevan/chron/length"
	"github.com/dustinevan/chron/rtime"
	"github.com/dustinevan/chron/seq"
)

func main() {
	mins, _ := seq.Sequence(chron.ThisHour().AddN(1)).
		Length(length.Hours(12)).
		Increment(length.Hours(1)).
		IncrementFn()
	Offset(length.Mins(15)).
		TimeChan()

	for m := range mins {
		fmt.Println(m)
	}

	s := rtime.Second{37}
	fmt.Println(s.Next(chron.Now()))
}
