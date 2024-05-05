package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

//ForCmdAnfFile 写入cmd及文件
func ForCmdAnfFile(logPath string, level logrus.Level) {
	if level == 0 {
		level = logrus.InfoLevel
	}
	//设置输出样式
	//自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	writers := []io.Writer{file, os.Stdout}
	//同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(writers...)
	if err == nil {
		logrus.SetOutput(fileAndStdoutWriter)
	} else {
		logrus.Infof("Failed to log to file. [%s]", logPath)
	}
	//设置最低loglevel
	logrus.SetLevel(level)
}
