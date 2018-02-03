## chron
it's time :]

![](https://github.com/dustinevan/chron/blob/master/chron.png "chron")

Chron is a library that wraps `time.Time` and can be used as a replacement for `time.Time` in general purpose programming. Chron only uses underlying `time.Time` functionality for time calculations, so you can trust it's accuracy. 

Chron provides a type system that breaks up the idea of time into three interfaces: `chron.Time`--a specific nanosecond in time, `dura.Time`--a specific length of time, and `chron.Span`--a combination of the two e.g. the year 2018. Specific implementation types that wrap `time.Time` are added to make it easier to name variables and reason about what they represent. The type system is shown below:

![](https://github.com/dustinevan/chron/blob/master/typesystem.png "type system")

