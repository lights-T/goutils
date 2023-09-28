package token

import (
	"testing"
	"time"
)

func TestToken_Encode(t *testing.T) {
	var err error
	DefaultToken, err = New("tGKN6AqvXCTYG3RrhdgLpTpYRgg6YBvT", time.Hour*7)
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
	DefaultToken, err = New("tGKN6AqvXCTYG3RrhdgLpTpYRgg6YBvT", time.Hour*7)
	if err != nil {
		t.Fatal(err)
	}
	s, err := DefaultToken.Decode("5YLKX35Hzvpu4w0YUcvBhTXzCTcKpUXu1muAPLFlaWy1pOfH4msXjc8OVgBU7U0cg")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}
