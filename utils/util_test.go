package utils

import "testing"

func Test_ScanPort(t *testing.T) {
	b := ScanPort("tcp", "127.0.0.1", 8050)
	t.Log(b)
}
