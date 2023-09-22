package regexp

import "regexp"

//IsAlpha 判断是否全部是英文
func IsAlpha(data string) (bool, error) {
	return regexp.MatchString("^[a-zA-Z]+$", data)
}
