package utils

import "time"

func MillSecondTime(millSecs int64) time.Time {
	//timestampMilli := time.Now().Add(time.Hour * 24).UnixMilli()
	seconds := millSecs / 1000
	nano := (millSecs % 1000) * 1e6
	return time.Unix(seconds, nano)
}
