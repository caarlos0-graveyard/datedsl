package datedsl

import "time"

type DateDSL struct {
	value time.Time
}

func New(t time.Time) DateDSL {
	return DateDSL{t}
}

func (d DateDSL) Value() time.Time {
	return d.value
}

func (d DateDSL) String() string {
	return d.value.String()
}
