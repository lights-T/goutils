package http

import (
	"errors"
	"time"

	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
)

type FastHttp struct {
	Client      *fasthttp.Client
	TimeOut     time.Duration
	ContentType []byte
	Headers     []*Headers
}

type Headers struct {
	Key   string
	Value string
}

func (c *FastHttp) Get(url string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(url)
	req.Header.SetContentTypeBytes(c.ContentType)
	req.Header.SetMethod(fasthttp.MethodGet)
	for _, header := range c.Headers {
		req.Header.Set(header.Key, header.Value)
	}
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)
	if err := c.Client.DoTimeout(req, res, c.TimeOut); err != nil {
		return nil, err
	}
	if res.StatusCode() != fasthttp.StatusOK {
		return nil, errors.New(fasthttp.StatusMessage(res.StatusCode()))
	}
	bb := bytebufferpool.Get()
	_, _ = bb.Write(res.Body())
	defer bytebufferpool.Put(bb)
	d := make([]byte, len(bb.Bytes()))
	copy(d, bb.Bytes())
	return d, nil
}

func (c *FastHttp) Post(url string, body []byte) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(url)
	req.Header.SetContentTypeBytes(c.ContentType)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetBody(body)
	for _, header := range c.Headers {
		req.Header.Set(header.Key, header.Value)
	}
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)
	if err := c.Client.DoTimeout(req, res, c.TimeOut); err != nil {
		return nil, err
	}
	if res.StatusCode() != fasthttp.StatusOK {
		return nil, errors.New(fasthttp.StatusMessage(res.StatusCode()))
	}
	bb := bytebufferpool.Get()
	_, _ = bb.Write(res.Body())
	defer bytebufferpool.Put(bb)
	d := make([]byte, len(bb.Bytes()))
	copy(d, bb.Bytes())
	return d, nil
}

//PostForHeader204 结果为204，获取header为结果
func (c *FastHttp) PostForHeader204(url string, body []byte) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(url)
	req.Header.SetContentTypeBytes(c.ContentType)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetBody(body)
	for _, header := range c.Headers {
		req.Header.Set(header.Key, header.Value)
	}
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)
	if err := c.Client.DoTimeout(req, res, c.TimeOut); err != nil {
		return nil, err
	}
	if res.StatusCode() != fasthttp.StatusNoContent {
		return nil, errors.New(fasthttp.StatusMessage(res.StatusCode()))
	}
	return res.Header.Header(), nil
}
