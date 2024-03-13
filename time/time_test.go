package time

import (
	"github.com/jinzhu/now"
	"github.com/lights-T/goutils/constant"
	"testing"
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
