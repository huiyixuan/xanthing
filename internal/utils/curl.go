package utils

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

type Request struct {
	URL     string
	Method  string
	Headers map[string]string
	Body    []byte
	Timeout time.Duration
}

type CurlResponse struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

func NewRequest() *Request {
	return &Request{
		Method:  "GET",
		Headers: make(map[string]string),
		Timeout: 30 * time.Second,
	}
}

func (r *Request) Send() (*CurlResponse, error) {
	req, err := http.NewRequest(r.Method, r.URL, bytes.NewBuffer(r.Body))
	if err != nil {
		return nil, err
	}

	// 设置请求头
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}

	// 创建HTTP客户端
	client := &http.Client{
		Timeout: r.Timeout,
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &CurlResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       body,
	}, nil
}