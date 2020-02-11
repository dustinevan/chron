package rtime

import (
	"fmt"
	"github.com/dustinevan/chron"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMonthSeriesProducesASeriesOfMonths(t *testing.T) {
	//second := func(year chron.Year) chron.Month {
	//	return chron.NewMonth(year.Year(), time.July)
	//}

	months := NewSeries(everyJuly, thisYear, func(t chron.Time) bool {
		return t.AsMonth().Year() >= 2100
	}, ctx)
	for i := 0; i < 2100-thisYear.Year(); i++ {
		actual := <-months
		expected := chron.NewMonth(thisYear.Year()+i, time.July)
		assert.Equal(t, expected, actual)
	}
	for m := range months {
		assert.Fail(t, fmt.Sprintf("months channel produced a month after cancel was called. %s recieved", m.String()))
	}
}

func TestSeries_NextN(t *testing.T) {

}
