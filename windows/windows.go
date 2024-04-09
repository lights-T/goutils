package windows

import (
	"github.com/lxn/walk"
	"gopkg.in/toast.v1"
)

func MessageBoxByErr(content string) {
	//w32.MessageBox(0, content, "Error", 0)
	walk.MsgBox(nil, "Error", content, walk.MsgBoxIconInformation)
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
