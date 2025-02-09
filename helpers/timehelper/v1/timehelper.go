package timehelper

import "time"

func Seconds(seconds int) time.Duration {
	return time.Duration(seconds) * time.Second
}