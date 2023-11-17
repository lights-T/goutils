package regexp

import "testing"

func TestIsAlpha(t *testing.T) {
	m, err := IsAlpha2("aQsQA.")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(m)
}
