package datedsl

import "time"

const (
	// Day represents a day in Nanoseconds
	Day = time.Hour * 24
)

// DateDSL the main type
type DateDSL struct {
	value time.Time
}

// New creates a DateDSL object
func New(t time.Time) DateDSL {
	return DateDSL{t}
}

// Value returns the inner value of the DateDSL
func (d DateDSL) Value() time.Time {
	return d.value
}

// String returns the inner value as a String
func (d DateDSL) String() string {
	return d.value.String()
}

func mutate(value time.Time, durations ...time.Duration) time.Time {
	for _, duration := range durations {
		value = value.Add(duration)
	}
	return value
}
