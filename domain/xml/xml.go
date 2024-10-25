package xml

import "encoding/xml"

// Configuration 定义XML主体结构体
type Configuration struct {
	XMLName        xml.Name `xml:"configuration"`
	AppServicePort int      `xml:"appServicePort"`
	AppBaseApi     string   `xml:"appBaseApi"`
	AppDB          *AppDB   `xml:"appDB"`
	Email          *Email   `xml:"emailConfig"`
}

type AppDB struct {
	XMLName  xml.Name `xml:"appDB"`
	ServerIP string   `xml:"serverIP"`
	Port     int      `xml:"port"`
	DBName   string   `xml:"dbName"`
	UserName string   `xml:"userName"`
	Password string   `xml:"password"`
}

type Email struct {
	XMLName  xml.Name `xml:"emailConfig"`
	Server   string   `xml:"server"`
	Port     int      `xml:"port"`
	UserName string   `xml:"userName"`
	Password string   `xml:"password"`
}
