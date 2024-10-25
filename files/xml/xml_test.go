package xml

import (
	"testing"

	dxml "github.com/lights-T/goutils/domain/xml"
)

func Test_Update(t *testing.T) {
	xmlName := "Web.Config"
	// 解码XML数据
	var config dxml.Configuration
	// 修改XML数据
	config.AppBaseApi = "http://127.0.0.1:8090"
	config.AppServicePort = 8090
	config.AppDB = &dxml.AppDB{
		ServerIP: "192.168.56.211",
		Port:     1433,
		DBName:   "UMP",
		UserName: "ACM",
		Password: "pwd",
	}
	config.Email = &dxml.Email{
		Server:   "smtp.qq.com",
		Port:     587,
		UserName: "xxx@qq.com",
		Password: "xxx",
	}
	if err := Update(xmlName, config); err != nil {
		t.Fatal(err.Error())
	}
	t.Log("XML文件已成功修改")
}
