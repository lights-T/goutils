package regexp

import "regexp"

func IsAlpha(data string) (bool, error) {
	return regexp.MatchString("^[a-zA-Z]+$", data)
}
