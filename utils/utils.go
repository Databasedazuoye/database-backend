package utils

import (
	"strconv"
	"time"
)

func StringToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func FormatTimeString(timeString string) string {
	t, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		panic(err)
	}
	formattedTime := t.Format("2006-01-02 15:04:05")
	return formattedTime
}

type Array struct {
	Array []int64
}
