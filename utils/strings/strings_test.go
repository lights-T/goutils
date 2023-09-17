package strings

import (
	"testing"

	lfiles "client-side/server/lib/files"
)

func Test_BytesToStringByAscii(t *testing.T) {
	b, _ := lfiles.ReadIcon("../../../../logo.ico")
	t.Log(b)
	str := BytesToStringByAscii(b)
	t.Log(str)
}
