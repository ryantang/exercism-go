package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	formattedDate := t.Format("Monday, January 2, 2006")
	formattedTime := t.Format("15:04")

	message := fmt.Sprintf("You have an appointment on %s, at %s.", formattedDate, formattedTime)
	return message
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	openingDate := time.Date(2012, time.September, 15, 0, 0, 0, 0, time.UTC)
	anniversaryDate := time.Date(
		time.Now().Year(),
		openingDate.Month(),
		openingDate.Day(),
		openingDate.Hour(),
		openingDate.Minute(),
		openingDate.Second(),
		openingDate.Nanosecond(),
		openingDate.Location(),
	)

	return anniversaryDate
}
