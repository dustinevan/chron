package date

import (
	"time"

	"github.com/dustinevan/time/chron"
)

type Holiday int

const (
	// US Federal Holidays
	NewYearsDay Holiday = iota
	MartinLutherKing
	PresidentsDay
	MemorialDay
	IndependenceDay
	LaborDay
	ColumbusDay
	VeteransDay
	ThanksgivingDay
	ChristmasDay

	// UK Holidays

)

type YearDayFunc func(chron.Year) chron.Day

var actualHolidays = []YearDayFunc{
	// NewYearsDay
	func(y chron.Year) chron.Day {
		return y.AsDay()
	},
	//MartinLutherKing
	func(y chron.Year) chron.Day {
		thirdMonday := NthWeekDay{weekday: time.Monday, n: 3}
		jan := chron.NewMonth(y.Year(), time.January)
		d, _ := thirdMonday.OfMonth(jan)
		return d
	},
	//PresidentsDay
	func(y chron.Year) chron.Day {
		thirdMonday := NthWeekDay{weekday: time.Monday, n: 3}
		feb := chron.NewMonth(y.Year(), time.February)
		d, _ := thirdMonday.OfMonth(feb)
		return d
	},
	//MemorialDay
	func(y chron.Year) chron.Day {
		return LastWeekdayOfMonth(chron.NewMonth(y.Year(), time.May), time.Monday)
	},
	//IndependenceDay
	func(y chron.Year) chron.Day {
		return chron.NewDay(y.Year(), 7, 4)
	},
	//LaborDay
	func(y chron.Year) chron.Day {
		firstMonday := NthWeekDay{weekday: time.Monday, n: 1}
		sep := chron.NewMonth(y.Year(), time.September)
		d, _ := firstMonday.OfMonth(sep)
		return d
	},
	//ColumbusDay
	func(y chron.Year) chron.Day {
		secondMonday := NthWeekDay{weekday: time.Monday, n: 2}
		oct := chron.NewMonth(y.Year(), time.October)
		d, _ := secondMonday.OfMonth(oct)
		return d
	},
	//VeteransDay
	func(y chron.Year) chron.Day {
		return chron.NewDay(y.Year(), time.November, 11)
	},
	//ThanksgivingDay
	func(y chron.Year) chron.Day {
		fourthThursday := NthWeekDay{weekday: time.Thursday, n: 4}
		nov := chron.NewMonth(y.Year(), time.November)
		d, _ := fourthThursday.OfMonth(nov)
		return d
	},
	//ChristmasDay
	func(y chron.Year) chron.Day {
		return chron.NewDay(y.Year(), time.December, 25)
	},
}

var holidayObservances = []YearDayFunc{
	func(y chron.Year) chron.Day {
		return ClosestNonWeekend(NewYearsDay.Date(y))
	},
	MartinLutherKing.Date,
	PresidentsDay.Date,
	MemorialDay.Date,
	func(y chron.Year) chron.Day {
		return ClosestNonWeekend(IndependenceDay.Date(y))
	},
	LaborDay.Date,
	ColumbusDay.Date,
	func(y chron.Year) chron.Day {
		return ClosestNonWeekend(VeteransDay.Date(y))
	},
	ThanksgivingDay.Date,
	func(y chron.Year) chron.Day {
		return ClosestNonWeekend(ChristmasDay.Date(y))
	},
}

func (h Holiday) Date(year chron.Year) chron.Day {
	return actualHolidays[int(h)](year)
}

// When a holiday falls on a weekend, the business day observance is moved to a weekday
func (h Holiday) Observance(year chron.Year) chron.Day {
	return holidayObservances[int(h)](year)
}
