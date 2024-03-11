package time

import (
	"github.com/jinzhu/now"
	"testing"
	"time"
)

func Test_GetDiffDays(t *testing.T) {
	t1, err := now.Parse("2024-03-12 15:12:12")
	if err != nil {
		t.Fatal(err.Error())
	}
	t2 := time.Now()
	d := GetDiffDays(t1, t2)
	t.Fatal(d)
}
