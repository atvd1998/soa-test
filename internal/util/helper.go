package util

import "time"

func FormatDateTime(dateTime *time.Time, format string) string {
	return dateTime.Format(format)
}
