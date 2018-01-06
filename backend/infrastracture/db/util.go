package db

import (
	"strings"
	"time"
)

func escape(str string) string {
	return strings.Replace(str, "'", "\\'", -1)
}

func toStrTime(date time.Time) string {
	layout := "2006-01-02 15:04:05"
	return date.Format(layout)
}
