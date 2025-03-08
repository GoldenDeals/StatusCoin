package share

import "time"

func GetDateFormatted() string {
	return time.Now().Format("02-01-2006")
}

// GetTimeFormatted returns the current time in hh-mm-ss format (24-hour format).
func GetTimeFormatted() string {
	return time.Now().Format("15-04-05")
}
