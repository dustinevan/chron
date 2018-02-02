package chron

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/dustinevan/chron/dura"
	"time"
)

var tchron = NewTime(2018, time.February, 27, 4, 0, 0, 0)
var tinterval = &Interval{start: tchron, end: tchron.Increment(dura.Week).AddNanos(-1), d: dura.Week}

var zeroInterval = &Interval{start: tchron, end: tchron, d: dura.Nano}

func TestNewInterval(t *testing.T) {
	assert.Exactly(t, tinterval, NewInterval(tchron, dura.Week))
	assert.Exactly(t, zeroInterval, NewInterval(tchron, dura.Nano))
}

func TestInterval_Contains(t *testing.T) {
	assert.True(t, tinterval.Contains(NewHour(2018, time.March, 2, 23)))
	assert.True(t, zeroInterval.Contains(tchron))
}

func TestInterval_Duration(t *testing.T) {
	assert.Exactly(t, dura.Week, tinterval.Duration())
}

func TestInterval_String(t *testing.T) {
	assert.Exactly(t, "start:2018-02-27 04:00:00 +0000 UTC, end:2018-03-06 03:59:59.999999999 +0000 UTC, len:Week", tinterval.String())
}