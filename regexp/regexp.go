package regexp

import "regexp"

//IsAlpha 判断是否全部是英文
func IsAlpha(data string) (bool, error) {
	return regexp.MatchString("^[a-zA-Z]+$", data)
}

//IsAlpha2 判断是否全部是英文及英文句号
func IsAlpha2(data string) (bool, error) {
	return regexp.MatchString("^[a-z.A-Z]+$", data)
}

//IsAlphaAndNum2 判断是否全部是英文及英文下滑线
func IsAlphaAndNum2(data string) (bool, error) {
	return regexp.MatchString("^[a-z_0-9A-Z]+$", data)
}

//IsAlphaAndNum 判断是否全部是英文及数字
func IsAlphaAndNum(data string) (bool, error) {
	return regexp.MatchString("^[a-z0-9A-Z]+$", data)
}

//IsNum 判断是否全部是数字
func IsNum(data string) (bool, error) {
	return regexp.MatchString("^[0-9]+$", data)
}

//IsDate 判断是否为时间
func IsDate(data string) (bool, error) {
	return regexp.MatchString("^\\d{4}-\\d{1,2}-\\d{1,2}", data)
}

//IsMobilePhone 判断是否为手机号
func IsMobilePhone(data string) (bool, error) {
	return regexp.MatchString("^1[3456789]\\d{9}$", data)
}

//IsQQ 判断是否为QQ
func IsQQ(data string) (bool, error) {
	return regexp.MatchString("^\\d{5,}$", data)
}

//IsEmail 判断是否为Email
func IsEmail(data string) (bool, error) {
	return regexp.MatchString("^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*\\.[a-zA-Z0-9]{2,6}$", data)
}

//IsHHMM 小时:分钟
func IsHHMM(data string) (bool, error) {
	return regexp.MatchString("^([0-1][0-9]|2[0-3]):([0-5][0-9])$", data)
}

//IsHHMMSS 小时:分钟:秒
func IsHHMMSS(data string) (bool, error) {
	return regexp.MatchString("^([0-1][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$", data)
}

//IsDD 日期 max=28
func IsDD(data string) (bool, error) {
	return regexp.MatchString("^([0-1][0-9]|2[0-8])$", data)
}

//IsWeekly 周几
func IsWeekly(data string) (bool, error) {
	return regexp.MatchString("^([1-7])$", data)
}

//IsDateTime 2023-01-01 小时:分钟:秒
func IsDateTime(data string) (bool, error) {
	return regexp.MatchString("^\\d{4}-\\d{1,2}-\\d{1,2} ([0-1][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$", data)
}
