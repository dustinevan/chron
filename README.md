## chron
it's time :]

![](https://github.com/dustinevan/chron/blob/master/chron.png "chron")

Chron is a general purpose time library that embeds `time.Time` and can be used as a replacement. Chron uses `time.Time` for time calculations, so you can trust it's accuracy.

Why? There are many reasons, but the central one is that `time.Time` is often used as an interface. Holidays, credit card expiration dates, hourly reporting times, postgres timestamps; all these things have different time precisions, some are used as instants, while others are used as time spans. 

Chron aims to wrap `time.Time` and provide a more specific type system that is consistent with simplicity of beauty of `time.Time` and `time.Duration`. Chron's type system breaks up the idea of time into three interfaces: 
```golang
chron.Time // a specific nanosecond in time
dura.Time  // an exact or fuzzy length of time
chron.Span // a time interval with an exact start and end, like the year 2018
```
The implementations of these interfaces map to different time precisions. These types make it easier to reason about the times, durations, and spans represented in your program. They also provide ways to operate on times in a way that clearly shows precision. The type system is shown below:

![](https://github.com/dustinevan/chron/blob/master/typesystem.png "type system")

#### chron.Time implementations

`chron.Year` ... `chron.Chron` are structs that embed `time.Time` and are truncated to a certain precision. You know that `chron.Hour` will always have 0 min, sec and nanoseconds. `chron.Chron` is the replacement for `time.Time` with nanosecond precision. These structs implement `chron.Time` which requires conversion functions to all other types.
```golang
now := chron.Now()          // type chron.Chron 2018-02-04 04:25:20.056473271 +0000 UTC
this_micro := now.AsMicro() // type chron.Micro 2018-02-04 04:25:20.056473 +0000 UTC
this_milli := now.AsMilli() // type chron.Milli 2018-02-04 04:25:20.056 +0000 UTC
...
this_month := now.AsMonth() // type chron.Month 2018-02-01 00:00:00 +0000 UTC
this_year := now.AsYear()   // type chron.Year 2018-01-01 00:00:00 +0000 UTC
time_time := now.AsTime()   // type time.Time 2018-02-04 04:25:20.056473271 +0000 UTC
```
Increment and Decrement functions are also required as part of the chron.Time interface. These take in a dura.Time and return a chron.Chron
```golang
next_hour := chron.Now().Increment(dura.Hour).AsHour()
previous_second := chron.Now().Decrement(dura.Second).AsSecond()
```
Increment and Decrement are meant to handle any possible fuzzy or exact duration. So these operations are better done useing the many convenience methods. 
```golang
next_hour := chron.ThisHour().AddN(1)
previous_second := chron.ThisSecond().AddN(-1)
```
JSON Unmarshaling methods support 25 different formats. Scan and Value methods are implemented to allow DB support. 


