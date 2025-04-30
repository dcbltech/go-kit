package timeutils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseTime(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Time
		hasError bool
	}{
		{"2025-04-30T12:34:56Z", time.Date(2025, 4, 30, 12, 34, 56, 0, time.UTC), false},
		{"", time.Time{}, false},
		{"invalid", time.Time{}, true},
	}

	for _, test := range tests {
		result, err := ParseTime(test.input)
		if test.hasError {
			assert.Error(t, err, "Expected an error for input %q", test.input)
		} else {
			assert.NoError(t, err, "Did not expect an error for input %q", test.input)
			assert.Equal(t, test.expected, result, "Unexpected result for input %q", test.input)
		}
	}
}

func TestParseNullTime(t *testing.T) {
	tests := []struct {
		input    string
		expected *time.Time
		hasError bool
	}{
		{"", nil, false},
	}

	for _, test := range tests {
		result, err := ParseNullTime(test.input)
		if test.hasError {
			assert.Error(t, err, "Expected an error for input %q", test.input)
		} else {
			assert.NoError(t, err, "Did not expect an error for input %q", test.input)
			assert.Equal(t, test.expected, result, "Unexpected result for input %q", test.input)
		}
	}
}

func TestStartOfToday(t *testing.T) {
	expected := StartOfDay(time.Now().UTC())
	result := StartOfToday()
	assert.Equal(t, expected, result, "Unexpected result for StartOfToday")
}

func TestEndOfToday(t *testing.T) {
	expected := EndOfDay(time.Now().UTC())
	result := EndOfToday()
	assert.Equal(t, expected, result, "Unexpected result for EndOfToday")
}

func TestStartOfDay(t *testing.T) {
	timeInput := time.Date(2025, 4, 30, 15, 45, 0, 0, time.UTC)
	expected := time.Date(2025, 4, 30, 0, 0, 0, 0, time.UTC)
	result := StartOfDay(timeInput)
	assert.Equal(t, expected, result, "Unexpected result for StartOfDay")
}

func TestEndOfDay(t *testing.T) {
	timeInput := time.Date(2025, 4, 30, 15, 45, 0, 0, time.UTC)
	expected := time.Date(2025, 4, 30, 23, 59, 59, 999999999, time.UTC)
	result := EndOfDay(timeInput)
	assert.Equal(t, expected, result, "Unexpected result for EndOfDay")
}

func TestNextSaturdayAt8PM(t *testing.T) {
	location := time.UTC
	expected := time.Date(2025, 5, 3, 20, 0, 0, 0, location)

	result := NextSaturdayAt8PM(location)
	assert.Equal(t, expected, result, "Unexpected result for NextSaturdayAt8PM")
}

func TestLocationFromOffset(t *testing.T) {
	tests := []struct {
		input    string
		expected *time.Location
	}{
		{input: "+0200", expected: time.FixedZone("UTC+0200", 2*60*60)},
		{input: "-0500", expected: time.FixedZone("UTC-0500", -5*60*60)},
		{input: "invalid", expected: time.UTC},
	}

	for _, test := range tests {
		result := LocationFromOffset(test.input)
		assert.Equal(t, test.expected.String(), result.String(), "Unexpected result for LocationFromOffset with input %q", test.input)
	}
}
