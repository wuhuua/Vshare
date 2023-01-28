package util

import "time"

func GetDate() string {
	getYear := time.Now().Format("01")
	getDay := time.Now().Format("02")
	return getYear + "-" + getDay
}

func GetTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
