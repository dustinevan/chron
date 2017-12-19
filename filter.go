package chron

import (
	"sync"
	"github.com/dustinevan/chron/dura"
)

// Filter
type FilterFunc func(Time) (filtered bool, span Span)


type Filter struct {
	fn FilterFunc
	priority int
	invert bool
}

func NewFilter(fn FilterFunc, priority int, invert bool) *Filter {
	return &Filter{
		fn: fn,
		priority: priority,
		invert: invert,
	}
}

func (f Filter) Check(t Time) (filtered bool, span Span) {
	filtered, span = f.fn(t)
	if f.invert {
		return !filtered, span
	}
	return filtered, span
}

func (f Filter) Priority() int {
	return f.priority
}

type Schedule struct {
	filters []Filter
	precision dura.Unit
}

