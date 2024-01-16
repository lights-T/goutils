package time

import (
	"time"

	"github.com/jinzhu/now"
	"github.com/lights-T/goutils/constant"
)

func GetCurrentTimeByNa() string {
	return time.Now().Format(constant.DatetimeLayoutNa)
}

func GetCurrentTime() string {
	return time.Now().Format(constant.DatetimeLayout)
}

func GetFormatDate(date string) string {
	if len(date) == 0 {
		return ""
	}
	_time, _ := now.Parse(date)
	return _time.Format(constant.DatetimeLayout)
}
