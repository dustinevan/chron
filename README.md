## chron
it's time :]

![](https://github.com/dustinevan/chron/blob/master/chron.png "chron")

Chron is a general purpose time library that embeds `time.Time` and can be used as a replacement. Chron uses `time.Time` for calculations, so you can trust it's accuracy.

Why? There are many reasons, but the central one is that `time.Time` is often used as an interface. Holidays, credit card expiration dates, hourly reporting times, postgres timestamps; all these things have different time precisions, some are used as instants, while others are used as timespans. 

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
`Increment` and `Decrement` functions are also required as part of the `chron.Time` interface. These functions handle any possible fuzzy or exact duration `(dura.Time)` and return a new `chron.Chron`
```golang
h := chron.Now().Increment(dura.NewDuration(1, 5, 32, time.Hour * 4 + time.Minute * 15 + time.Second * 30)).AsHour()
// the hour 1 year, 5 months, 32 days, 4 hour, 15 minutes, and 30 seconds from now
```
While `Increment` and `Decrement` handle any time duration, simple operations are better done using the many convenience methods. 
```golang
now := chron.Now() // type chron.Chron 
next_hour := chron.ThisHour().AddN(1) // type chron.Hour
five_minutes_ago := now.AddMinutes(-5) // type chron.Chron
previous_second := chron.ThisSecond().AddN(-1) // type chron.Second
```
JSON Unmarshaling methods support 25 different formats--more can be added by appending to `chron.ParseFormats`. `Scan` and `Value` methods are also implemented to allow DB support. 

Becuase `time.Time` is embedded, time package methods can be accessed directly. `Before`, `After`, and `UnmarshalJSON` are overwritten, but will provide the same functionality. `Before` and `After` now handle the overlapping nature of timespans, and `UnmarshalJSON` adds more formats besides `time.RFC3339`.     

#### Time Zones
I have been burned by timezoned time data. I am of the opinion that all times belong in UTC until a human being wants to see them. I could be naive or wrong about this. Currently chron converts all times to UTC, so using the constructors will guarantee UTC internal times. If a chron user wants to create intances via `chron.Chron{...}`, it is their responsibility to ensure the underlying time is in UTC. 

#### dura.Time implmentations
`dura.Time` is an interface that represents the same thing as a `time.Duration` with one key difference. `dura.Time` implementations can handle fuzzy durations. The concepts Month and Year don't have fixed lengths until matched with a specific instant in time. 

`time.Time` uses `Add` and `AddDate` to deal with these differences, but there isn't currently a way to hold an pass around a duration that has months and years in it. If time libraries start handling leap seconds in a similar way, week, day, hour, minute and second will also become fuzzy. 

There are currently 2 implementations of `dura.Time`: `dura.Duration` as struct that holds years, months, days, and a time.Duration, and `dura.Unit` an int enum that defines standard units of time. 
```golang
d := dura.NewDuration(1, 3, 15, time.Hour*12)
d = d.Mult(3) // 3 years, 9 months, 45 days, 36 hours
now := chron.Now() // 2018-02-04 21:09:50.096961028 +0000 UTC
future := now.Increment(d) //2021-12-21 09:09:50.096961028 +0000 UTC
```
Convenience methods have overcome the original uses for `dura.Unit` constants, there are left here for possible use in switch statements and as the hard coded durations in chron.Span implementations. 
```golang
today := chron.today()
// before 
noon := today.Increment(dura.Hour.Mult(12)).AsHour()
// new
noon := today.AddHours(12)
```
#### chron.Span implementations
`chron.Span` is just a time combined with a duration. Each of the `chron.Time` implementations also implement `chron.Span`. This interface allows the compairson of timespans rather than just time instants. 
```golang
tomorrow := chron.Today().AddDays(1)
noon_tomorrow := tomorrow.AddHours(12)
if tomorrow.Contains(noon_tomorrow) {
    fmt.Println(:])
}
if chron.Today().Before(tomorrow) {
    fmt.Println(:])
}
```

#### Future Plans
I actually set out to write a scheduler, then I decided I needed a library that could output a stream of times base on input arguments. Then I decided to write some time conveniece stuff to make all the odd time precision and fuzzy duration issues easier to deal with. Chron will eventually become the second thing. I plan to add time series, time sequence and relative time functionality in the near future. What exists now though is solid, all changes to current code will preserve backward compatibility. 

#### Issues
Please make issues if you have things you want to discuss or that you think need fixing. I'm all ears.  
