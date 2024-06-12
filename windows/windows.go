package windows

import (
	"syscall"

	"github.com/JamesHovious/w32"
	"github.com/lxn/walk"
	"gopkg.in/toast.v1"
)

func MessageBoxByErr(content string) {
	//w32.MessageBox(0, content, "Error", 0)
	walk.MsgBox(nil, "Error", content, walk.MsgBoxIconInformation)
}

func MessageBoxByInfo(content string) {
	//w32.MessageBox(0, content, "Error", 0)
	walk.MsgBox(nil, "Info", content, walk.MsgBoxIconInformation)
}

//GetSystemMetrics 获取windows分辨率，0宽，1高
func GetSystemMetrics(index int) int {
	return w32.GetSystemMetrics(index)
}

func SystemMetricsByOne(text string) int {
	return w32.MessageBox(0, text, "Tips", 1)
}

func Notify(appID, title, message string) {
	notification := toast.Notification{
		AppID:   appID,
		Title:   title,
		Message: message,
		//Icon:    "go.png", // This file must exist (remove this line if it doesn't)
		//Actions: []toast.Action{
		//	{"protocol", "I'm a button", ""},
		//	{"protocol", "Me too!", ""},
		//},
	}
	_ = notification.Push()
}

type HandleCMDRsp struct {
	GetConsoleWindows *syscall.LazyProc
	ShowWindowAsync   *syscall.LazyProc
	ConsoleHandle     uintptr
	R2                uintptr
}

func HandleCMD() (rsp *HandleCMDRsp, err error) {
	rsp = &HandleCMDRsp{}
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	user32 := syscall.NewLazyDLL("user32.dll")
	// https://docs.microsoft.com/en-us/windows/console/getconsolewindow
	rsp.GetConsoleWindows = kernel32.NewProc("GetConsoleWindow")
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showwindowasync
	rsp.ShowWindowAsync = user32.NewProc("ShowWindowAsync")
	rsp.ConsoleHandle, rsp.R2, err = rsp.GetConsoleWindows.Call()
	//使用返回值时，注意err是否为nil
	if rsp.ConsoleHandle == 0 {
		return rsp, err
	}
	return rsp, nil
}
