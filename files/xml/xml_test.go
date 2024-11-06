package xml

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"os"
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
		Server:       "smtp.qq.com",
		Port:         587,
		UserName:     "xxx@qq.com",
		Password:     "xxx",
		AutoSendTime: 7,
	}
	if err := Update(xmlName, config); err != nil {
		t.Fatal(err.Error())
	}
	t.Log("XML文件已成功修改")
}

func Test_Read2(t *testing.T) {
	xmlName := "Web.Config"
	// 解码XML数据
	var config = &dxml.Configuration{}
	if err := Read(xmlName, config); err != nil {
		t.Fatal(err.Error())
	}
	t.Log(config)
	t.Log(config.AppBaseApi)
	t.Log(config.AppServicePort)
	t.Log(config.XMLName)
	t.Log(config.AppDB)
	t.Log(config.Email)
}

func Test_Read(t *testing.T) {
	xmlName := "Web.Config"
	// 解码XML数据
	var config dxml.Configuration
	// 打开XML文件
	file, err := os.Open(xmlName)
	if err != nil {
		t.Logf("无法打开文件: %s", err.Error())
	}
	defer file.Close()

	// 读取文件内容到字节切片
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		t.Logf("读取文件内容时发生错误: %s", err.Error())
	}
	// 创建解码器并解析XML内容
	decoder := xml.NewDecoder(bytes.NewBuffer(byteValue))
	if err = decoder.Decode(&config); err != nil {
		t.Logf("解码XML时发生错误: %s", err.Error())
	}
	t.Log(config)
	t.Log(config.AppBaseApi)
	t.Log(config.AppServicePort)
	t.Log(config.XMLName)
	t.Log(config.AppDB)
	t.Log(config.Email)
}
