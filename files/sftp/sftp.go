package sftp

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var (
	auth         []ssh.AuthMethod
	addr         string
	clientConfig *ssh.ClientConfig
	sshClient    *ssh.Client
	sftpClient   *sftp.Client
	err          error
)

type Sftp struct {
	Ctx      *sftp.Client
	User     string
	Password string
	Host     string
	Port     int64
	Timeout  time.Duration
}

//Connect 连接SFTP服务器，用完手动必须关闭
func Connect(req *Sftp) (*Sftp, error) {
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

	// get auth method
	auth = make([]ssh.AuthMethod, 0, 5)
	auth = append(auth, ssh.Password(password))
	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         req.Timeout,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //ssh.FixedHostKey(hostKey),
	}

	// connect to ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		fmt.Println("Failed to Dial.", err)
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		fmt.Println("Failed to NewClient.", err)
		return nil, err
	}

	req.Ctx = sftpClient

	return req, nil
}

func (s *Sftp) Close() error {
	return s.Ctx.Close()
}

//Upload 本地->远程
func (s *Sftp) Upload(sourcePath, destPath string) error {
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
	if err = s.Ctx.Mkdir(destPath); err != nil {
		fmt.Println("Failed to Mkdir.", err)
	}

	if info.IsDir() {
		nextDestPath := path.Join(destPath, info.Name())
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
		destFileName := path.Join(destPath, info.Name())
		dstFile, err := s.Ctx.Create(destFileName)
		if err != nil {
			fmt.Println("Failed to Create.", err)
			return err
		}
		defer dstFile.Close()
		ff, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println("Failed to ioutil.ReadAll.", err)
			return err
		}
		if _, err = dstFile.Write(ff); err != nil {
			fmt.Println(f.Name(), " -> ", destFileName, "Failure")
			return err
		} else {
			fmt.Println(f.Name(), " -> ", destFileName, "Succeed")
		}
	}
	return nil
}

func (s *Sftp) CheckSFTPDir(sftpPath string) error {
	if _, err := s.Ctx.Stat(sftpPath); err != nil {
		return err
	}
	return nil
}
