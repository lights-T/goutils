package xml

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"os"

	"github.com/lights-T/goutils"
	"github.com/lights-T/goutils/files"
)

func Create(xmlName string, v any) error {
	// 重新序列化XML数据
	updatedData, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return goutils.Errorf("XML数据序列化失败: %s", err.Error())
	}

	// 写入修改后的XML数据
	if err = files.WriteFileByOptional(xmlName, os.O_CREATE, xml.Header+string(updatedData)); err != nil {
		return goutils.Errorf("写入XML文件失败: %s", err.Error())
	}
	return nil
}

func Update(xmlName string, v any) error {
	// 重新序列化XML数据
	updatedData, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return goutils.Errorf("XML数据序列化失败: %s", err.Error())
	}

	// 写入修改后的XML数据
	if err = ioutil.WriteFile(xmlName, []byte(xml.Header+string(updatedData)), 0644); err != nil {
		return goutils.Errorf("写入XML文件失败: %s", err.Error())
	}
	return nil
}

func Read(filePath string, v interface{}) (err error) {
	// 打开XML文件
	file, err := os.Open(filePath)
	if err != nil {
		return goutils.Errorf("无法打开文件: %s", err.Error())
	}
	defer file.Close()

	// 读取文件内容到字节切片
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return goutils.Errorf("读取文件内容时发生错误: %s", err.Error())
	}
	// 创建解码器并解析XML内容
	decoder := xml.NewDecoder(bytes.NewBuffer(byteValue))
	if err = decoder.Decode(v); err != nil {
		return goutils.Errorf("解码XML时发生错误: %s", err.Error())
	}
	return nil
}
