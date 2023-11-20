package net

import "testing"

func Test_port(t *testing.T) {
	if res := CheckPort("192.168.56.211", 22); !res {
		t.Fatalf("不通")
	}
	t.Fatal("Done.")
}

func Test_ip(t *testing.T) {
	if res := CheckPing("192.168.56.211"); !res {
		t.Fatalf("不通")
	}
	t.Fatal("Done.")
}
