package calculate

import "testing"

func Test_LimitPercent(t *testing.T) {
	res := LimitPercent(900)
	t.Log(res)
}
