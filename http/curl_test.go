package http

import (
	"testing"
	"time"

	"github.com/valyala/fasthttp"
)

func Test_Get(t *testing.T) {
	newCurl := &FastHttp{
		Client: &fasthttp.Client{
			MaxIdleConnDuration: time.Minute,
			MaxConnDuration:     time.Minute,
			ReadTimeout:         time.Second * 10,
			WriteTimeout:        time.Second * 10,
		},
		TimeOut:     time.Second * 10,
		ContentType: []byte("application/json"),
		Headers:     nil,
	}
	rb, err := newCurl.Get("https://www.baidu.com")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(string(rb))
}
