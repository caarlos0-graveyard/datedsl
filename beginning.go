package datedsl

import "time"

const (
	Day = time.Hour * 24
)

func (d DateDSL) BeginningOfDay() DateDSL {
	durations := []time.Duration{
		time.Duration(-d.Value().Hour()) * time.Hour,
		time.Duration(-d.Value().Minute()) * time.Minute,
		time.Duration(-d.Value().Second()) * time.Second,
		time.Duration(-d.Value().Nanosecond()) * time.Nanosecond,
	}
	return DateDSL{mutate(d.Value(), durations...)}
}

func (d DateDSL) BeginningOfMonth() DateDSL {
	value := d.BeginningOfDay().Value()
	return DateDSL{mutate(value, time.Duration(-d.Value().Day()+1)*Day)}
}

func mutate(value time.Time, durations ...time.Duration) time.Time {
	for _, duration := range durations {
		value = value.Add(duration)
	}
	return value
}
