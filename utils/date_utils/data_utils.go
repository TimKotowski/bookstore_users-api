package date_utils

import "time"

const (
	apiDateLayout = "01-02-2006T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
