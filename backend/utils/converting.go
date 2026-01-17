package utils

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

func FormatDate(t time.Time) string {
	return t.Format("02 Jan 2006")
}

func ParseDate(str string) time.Time {
	layout := "2006-01-02"
	tanggal, err := time.Parse(layout, str)
	if err != nil {
		panic(err)
	}
	return tanggal
}

func StringToUUID(str string) (uuid.UUID, error) {
	return uuid.Parse(str)
}

func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}
