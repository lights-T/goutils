package strings

import (
	"testing"

	lfiles "github.com/lights-T/goutil/files"
)

func Test_BytesToStringByAscii(t *testing.T) {
	b, _ := lfiles.ReadIcon("../../../../logo.ico")
	t.Log(b)
	str := BytesToStringByAscii(b)
	t.Log(str)
}
