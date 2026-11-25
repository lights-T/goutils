package xstrings

import (
	"fmt"
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
	A := "xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;...重复至超过100次..."
	//A := "xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;xxx;...重复至超过100次..."

	A, B := SplitData(A, 5)
	t.Log("A变量:", A)
	if B != "" {
		t.Log("B变量:", B)
	} else {
		t.Log("没有超出部分")
	}
}

func Test_ParsePath(t *testing.T) {
	cases := []string{
		`F:\ACMas\ACMSvcwe\ACMSvcws.exe`,
		`C:\ACM\ACMSvc\ACMSvc.exe`,
		`C:\ACM9\ACMSvc\ACMSvc9.exe`,
		`C:\ACM10\ACMSvc\ACMSvc10.exe`,
		`D:\ACMSvc3\ACMSvc\ACMSvc3.exe`,
		`E:\ACMSvc12\ACMSvc\ACMSvc12.exe`,
	}
	for _, c := range cases {
		d, n, err := GetSystemDrive(c)
		if err != nil {
			fmt.Println("解析失败:", err)
			continue
		}
		fmt.Printf("路径=%-35s -> 盘符=%s 数字=%s\n", c, d, n)
	}
}
