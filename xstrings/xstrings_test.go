package xstrings

import (
	"testing"

	lfiles "github.com/lights-T/goutils/files"
)

func Test_BytesToStringByAscii(t *testing.T) {
	b, _ := lfiles.ReadIcon("../../../../logo.ico")
	t.Log(b)
	str := BytesToStringByAscii(b)
	t.Log(str)
}

func Test_SplitData(t *testing.T) {
	// 示例数据，假设这是A变量的数据
	A := "xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;...重复至超过100次..."

	A, B := SplitData(A, 100)
	t.Log("A变量:", A)
	if B != "" {
		t.Log("B变量:", B)
	} else {
		t.Log("没有超出部分")
	}
}
