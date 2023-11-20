package net

import (
	"fmt"
	"net"
	"os/exec"
	"time"
)

// CheckPort Golang中的net标准库来检查特定主机和端口是否可连接,而不仅限于Telnet服务,即使未勾选Telnet服务,即使未勾选Telnet服务也支持
func CheckPort(host string, port int) bool {
	address := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout("tcp", address, time.Second*2)
	if err != nil {
		return false
	}

	defer conn.Close()

	return true
}

// CheckPing windows系统下检查指定ip地址是否能ping通
func CheckPing(ip string) bool {
	args := []string{"-n", "1", "-w", "500", ip}

	cmd := exec.Command("ping", args...)
	err := cmd.Run()
	if err != nil {
		return false
	}

	return true
}
