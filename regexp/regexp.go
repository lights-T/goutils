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
