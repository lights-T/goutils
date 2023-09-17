package computer

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unsafe"

	ustrings "github.com/lights-T/goutils/utils/strings"

	"github.com/StackExchange/wmi"
	"golang.org/x/sys/windows"
)

//GetPhysicalID 获取电脑id
func GetPhysicalID() (string, error) {
	var id string
	var ids []string
	if guid, err := GetMachineGuid(); err != nil {
		return id, err
	} else {
		fmt.Printf("guid: %s\n", guid)
		ids = append(ids, guid)
	}
	if cpuInfo, err := GetCPUInfo(); err != nil {
		return id, err
	} else {
		fmt.Printf("cpuInfo.PhysicalID: %s\n", cpuInfo[0].VendorID+cpuInfo[0].PhysicalID)
		ids = append(ids, cpuInfo[0].VendorID+cpuInfo[0].PhysicalID)
	}
	if mac, err := GetMACAddress(); err != nil {
		return id, err
	} else {
		fmt.Printf("mac: %s\n", mac)
		ids = append(ids, mac)
	}
	sort.Strings(ids)
	idsStr := strings.Join(ids, "|/|")
	return ustrings.GetMd5String(idsStr, true, true), nil
}

// GetMachineGuid 开始+r: regedit，打开注册表后find：MachineGuid
func GetMachineGuid() (string, error) {
	// there has been reports of issues on 32bit using golang.org/x/sys/windows/registry, see https://github.com/shirou/gopsutil/pull/312#issuecomment-277422612
	// for rationale of using windows.RegOpenKeyEx/RegQueryValueEx instead of registry.OpenKey/GetStringValue
	var h windows.Handle
	err := windows.RegOpenKeyEx(windows.HKEY_LOCAL_MACHINE, windows.StringToUTF16Ptr(`SOFTWARE\Microsoft\Cryptography`), 0, windows.KEY_READ|windows.KEY_WOW64_64KEY, &h)
	if err != nil {
		return "", err
	}
	defer windows.RegCloseKey(h)

	const windowsRegBufLen = 74 // len(`{`) + len(`abcdefgh-1234-456789012-123345456671` * 2) + len(`}`) // 2 == bytes/UTF16
	const uuidLen = 36

	var regBuf [windowsRegBufLen]uint16
	bufLen := uint32(windowsRegBufLen)
	var valType uint32
	err = windows.RegQueryValueEx(h, windows.StringToUTF16Ptr(`MachineGuid`), nil, &valType, (*byte)(unsafe.Pointer(&regBuf[0])), &bufLen)
	if err != nil {
		return "", err
	}

	hostID := windows.UTF16ToString(regBuf[:])
	hostIDLen := len(hostID)
	if hostIDLen != uuidLen {
		return "", fmt.Errorf("HostID incorrect: %q\n", hostID)
	}

	return hostID, nil
}

type cpuInfo struct {
	CPU        int32  `json:"cpu"`
	VendorID   string `json:"vendorId"`
	PhysicalID string `json:"physicalId"`
}

type win32Processor struct {
	Manufacturer string
	ProcessorID  *string
}

func GetCPUInfo() ([]cpuInfo, error) {
	var ret []cpuInfo
	var dst []win32Processor
	q := wmi.CreateQuery(&dst, "")
	fmt.Println(q)
	if err := WmiQuery(q, &dst); err != nil {
		return ret, err
	}

	var procID string
	for i, l := range dst {
		procID = ""
		if l.ProcessorID != nil {
			procID = *l.ProcessorID
		}

		cpu := cpuInfo{
			CPU:        int32(i),
			VendorID:   l.Manufacturer,
			PhysicalID: procID,
		}
		ret = append(ret, cpu)
	}

	return ret, nil
}

//WmiQuery WithContext - wraps wmi.Query with a timed-out context to avoid hanging
func WmiQuery(query string, dst interface{}, connectServerArgs ...interface{}) error {
	ctx := context.Background()
	if _, ok := ctx.Deadline(); !ok {
		ctxTimeout, cancel := context.WithTimeout(ctx, 3000000000) //超时时间3s
		defer cancel()
		ctx = ctxTimeout
	}

	errChan := make(chan error, 1)
	go func() {
		errChan <- wmi.Query(query, dst, connectServerArgs...)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}

//GetMACAddress 获取电脑MAC地址
func GetMACAddress() (string, error) {
	macAddress := ""
	var err error
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return macAddress, err
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags&net.FlagUp) != 0 && (netInterfaces[i].Flags&net.FlagLoopback) == 0 {
			address, _ := netInterfaces[i].Addrs()
			for _, address := range address {
				ipNet, ok := address.(*net.IPNet)
				if ok && ipNet.IP.IsGlobalUnicast() {
					// 如果IP是全局单拨地址，则返回MAC地址
					macAddress = netInterfaces[i].HardwareAddr.String()
					return macAddress, err
				}
			}
		}
	}
	return macAddress, err
}

//PortInUse 传入查询的端口号
// 返回端口号对应的进程PID，若没有找到相关进程，返回-1
func PortInUse(portNumber int) int {
	res := -1
	var outBytes bytes.Buffer
	cmdStr := fmt.Sprintf("netstat -ano -p tcp | findstr %d", portNumber)
	cmd := exec.Command("cmd", "/c", cmdStr)
	cmd.Stdout = &outBytes
	cmd.Run()
	resStr := outBytes.String()
	r := regexp.MustCompile(`\s\d+\s`).FindAllString(resStr, -1)
	if len(r) > 0 {
		pid, err := strconv.Atoi(strings.TrimSpace(r[0]))
		if err != nil {
			res = -1
		} else {
			res = pid
		}
	}
	return res
}

func ListProcess(text string) (string, error) {
	cmd := exec.Command("tasklist", "/fi", text)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func ListProcessByPid(pid int) (string, error) {
	cmd := exec.Command("tasklist", "/v", "/fi", fmt.Sprintf("pid eq %d", pid))
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func ListAllProcess(text string) (string, error) {
	cmd := exec.Command("tasklist")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func KillProcess(text string) (string, error) {
	cmd := exec.Command("taskkill", "/f", "/t", "/fi", text)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
