package goutils

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"testing"

	"github.com/jinzhu/now"
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

func TestMinToDHM(t *testing.T) {
	d, h, m := MinToDHM(86400)
	t.Log(d, h, m)
}

func TestSecToDHM(t *testing.T) {
	t.Log(SecToDHM(59))
}

func TestGetDurationToMin(t *testing.T) {
	t.Log(GetDurationToMin("2025-06-05 07:53:31", "2025-06-05 08:10:31"))
}

func Test(t *testing.T) {
	a, _ := now.Parse("2024-02-19 7:30:00")
	aa, _ := now.Parse("2024-02-19 10:32")
	aaa := aa.Unix() - a.Unix()
	t.Log(aaa) //112

	b, _ := now.Parse("2024-02-19 10:32:50")
	bb, _ := now.Parse("2024-02-19 10:43:51")
	bbb := bb.Unix() - b.Unix()
	t.Log(bbb) //12

	c, _ := now.Parse("2024-02-19 10:53:51")
	cc, _ := now.Parse("2024-02-19 11:24:19")
	ccc := cc.Unix() - c.Unix()
	t.Log(ccc) //31

	d, _ := now.Parse("2024-02-19 11:34:19")
	dd, _ := now.Parse("2024-02-19 13:58:32")
	ddd := dd.Unix() - d.Unix()
	t.Log(ddd)

	e, _ := now.Parse("2024-02-19 14:08:32")
	ee, _ := now.Parse("2024-02-19 19:45:40")
	eee := ee.Unix() - e.Unix()
	t.Log(eee)

	f, _ := now.Parse("2024-02-19 19:45:40")
	ff, _ := now.Parse("2024-02-20 7:30:00")
	fff := ff.Unix() - f.Unix() //755
	t.Log(fff)                  //755

	total := float64(aaa + ddd + eee + fff)
	t.Log(total)
	t.Log(total / float64(86400))
}
