package chron

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseUnixSeconds(t *testing.T) {
	d := NewDay(2016, 03, 17)
	ti, err := ParseUnixSeconds("1458172800")
	assert.True(t, err == nil)
	assert.Exactly(t, d.Time, ti)

}

