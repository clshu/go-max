package util

import "time"

// GetDBTimeFormat return a format to match database time format in time.Time structure
//
// e.g. "2025-04-18T09:00:57Z"
func GetDBTimeFormat(timestamp time.Time) time.Time {
	// Convert to UTC to get suffix Z, truncate to seconds
	return timestamp.UTC().Truncate(time.Second)
}
