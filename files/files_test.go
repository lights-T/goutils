package files

import (
	"os"
	"testing"
)

func Test_Read(t *testing.T) {
	b, _ := ReadIcon("../../../logo.ico")
	t.Log(b)
	t.Log(string(b))
}

func Test_GetRunPath(t *testing.T) {
	cwd, _ := os.Getwd()
	t.Log(cwd)
	path, err := GetRunPath()
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(path)
}
