package main

import (
	"fmt"

	"github.com/dustinevan/chron/seq"
)

func main() {
	random_hourly, stop := seq.RandHourly()
	i := 0
	for t := range random_hourly {
		if i == 5 {
			stop()
		}
		fmt.Println(t)
		i++
	}
}
