package windows

import (
	"github.com/lxn/walk"
)

func MessageBoxByErr(content string) {
	//w32.MessageBox(0, content, "Error", 0)
	walk.MsgBox(nil, "Error", content, walk.MsgBoxIconInformation)
}
