package files

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_Read(t *testing.T) {
	b, _ := ReadIcon("../../../logo.ico")
	t.Log(b)
	t.Log(string(b))
}

func Test_GetRunPath(t *testing.T) {
	cwd, _ := os.Getwd()
	t.Log(cwd)
	path, err := GetRunPath()
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(path)
}

func TestDownloadFile(t *testing.T) {
	handler := func(rw http.ResponseWriter, req *http.Request) {
		filename := filepath.Join("testdata", "download.txt")
		err := DownloadFile(filename, rw)
		if err != nil {
			t.Fatal(err)
		}
	}
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := "golang"
	if string(body) != expected {
		t.Fatalf("got file content [%s], want content [%s]", body, expected)
	}
	fields := strings.Split(resp.Header.Get("Content-Disposition"), "=")
	if len(fields) != 2 {
		t.Fatalf("unexpected download filename")
	}
	if strings.TrimSpace(fields[1]) != `"download.txt"` {
		t.Fatalf("unexpected download filename: %s", fields[1])
	}
}
