package datedsl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndOfSecond(t *testing.T) {
	value := New(parse("Fri Jul 17 10:13:50 MST 2015")).EndOfSecond()
	expected := parse("Fri Jul 17 10:13:50.999999999 MST 2015")
	assert.Equal(t, expected.String(), value.String())
}

func TestEndOfMinute(t *testing.T) {
	value := New(parse("Fri Jul 17 10:13:59 MST 2015")).EndOfMinute()
	expected := parse("Fri Jul 17 10:13:59.999999999 MST 2015")
	assert.Equal(t, expected.String(), value.String())
}

func TestEndOfHour(t *testing.T) {
	value := New(parse("Fri Jul 17 10:13:59 MST 2015")).EndOfHour()
	expected := parse("Fri Jul 17 10:59:59.999999999 MST 2015")
	assert.Equal(t, expected.String(), value.String())
}

func TestEndOfDay(t *testing.T) {
	value := New(parse("Fri Jul 17 10:13:59 MST 2015")).EndOfDay()
	expected := parse("Fri Jul 17 23:59:59.999999999 MST 2015")
	assert.Equal(t, expected.String(), value.String())
}

func TestEndOfMonth(t *testing.T) {
	value := New(parse("Fri Jul 17 09:54:37 MST 2015")).EndOfMonth()
	expected := parse("Wed Jul 31 23:59:59.999999999 MST 2015")
	assert.Equal(t, expected.String(), value.String())
}

func TestEndOfYear(t *testing.T) {
	value := New(parse("Fri Jul 17 09:54:37 MST 2015")).EndOfYear()
	expected := parse("Thu Dec 31 23:59:59.999999999 MST 2015")
	assert.Equal(t, expected.String(), value.String())
}
