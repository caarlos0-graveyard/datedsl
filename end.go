package datedsl

import "time"

// EndOfSecond returns a DSL with the current minute at 0000ns
func (d DateDSL) EndOfSecond() DateDSL {
	return DateDSL{mutate(
		d.BeginningOfSecond().Value(),
		time.Second-time.Nanosecond,
	)}
}

// EndOfMinute returns a DSL with the current minute at 00s.0000ns
func (d DateDSL) EndOfMinute() DateDSL {
	return DateDSL{mutate(
		d.BeginningOfMinute().Value(),
		time.Minute-time.Nanosecond,
	)}
}

// EndOfHour returns a DSL with the current hour at 00m:00s.0000ns
func (d DateDSL) EndOfHour() DateDSL {
	return DateDSL{mutate(
		d.BeginningOfHour().Value(),
		time.Hour-time.Nanosecond,
	)}
}

// EndOfDay returns a DSL with the current day at 00:00:00.0000
func (d DateDSL) EndOfDay() DateDSL {
	return DateDSL{mutate(
		d.BeginningOfDay().Value(),
		Day-time.Nanosecond,
	)}
}

// EndOfMonth returns a DSL at first day of the current month,
// at 00:00:00.0000
func (d DateDSL) EndOfMonth() DateDSL {
	return DateDSL{mutate(
		d.EndOfDay().Value(),
		Month,
		time.Duration(-d.Value().Day()+1)*Day,
	)}
}

// EndOfYear returns a DSL at 01/01 of the current year, at 00:00:00
func (d DateDSL) EndOfYear() DateDSL {
	dsl := DateDSL{mutate(
		d.Value(),
		time.Duration(d.Value().Month()-2)*Month,
	)}
	return dsl.EndOfMonth()
}
