package files

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lights-T/goutils/constant"
	"github.com/lights-T/goutils/domain"
)

func FileHandler(filename, content string) error {
	var f *os.File
	var err error
	f, err = os.Create(filename)
	if err != nil {
		return err
	}
	if _, err = io.WriteString(f, content); err != nil {
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

// DownloadFile 文件下载
func DownloadFile(filePath string, rw http.ResponseWriter) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	filename := path.Base(filePath)
	rw.Header().Set("Content-Type", "application/octet-stream")
	rw.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	_, err = io.Copy(rw, file)

	return err
}

// CreateDirIfNotExists 检测目录是否存在
func CreateDirIfNotExists(path ...string) error {
	for _, value := range path {
		if FileExist(value) {
			continue
		}
		err := os.Mkdir(value, 0755)
		if err != nil {
			return fmt.Errorf("创建目录失败:%s", err.Error())
		}
	}
	return nil
}

// FileExist 判断文件是否存在及是否有权限访问
func FileExist(file string) bool {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	if os.IsPermission(err) {
		return false
	}

	return true
}

func Upload(ctx *gin.Context, fileDir string, req string) (*domain.UploadFilesInfo, map[string][]string, error) {
	if len(req) == 0 {
		req = "file"
	}
	filesInfo := &domain.UploadFilesInfo{}
	var value map[string][]string
	//获取所有上传文件信息
	form, err := ctx.MultipartForm()
	if err != nil {
		return filesInfo, value, err
	}
	if form == nil {
		return filesInfo, value, fmt.Errorf(constant.ErrUploadParamsIsNotExist)
	}
	value = form.Value
	fhs := form.File[req]
	if len(fhs) == 0 {
		return filesInfo, value, fmt.Errorf(constant.ErrUploadFileIsNotExist)
	}
	if err := CheckFileDirAndCreate(fileDir); err != nil {
		return filesInfo, value, err
	}
	var filePaths []string
	var fileNames []string
	var fileSizes []int64
	for _, f := range fhs {
		currentTime := strconv.FormatInt(time.Now().UnixNano(), 10)
		if len(strings.Split(f.Filename, ".")) < 2 {
			return filesInfo, value, fmt.Errorf("图片不合规")
		}
		filePath := fmt.Sprintf("%s/%s", fileDir, currentTime+fmt.Sprintf(".%s", strings.Split(f.Filename, ".")[1]))
		if err := ctx.SaveUploadedFile(f, filePath); err != nil {
			return filesInfo, value, err
		}
		filePaths = append(filePaths, filePath)
		fileNames = append(fileNames, f.Filename)
		fileSizes = append(fileSizes, f.Size)
	}

	filesInfo.FileNames = fileNames
	filesInfo.FilePaths = filePaths
	filesInfo.FileSizes = fileSizes

	return filesInfo, value, nil
}

func Remove(path string) error {
	return os.Remove(path)
}

//WriteFile 自动创建或追加写文件
func WriteFile(filename, content string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if _, err = io.WriteString(f, content); err != nil {
		return err
	}

	return nil
}
