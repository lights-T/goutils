package goutils

import (
	"bytes"
	"fmt"
	"html/template"
	"math"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
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

// RandNumber 生成min - max之间的随机数
// 如果min大于max, panic
func RandNumber(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	switch {
	case min == max:
		return min
	case min > max:
		panic("min must be less than or equal to max")
	}

	return min + r.Intn(max-min)
}

// PanicToError Panic转换为error
func PanicToError(f func()) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf(PanicTrace(e))
		}
	}()
	f()
	return
}

// PanicTrace panic调用链跟踪
func PanicTrace(err interface{}) string {
	stackBuf := make([]byte, 4096)
	n := runtime.Stack(stackBuf, false)

	return fmt.Sprintf("panic: %v %s", err, stackBuf[:n])
}

// PrintAppVersion 打印应用版本
func PrintAppVersion(appVersion, GitCommit, BuildDate string) {
	versionInfo, err := FormatAppVersion(appVersion, GitCommit, BuildDate)
	if err != nil {
		panic(err)
	}
	fmt.Println(versionInfo)
}

// FormatAppVersion 格式化应用版本信息
func FormatAppVersion(appVersion, GitCommit, BuildDate string) (string, error) {
	content := `
   Version: {{.Version}}
Go Version: {{.GoVersion}}
Git Commit: {{.GitCommit}}
     Built: {{.BuildDate}}
   OS/ARCH: {{.GOOS}}/{{.GOARCH}}
`
	tpl, err := template.New("version").Parse(content)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, map[string]string{
		"Version":   appVersion,
		"GoVersion": runtime.Version(),
		"GitCommit": GitCommit,
		"BuildDate": BuildDate,
		"GOOS":      runtime.GOOS,
		"GOARCH":    runtime.GOARCH,
	})
	if err != nil {
		return "", err
	}

	return buf.String(), err
}

// WorkDir 获取程序运行时根目录
func WorkDir() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}
	wd := filepath.Dir(execPath)
	if filepath.Base(wd) == "bin" {
		wd = filepath.Dir(wd)
	}
	return wd, nil
}

// WaitGroupWrapper waitGroup包装
type WaitGroupWrapper struct {
	sync.WaitGroup
}

// Wrap 包装Add, Done方法
func (w *WaitGroupWrapper) Wrap(f func()) {
	w.Add(1)
	go func() {
		defer w.Done()
		f()
	}()
}

func Errorf(format string, a ...any) error {
	return fmt.Errorf(format, a)
}
