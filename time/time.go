package time

import (
	"time"

	"github.com/lights-T/goutils/constant"
)

func GetCurrentTimeByNa() string {
	return time.Now().Format(constant.DatetimeLayoutNa)
}

func GetCurrentTime() string {
	return time.Now().Format(constant.DatetimeLayout)
}
