package regexp

import "testing"

func TestIsAlpha(t *testing.T) {
	m, err := IsAlphaAndNum("aQsQA1234567890")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(m)
}

func TestIsNum(t *testing.T) {
	m, err := IsNum("123")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(m)
}
