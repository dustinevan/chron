## chron
it's time :]

![](https://github.com/dustinevan/chron/blob/master/chron.png "chron")

Chron is a general purpose time library that embeds `time.Time` and can be used as a replacement. Chron uses `time.Time` for time calculations, so you can trust it's accuracy.

Why? There are many reasons, but the central one is that like `string`, `time.Time` is often used as an interface. Holidays, credit card expiration dates, hourly reporting times, postgres timestamps; all these things have different time precisions, some are used as instants, while others are used as time spans. Chron aims to wrap `time.Time` and provide a more specific type system that is consistent with simplicity of beauty of `time.Time` and `time.Duration`. 

Chron's type system breaks up the idea of time into three interfaces: `chron.Time`--a specific nanosecond in time, `dura.Time`--an exact or fuzzy length of time, and `chron.Span`--a time interval with an exact start and end, like the year 2018. The implementations of these interfaces map to different time precisions. These types make it easier to reason about the times, durations, and spans represented in your program. They also provide ways to operate on times in a way that clearly shows precision. The type system is shown below:

![](https://github.com/dustinevan/chron/blob/master/typesystem.png "type system")
