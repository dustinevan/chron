package rtime

import (
	"github.com/dustinevan/chron"
	"testing"
)

func BenchmarkIsUSBusinessHoliday(b *testing.B) {
	today := chron.Today()
	for i := 0; i < b.N; i++ {
		IsUSBusinessHoliday(today.AddN(i).AsTime())
	}
}
