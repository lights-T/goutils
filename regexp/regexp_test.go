package regexp

import "testing"

func TestIsAlpha(t *testing.T) {
	m, err := IsAlpha("aQsQW啊哈哈A")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(m)
}
