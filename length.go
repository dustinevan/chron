package chron

import (
	"time"
	"github.com/dustinevan/chron/length"
)

type Length interface {
	Mult(int) length.Duration
	Years() int
	Months() int
	Days() int
	Duration() time.Duration
}
