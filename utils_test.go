package goutils

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"testing"
)

func Test_ScanPort(t *testing.T) {
	b := ScanPort("tcp", "127.0.0.1", 8050)
	t.Log(b)
}

func TestRandNumber(t *testing.T) {
	f := func(min, max int) {
		for i := min; i <= max; i++ {
			num := RandNumber(i, max)
			if min <= num && num <= max {
				continue
			}
			t.Errorf("got %d, want range %d - %d", num, min, max)
		}
	}
	f(0, 1000)
	f(0, 0)
	f(-1000, 0)
	f(-1000, 1000)
}

func BenchmarkRandNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandNumber(1, math.MaxInt32)
	}
}

func ExampleRandNumber() {
	num := RandNumber(1, 1000)
	fmt.Println(num)
}

func TestPanicToError(t *testing.T) {
	err := PanicToError(func() {
		panic("error")
	})
	if err == nil {
		t.Errorf("got err is nil, want err is not nil")
	}
}

func ExamplePanicToError() {
	err := PanicToError(func() {
		panic("error")
	})
	fmt.Println(err)
}

func TestWorkDir(t *testing.T) {
	wd, err := WorkDirByExecutable()
	if err != nil {
		t.Fatal(err)
	}
	if wd != filepath.Dir(os.Args[0]) {
		t.Fatalf("got working dir [%s], want working dir [%s]", wd, os.Args[0])
	}
}

func TestWaitGroupWrapper_Wrap(t *testing.T) {
	wg := WaitGroupWrapper{}
	n := 10
	exited := make(chan struct{}, n)
	f := func() {
		exited <- struct{}{}
	}
	for i := 0; i < n; i++ {
		wg.Wrap(f)
	}
	wg.Wait()
	for i := 0; i < n; i++ {
		<-exited
	}
}
