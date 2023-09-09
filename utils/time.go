package utils

import "time"

func Timestamp(dt time.Time) string {
	return dt.Format("2006_01_02_15_04_05")
}

func TimestampNow() string {
	return time.Now().Format("2006_01_02_15_04_05")
}
