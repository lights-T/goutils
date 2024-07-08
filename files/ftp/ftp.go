package ftp

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/lights-T/goutils"
)

var (
	err  error
	addr string
)

type Ftp struct {
	Ctx               *ftp.ServerConn
	User              string
	Password          string
	Host              string
	Port              int64
	Timeout           time.Duration
	AutoCreateDestDir bool //是否自动创建目标文件夹
}

//Connect 连接FTP服务器，用完手动必须关闭
func Connect(req *Ftp) (*Ftp, error) {
	user := req.User
	password := req.Password
	host := req.Host
	port := req.Port
	if req.Timeout == 0 {
		req.Timeout = 30 * time.Second
	}
	if port == 0 {
		port = 21
	}

	// 连接到FTP服务器
	addr = fmt.Sprintf("%s:%d", host, port)
	client, err := ftp.Dial(addr)
	if err != nil {
		fmt.Println("Failed to ftp.Dial.", err)
		return req, err
	}
	//defer client.Quit()

	// 登录
	if err = client.Login(user, password); err != nil {
		fmt.Println("Failed to client.Login.", err)
		return req, err
	}

	req.Ctx = client

	return req, nil
}

func (s *Ftp) Close() error {
	return s.Ctx.Quit()
}

//Upload 本地->远程
//ftp目录支持多级目录，但是只支持创建多级目录的上级目录
//被上传目录不支持同时上传多级目录
//文件支持重复上传，会覆盖
func (s *Ftp) Upload(sourcePath, remotePath string) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Local to remote, Program abnormal exit.", err)
			for i := 0; ; i++ {
				pc, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				fmt.Println("pc: ", pc, "file: ", file, "line: ", line)
			}
			return
		}
	}()

	f, err := os.Open(sourcePath)
	if err != nil {
		fmt.Println("Failed to os.Open.", err)
		return err
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		fmt.Println("Failed to f.Stat.", err)
		return err
	}

	// 创建对应目录
	if err = s.Ctx.MakeDir(remotePath); err != nil {
		fmt.Println("Failed to MakeDir.", err)
	}

	if info.IsDir() {
		nextDestPath := path.Join(remotePath, info.Name())
		localFiles, _ := ioutil.ReadDir(sourcePath)
		for _, localFile := range localFiles {
			nextSourcePath := path.Join(sourcePath, localFile.Name())
			if err = s.Upload(nextSourcePath, nextDestPath); err != nil {
				fmt.Println("Failed to upload.", err)
				return err
			}
		}
	} else {
		info, err = f.Stat()
		if err != nil {
			fmt.Println("Failed to f.Stat.", err)
			return err
		}
		destFileName := path.Join(remotePath, info.Name())
		if err = s.Ctx.Stor(destFileName, f); err != nil {
			fmt.Println("Failed to Stor.", err)
			return err
		}
	}
	return nil
}

//Download 远程->本地
//本地文件已存在，下载会被覆盖
//不支持下载文件夹
func (s *Ftp) Download(remotePath, localPath string) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Local to remote, Program abnormal exit.", err)
			for i := 0; ; i++ {
				pc, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				fmt.Println("pc: ", pc, "file: ", file, "line: ", line)
			}
			return
		}
	}()

	// 创建本地文件
	file, err := os.Create(localPath)
	if err != nil {
		fmt.Println("Failed to os.Create.", err)
		return err
	}
	defer file.Close()
	// 下载文件
	r, err := s.Ctx.Retr(remotePath)
	if err != nil {
		fmt.Println("Failed to Retr.", err)
		return err
	}
	defer r.Close()
	buf, err := io.ReadAll(r)
	if err != nil {
		fmt.Println("Failed to io.ReadAll.", err)
		return err
	}
	if _, err = file.Write(buf); err != nil {
		fmt.Println("Failed to Write.", err)
		return err
	}

	return nil
}

//CheckFTPDir 返回错误为不存在，反之存在
func (s *Ftp) CheckFTPDir(ftpPath string) error {
	// 列出当前目录下的文件和文件夹
	entries, err := s.Ctx.NameList(".")
	if err != nil {
		return err
	}
	ftpPath = strings.Trim(ftpPath, "./")

	// 检查文件夹是否存在
	found := false
	for _, entry := range entries {
		fmt.Println(entry)
		if strings.Contains(entry, ftpPath) {
			found = true
			break
		}
	}

	if found {
		return goutils.Errorf("The path already exists.")
	}
	return nil
}
