package time

import (
	"testing"
	"time"

	"github.com/jinzhu/now"
	"github.com/lights-T/goutils/constant"
)

func Test_GetDiffDays(t *testing.T) {
	t1, err := now.Parse("2024-03-12 15:12:12")
	if err != nil {
		t.Fatal(err.Error())
	}
	t2, err := now.Parse("2024-03-11 15:12:12")
	if err != nil {
		t.Fatal(err.Error())
	}
	d := GetDiffDays(t1, t2) //1
	t.Fatal(d)
}

func Test_GetBeforeTime(t *testing.T) {
	now := "2024-03-1 4:10:10"
	sec := int64(8 * 60 * 60)
	t.Log(GetBeforeTime(now, sec).Format(constant.DatetimeLayoutNa))
}

func Test_GetAfterTime(t *testing.T) {
	now := "2024-03-1 4:10:10"
	sec := int64(8 * 60 * 60)
	t.Log(GetAfterTime(now, sec).Format(constant.DatetimeLayoutNa))
}

func Test_GetAfterTimeByNow(t *testing.T) {
	sec := int64(8 * 60 * 60)
	t.Log(GetAfterTimeByNow(sec).Format(constant.DatetimeLayoutNa))
}

func Test_GetDurationToSec(t *testing.T) {
	t.Log(GetDurationToSec("2024-01-02 12:14:10.641464", "2024-01-02 12:14:36.625598"))
}

func Test(t *testing.T) {
	//currentDate: 2026-01-01 00:00:00.016, beforeDay: 2025-12-01 00:00:00.016, reportTotalTime: 2678400
	//currentDate: 2026-02-01 00:00:00.016, beforeDay: 2026-01-01 00:00:00.016, reportTotalTime: 2678400
	//currentDate: 2026-03-01 00:00:00.016, beforeDay: 2026-02-01 00:00:00.016, reportTotalTime: 2419200
	//currentDate: 2026-04-01 00:00:00.016, beforeDay: 2026-03-01 00:00:00.016, reportTotalTime: 2678400
	//currentDate: 2026-05-01 00:00:00.016, beforeDay: 2026-04-01 00:00:00.016, reportTotalTime: 2592000
	//currentDate: 2026-06-01 00:00:00.016, beforeDay: 2026-05-01 00:00:00.016, reportTotalTime: 2678400
	//currentDate: 2026-07-01 00:00:00.016, beforeDay: 2026-06-01 00:00:00.016, reportTotalTime: 2592000
	//currentDate: 2026-08-01 00:00:00.016, beforeDay: 2026-07-01 00:00:00.016, reportTotalTime: 2678400
	//currentDate: 2026-09-01 00:00:00.016, beforeDay: 2026-08-01 00:00:00.016, reportTotalTime: 2678400
	//currentDate: 2026-10-01 00:00:00.016, beforeDay: 2026-09-01 00:00:00.016, reportTotalTime: 2592000
	//currentDate: 2026-11-01 00:00:00.016, beforeDay: 2026-10-01 00:00:00.016, reportTotalTime: 2678400
	//currentDate: 2026-12-01 00:00:00.016, beforeDay: 2026-11-01 00:00:00.016, reportTotalTime: 2592000
	//currentDate: 2026-12-01 10:06:11.016, beforeDay: 2026-11-01 10:06:11.016, reportTotalTime: 2592000
	tt, err := time.Parse("2006-01-02 15:04:05.999", "2026-12-01 10:06:11.016")
	if err != nil {
		panic(err)
	}
	//month := -1
	//day := 0

	currentUtcDateTime, _ := GetForwardOffsetDate(tt.Format(constant.DatetimeLayoutNa), 8)
	currentDate := currentUtcDateTime.Format(constant.DatetimeLayoutNa)
	//beforeDayUnix := tt.AddDate(0, month, day)
	//beforeDay := beforeDayUnix.Format(constant.DatetimeLayoutNa)
	//reportTotalTime := tt.Unix() - beforeDayUnix.Unix()
	//t.Logf("===> export excel: currentDate: %s, beforeDay: %s, reportTotalTime: %d", currentDate, beforeDay, reportTotalTime)

	lastDayTime := GetMonthsAgo(tt, 1)
	reportTotalTime2 := tt.Unix() - lastDayTime.Unix()
	lastDayUtcTime, _ := GetForwardOffsetDate(lastDayTime.Format(constant.DatetimeLayoutNa), 8)
	lastDay := lastDayUtcTime.Format(constant.DatetimeLayoutNa)
	t.Logf("===> export excel: currentDate: %s, beforeDay: %s, reportTotalTime: %d", tt.Format(constant.DatetimeLayoutNa), lastDayTime.Format(constant.DatetimeLayoutNa), reportTotalTime2)
	t.Logf("===> export excel: utc-currentDate: %s, utc-beforeDay: %s, reportTotalTime: %d", currentDate, lastDay, reportTotalTime2)
}
