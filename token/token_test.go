package token

import (
	"testing"
	"time"
)

func TestToken_Encode(t *testing.T) {
	var err error
	DefaultToken, err = New("tG1N6AqvXCTYG3Rrh5gLpTpYRgg6Y0vT", time.Hour*7)
	if err != nil {
		t.Fatal(err)
	}
	s, err := DefaultToken.Encode(175195)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}

func TestToken_Decode(t *testing.T) {
	var err error
	DefaultToken, err = New("tG1N6AqvXCTYG3Rrh5gLpTpYRgg6Y0vT", time.Hour*7)
	if err != nil {
		t.Fatal(err)
	}
	s, err := DefaultToken.Decode("6J0gSdw2bErtOfRwVcTd5jz4MsoQYAoHuZi4lg43cXyhstGwFVBFLa0l2Gz0nFiiUuxIH")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}
