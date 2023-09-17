package logger

import (
	"testing"
)

func init() {
	conf := &Config{
		ConsoleLoggingEnabled: false,
		FileLoggingEnabled:    true,
		EncodeLogsAsJson:      false,
		Directory:             "./log_files",
		Filename:              "log.txt",
		MaxSize:               30,
		MaxBackups:            1,
		MaxAge:                30,
		Debug:                 false,
	}
	InitLogger(conf)
}

func TestNew(t *testing.T) {
	Infof("user:%v, test:%s", 111, "测试")
	Warnf("111")
	Error("发生了致命性错误")
}
