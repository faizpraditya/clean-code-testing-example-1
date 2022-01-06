package utils

import "time"

func StringtoDate(stringDate string, format string) time.Time {
	dateTime, _ := time.Parse(format, stringDate)

	return dateTime
}
