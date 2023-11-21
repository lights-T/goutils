package regexp

import "testing"

func TestIsAlpha(t *testing.T) {
	m, err := IsAlphaAndNum("aQsQA1234567890")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(m)
}
