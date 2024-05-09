package windows

import "testing"

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
