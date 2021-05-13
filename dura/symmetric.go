package dura

// import "time"
//
// // A Symmetric Duration is a dura.Time for which, like time.Duration, the commutative property
// // is preserved when pared with time.Time.AddDate. It however is not limited to 290 years
// type Symmetric struct {
// 	years int
// 	duration time.Duration
// }
//
// func SymmetricDuration(t1, t2 time.Time) Symmetric {
// 	if t1.Equal(t2) {
// 		return Symmetric{}
// 	}
// 	if t1.After(t2) {
// 		s := SymmetricDuration(t2, t1)
// 		return Symmetric{s.years*-1, s.duration*-1}
// 	}
//
// 	years := t1.Year() - t2.Year()
// 	centuries := years/100
// 	tmp := t1.AddDate(100 * centuries, 0 ,0)
// 	dur := t2.Sub(tmp)
// 	return Symmetric{100*centuries, dur}
// }
