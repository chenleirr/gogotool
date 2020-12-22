package help

import (
	"fmt"
	"time"
)

func StrToTime(timeStr string) (time.Time, error) {
	if len(timeStr) != 19 {
		return time.Time{}, fmt.Errorf("StrToTime unexpect params length")
	}
	return time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
}

func TimeToStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}