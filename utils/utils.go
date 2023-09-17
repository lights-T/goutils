package utils

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/now"
)

func ScanPort(protocol string, hostname string, port int) bool {
	p := strconv.Itoa(port)
	addr := net.JoinHostPort(hostname, p)
	conn, err := net.DialTimeout(protocol, addr, 3*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func replaceDescDate(_time time.Time) string {
	return strings.Replace(_time.Format("2006-01-02"), "-", "", -1)
}

func replaceDescTime(_time string) string {
	return strings.Replace(_time, ":", "", -1)
}

func GetDurationByCurrentTime(startDate string) (string, string, string, string) {
	startTime, _ := now.Parse(startDate)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	var dayStr, hourStr, minStr, secsStr string
	res := fmt.Sprintf("%v", duration)
	getDayArr := strings.Split(fmt.Sprintf("%s", res), "h")
	if len(getDayArr) > 1 {
		d, _ := strconv.Atoi(getDayArr[0])
		dayInt := d / 24
		dayStr = fmt.Sprintf("%d", d/24)
		hourStr = fmt.Sprintf("%d", d-(dayInt*24))
		//开始取得min
		getMinArr := strings.Split(fmt.Sprintf("%+v", getDayArr[1]), "m")
		if len(getMinArr) > 1 {
			minStr = getMinArr[0]
			getSecsArr := strings.Split(fmt.Sprintf("%+v", getMinArr[1]), "s")
			getSecsDArr := strings.Split(fmt.Sprintf("%+v", getSecsArr[0]), ".")
			secsStr = getSecsArr[0]
			if len(getSecsDArr) > 1 {
				secsStr = getSecsDArr[0]
			}
		}
		res = fmt.Sprintf("%sd%sh%s%ss", dayStr, hourStr, minStr, secsStr)
	} else {
		getMinArr := strings.Split(fmt.Sprintf("%s", res), "m")
		if len(getMinArr) > 1 {
			minStr = getMinArr[0]
			getSecsArr := strings.Split(fmt.Sprintf("%+v", getMinArr[1]), "s")
			getSecsDArr := strings.Split(fmt.Sprintf("%+v", getSecsArr[0]), ".")
			secsStr = getSecsArr[0]
			if len(getSecsDArr) > 1 {
				secsStr = getSecsDArr[0]
			}
		} else {
			getSecsArr := strings.Split(fmt.Sprintf("%s", res), "s")
			getSecsDArr := strings.Split(fmt.Sprintf("%+v", getSecsArr[0]), ".")
			secsStr = getSecsArr[0]
			if len(getSecsDArr) > 1 {
				secsStr = getSecsDArr[0]
			}
		}
	}
	if len(dayStr) == 0 {
		dayStr = "0"
	}
	if len(hourStr) == 0 {
		hourStr = "0"
	}
	if len(minStr) == 0 {
		minStr = "0"
	}
	if len(secsStr) == 0 {
		secsStr = "0"
	}
	return dayStr, hourStr, minStr, secsStr
}

func GetDurationByCurrentTimeToMin(startDate string) string {
	startTime, _ := now.Parse(startDate)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	return fmt.Sprintf("%d", int64(math.Ceil(duration.Minutes())+0/5))
}

//MinToDHM 分钟转化天、小时、分钟
func MinToDHM(min int) (int, int, int) {
	return min / 60 / 24 % 365, min / 60 % 24, min % 60
}

//ZnToUnicode 中文转unicode
func ZnToUnicode(str string) string {
	textQuoted := strconv.QuoteToASCII(str)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	return textUnquoted
}

//UnicodeToZn unicode转中文
func UnicodeToZn(str string) (string, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(str), `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	return strings.Replace(str, `\`, ``, -1), nil
}

func GetDurationToMin(startDate, endDate string) string {
	startTime, _ := now.Parse(startDate)
	endTime, _ := now.Parse(endDate)
	duration := endTime.Sub(startTime)
	return fmt.Sprintf("%d", int64(math.Ceil(duration.Minutes())+0/5))
}
