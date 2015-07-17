package datedsl

import "time"

// BeginningOfDay returns a DSL with the current day at 00:00:00.0000
func (d DateDSL) BeginningOfDay() DateDSL {
	durations := []time.Duration{
		time.Duration(-d.Value().Hour()) * time.Hour,
		time.Duration(-d.Value().Minute()) * time.Minute,
		time.Duration(-d.Value().Second()) * time.Second,
		time.Duration(-d.Value().Nanosecond()) * time.Nanosecond,
	}
	return DateDSL{mutate(d.Value(), durations...)}
}

// BeginningOfMonth returns a DSL at first day of the current month,
// at 00:00:00.0000
func (d DateDSL) BeginningOfMonth() DateDSL {
	value := d.BeginningOfDay().Value()
	return DateDSL{mutate(value, time.Duration(-d.Value().Day()+1)*Day)}
}

// BeginningOfYear returns a DSL at 01/01 of the current year, at 00:00:00
func (d DateDSL) BeginningOfYear() DateDSL {
	value := d.Value()
	for value.Month() != time.January {
		value = value.Add(-time.Duration(30) * Day)
	}
	dsl := DateDSL{value}
	return dsl.BeginningOfMonth()
}
