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
func TestIsDate(t *testing.T) {
	m, err := IsDate("2020-13-01")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(m)
}

func TestIsHHMM(t *testing.T) {
	m, err := IsHHMM("07:12")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(m)
}

func TestIsHHMMSS(t *testing.T) {
	m, err := IsHHMMSS("23:12:19")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(m)
}

func TestIsWeekly(t *testing.T) {
	m, err := IsWeekly("0")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(m)
}

func TestIsDateTime(t *testing.T) {
	//m, err := IsDateTime("2023-01-01")
	m, err := IsDateTime("2023-01-01 00:00:00")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(m)
}
