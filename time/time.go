package time

import (
	"fmt"
	"math"
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

func GetFormatDateTime(date string) string {
	if len(date) == 0 {
		return ""
	}
	_time, _ := now.Parse(date)
	return _time.Format(constant.DatetimeLayout)
}

func GetFormatDateTimeNa(date string) string {
	if len(date) == 0 {
		return ""
	}
	_time, _ := now.Parse(date)
	return _time.Format(constant.DatetimeLayoutNa)
}

func GetFormatDate(date string) string {
	if len(date) == 0 {
		return ""
	}
	_time, _ := now.Parse(date)
	return _time.Format(constant.DateLayout)
}

//GetDiffDays 获取两个时间相差的天数，0表同一天，正数表t1>t2，负数表t1<t2
func GetDiffDays(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)
	return int(t1.Sub(t2).Hours() / 24)
}

//GetDiffDaysBySecond 获取t1和t2的相差天数，单位：秒，0表同一天，正数表t1>t2，负数表t1<t2
func GetDiffDaysBySecond(t1, t2 int64) int {
	time1 := time.Unix(t1, 0)
	time2 := time.Unix(t2, 0)

	// 调用上面的函数
	return GetDiffDays(time1, time2)
}

func GetBeforeTime(target string, sec int64) time.Time {
	now, _ := now.Parse(target)
	return now.Add(-time.Second * time.Duration(sec))
}

func GetAfterTime(target string, sec int64) time.Time {
	now, _ := now.Parse(target)
	return now.Add(time.Second * time.Duration(sec))
}

func GetAfterTimeByNow(sec int64) time.Time {
	return time.Now().Add(time.Second * time.Duration(sec))
}

//GetDurationToSec 两者之差，秒
func GetDurationToSec(startDate, endDate string) string {
	startTime, _ := now.Parse(startDate)
	endTime, _ := now.Parse(endDate)
	duration := endTime.Sub(startTime)
	return fmt.Sprintf("%d", int64(math.Ceil(duration.Seconds())+0/5))
}

func GetForwardOffsetDate(dateReq string, offset int) (time.Time, error) {
	tt, err := time.Parse("2006-01-02 15:04:05.999", dateReq)
	if err != nil {
		return tt, err
	}
	sth, _ := time.ParseDuration(fmt.Sprintf("-%dh", offset))
	tt = tt.Add(sth)
	return tt, nil
}

func GetBackwardOffsetDate(dateReq string, offset int) (time.Time, error) {
	tt, err := time.Parse("2006-01-02 15:04:05.999", dateReq)
	if err != nil {
		return tt, err
	}
	sth, _ := time.ParseDuration(fmt.Sprintf("%dh", offset))
	tt = tt.Add(sth)
	return tt, nil
}

func GetMonthsAgo(t time.Time, monthsBack int) time.Time {
	// 1. 计算目标年份和月份
	// 注意：这里要处理好跨年，比如 1月 - 2个月 = 去年11月
	targetYear := t.Year()
	targetMonth := t.Month() - time.Month(monthsBack)

	for targetMonth <= 0 {
		targetMonth += 12
		targetYear--
	}

	// 2. 核心逻辑：尝试构建目标日期的“同一天”
	// 例如：t是3月31日，target是2月。Go会尝试构建 2月31日。
	// Go的time.Date会自动进位：2月31日 -> 3月3日（非闰年）
	tempTime := time.Date(targetYear, targetMonth, t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())

	// 3. 检查是否发生了“月份溢出”
	// 如果我们想要2月，结果却跑到了3月，说明目标月份没有这一天（如2月没有31号）
	if tempTime.Month() != targetMonth {
		// 解决方法：将日期设为目标月份的“第0天”，即上个月的最后一天（也就是目标月的最后一天）
		// 例如：3月0日 = 2月最后一天
		return time.Date(targetYear, targetMonth+1, 0, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	}

	// 如果没有溢出，直接返回
	return tempTime
}
