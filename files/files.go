package files

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func FileHandler(filename, content string) error {
	var f *os.File
	var err error
	f, err = os.Create(filename)
	if err != nil {
		return err
	}
	_, err = io.WriteString(f, content)
	if err != nil {
		return err
	}

	return nil
}

//CheckFileIsExist 判断文件是否存在  存在返回 true 不存在返回false
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func CheckFileDir(path string, typeStr string) error {
	f, err := os.Stat(path)
	if err != nil {
		return err
	}
	if os.IsNotExist(err) {
		return err
	}

	switch typeStr {
	case "file":
		if f.IsDir() {
			return fmt.Errorf("file path is not a file, path:%s", path)
		}
	case "dir":
		if !f.IsDir() {
			return fmt.Errorf("file path is not a folder, path:%s", path)
		}
	}

	return nil
}

func CheckFileDirAndCreate(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		if err = os.Mkdir(path, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func ReadFile(fileName string) ([]byte, error) {
	var content []byte
	var err error
	file, err := os.Open(fileName)
	if err != nil {
		return content, err
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return content, err
	}
	return b, nil
}

func ReadIcon(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}
	data := make([]byte, stats.Size())
	if _, err = file.Read(data); err != nil {
		return nil, err
	}
	return data, nil
}

//GetRunPath 获取程序执行目录
func GetRunPath() (string, error) {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	return path, err
}

func CheckPath(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}
	if os.IsNotExist(err) {
		return err
	}
	return nil
}
