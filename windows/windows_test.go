package windows

import (
	"testing"
)

func Test_Notify(t *testing.T) {
	Notify("Example App", "notification", "Some message about how important something is...")
}

func Test_HandleCMD(t *testing.T) {
	rsp, err := HandleCMD()
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("GetConsoleWindows: %+v", rsp.GetConsoleWindows)
	t.Logf("ShowWindowAsync: %+v", rsp.ShowWindowAsync)
	t.Logf("ConsoleHandle: %+v", rsp.ConsoleHandle)
	t.Logf("R2: %+v", rsp.R2)
}

func Test_PowerShell(t *testing.T) {
	ps, err := New()
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Run("创建startup快捷方式", func(t *testing.T) {
		//execPath := "C:\\Users\\E1394288\\Documents\\Product\\soft\\test\\goutils\\windows\\UMP.exe"
		execPath := "soft\\test\\goutils\\windows\\UMP.exe"
		if err = ps.EnableAutostartWin(execPath); err != nil {
			t.Fatal(err.Error())
		}
		t.Log("Shortcut created successfully")
	})
	t.Run("创建桌面快捷方式", func(t *testing.T) {
		execPath := "C:\\Users\\E1394288\\Documents\\Product\\soft\\test\\goutils\\windows\\UMP.exe"
		if err = ps.EnableDesktopWin(execPath); err != nil {
			t.Fatal(err.Error())
		}
		t.Log("Shortcut created successfully")
	})
}
