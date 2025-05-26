package timeutils

import (
	"fmt"
	"strconv"
	"time"
)

func ParseTime(s string) (time.Time, error) {
	if s == "" {
		return time.Time{}, nil
	}

	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func ParseNullTime(s string) (*time.Time, error) {
	if s == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func StartOfToday() time.Time {
	return StartOfDay(time.Now().UTC())
}

func EndOfToday() time.Time {
	return EndOfDay(time.Now().UTC())
}

func StartOfDay(t time.Time) time.Time {
	y, m, d := t.UTC().Date()

	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

func EndOfDay(t time.Time) time.Time {
	y, m, d := t.UTC().Date()

	return time.Date(y, m, d, 23, 59, 59, 999999999, time.UTC)
}

func NextSaturdayAt8PM(location *time.Location) time.Time {
	if location == nil {
		location = time.UTC
	}

	// Get the current time in the user's timezone
	now := time.Now().In(location)

	// If it is already Saturday, and not 20:00 yet, return the current day at 20:00
	if now.Weekday() == time.Saturday && now.Hour() < 20 {
		return time.Date(now.Year(), now.Month(), now.Day(), 20, 0, 0, 0, location)
	}

	// Calculate the number of days until the next Saturday
	daysUntilNext := (6 - int(now.Weekday()) + 7) % 7
	if daysUntilNext == 0 {
		// If today is already Saturday, set the target to the next Saturday (7 days later)
		daysUntilNext = 7
	}

	// Calculate the next Saturday
	next := now.AddDate(0, 0, daysUntilNext)

	// Set the time to 8 PM (20:00) on that Saturday
	return time.Date(next.Year(), next.Month(), next.Day(), 20, 0, 0, 0, location)
}

// LocationFromOffset creates a location from an offset string.
// Format: [+|-]\d\d\d\d
func LocationFromOffset(offset string) *time.Location {
	if len(offset) != 5 {
		return time.UTC
	}

	if offset[0] != '+' && offset[0] != '-' {
		return time.UTC
	}

	return time.FixedZone(fmt.Sprintf("UTC%s", offset), secondsForOffset(offset))
}

func secondsForOffset(offset string) int {
	negative := offset[0] == '-'

	hi, err := strconv.ParseInt(offset[1:3], 10, 64)
	if err != nil {
		return 0
	}

	mi, err := strconv.ParseInt(offset[3:], 10, 64)
	if err != nil {
		return 0
	}

	ti := (hi * 60 * 60) + (mi * 60)

	if negative {
		ti *= -1
	}

	return int(ti)
}
