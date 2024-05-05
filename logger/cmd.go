package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

//CmdOutput cmd标准输出
func CmdOutput(level logrus.Level) {
	if level == 0 {
		level = logrus.InfoLevel
	}
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	writers := []io.Writer{os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	logrus.SetOutput(fileAndStdoutWriter)
	logrus.SetLevel(level)
}

//CmdOutput cmd标准输出
//func CmdOutput(req *CmdOutputReq) {
//	logPath := req.LogPath
//	level := req.Level
//	saveMaxDay := req.SaveMaxDay
//	isWriteFile := req.IsWriteFile
//	if level == 0 {
//		level = logrus.InfoLevel
//	}
//	if saveMaxDay == 0 {
//		saveMaxDay = 365
//	}
//	logrus.SetFormatter(&logrus.TextFormatter{})
//	logrus.SetOutput(os.Stdout)
//	writers := []io.Writer{os.Stdout}
//	if isWriteFile {
//		//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
//		file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
//		fileInfo, err := file.Stat()
//		if err != nil {
//			logrus.Errorf("Failed to get file information. [%s]", err.Error())
//			return
//		}
//		if time.Now().Unix()-fileInfo.ModTime().Unix() > saveMaxDay {
//			if err = files.Remove(logPath); err != nil {
//				logrus.Errorf("Failed to remove file. [%s]", err.Error())
//				return
//			}
//			//删除后重新创建
//			file, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
//			if err != nil {
//				logrus.Errorf("Failed to log to file. [%s]", logPath)
//			}
//		}
//		writers = append(writers, file)
//	}
//	fileAndStdoutWriter := io.MultiWriter(writers...)
//	logrus.SetOutput(fileAndStdoutWriter)
//	logrus.SetLevel(level)
//}
