package excel

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	path := "G:\\project\\emerson\\soft_project\\akzon-gui-vue\\Excel.xlsx"
	//data, err := ReadeExcel(path)
	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	//t.Log(data)

	t.Run("isExist", func(t *testing.T) {
		_, err := os.Stat(path)
		if err != nil {
			t.Fatal(err.Error())
		}
		//isnotexist来判断，是不是不存在的错误
		if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
			t.Fatal(err.Error())
		} else {
			t.Log(true)
		}
	})
}

func TestReadExcel(t *testing.T) {
	filePath := ""
	cols, list, err := ReadeExcel(filePath, false)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	t.Log(cols)
	t.Log(list)
	//for _, v := range list {
	//t.Logf("第一行"v["名称"])
	//}
}
