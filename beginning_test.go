package datedsl

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBeginningOfSecond(t *testing.T) {
	value := New(parse("Fri Jul 17 10:13:59 MST 2015")).BeginningOfSecond()
	expected := parse("Fri Jul 17 10:13:59 MST 2015")
	assert.Equal(t, expected.String(), value.String())
	assert.Equal(t, 0, value.Value().Nanosecond())
}

func TestBeginningOfMinute(t *testing.T) {
	value := New(parse("Fri Jul 17 10:13:59 MST 2015")).BeginningOfMinute()
	expected := parse("Fri Jul 17 10:13:00 MST 2015")
	assert.Equal(t, expected.String(), value.String())
}

func TestBeginningOfHour(t *testing.T) {
	value := New(parse("Fri Jul 17 10:13:59 MST 2015")).BeginningOfHour()
	expected := parse("Fri Jul 17 10:00:00 MST 2015")
	assert.Equal(t, expected.String(), value.String())
}

func TestBeginningOfDay(t *testing.T) {
	value := New(parse("Fri Jul 17 10:13:59 MST 2015")).BeginningOfDay()
	expected := parse("Fri Jul 17 00:00:00 MST 2015")
	assert.Equal(t, expected.String(), value.String())
}

func TestFirstDayOfMonth(t *testing.T) {
	value := New(parse("Fri Jul 17 09:54:37 MST 2015")).BeginningOfMonth()
	expected := parse("Wed Jul 1 00:00:00 MST 2015")
	assert.Equal(t, expected.String(), value.String())
}

func TestFirstDayOfYear(t *testing.T) {
	value := New(parse("Fri Jul 17 09:54:37 MST 2015")).BeginningOfYear()
	expected := parse("Thu Jan 1 00:00:00 MST 2015")
	assert.Equal(t, expected.String(), value.String())
}

func parse(s string) time.Time {
	t, _ := time.Parse(time.UnixDate, s)
	return t
}
