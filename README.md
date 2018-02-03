## chron
it's time :]

![](https://github.com/dustinevan/chron/blob/master/chron.png "chron")

Chron is a library that wraps time.Time and can be used as a replacement for time.Time in general purpose programming. Chron only uses underlying time.Time functionality for time calculations, so you can trust it's accuracy. Chron provides a type system that breaks up the idea of Time into three interfaces. chron.Time, dura.Time, and chron.Span. More specific implementations are added as well which make it easier to name variables and reason about what they represent. The type system is shown below:

![](https://github.com/dustinevan/chron/blob/master/typesystem.png "type system")

