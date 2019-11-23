package dateutils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowUint() uint64 {
	return uint64(GetNow().UnixNano() / int64(time.Millisecond))
}
