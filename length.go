package chron

import (
	"time"
)

type Length interface {
	Years() int
	Months() int
	Days() int
	Duration() time.Duration
}
