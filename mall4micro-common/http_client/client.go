package http_client

import (
	"bytes"
	"net/http"
)

type Client struct {
	Method      string            `json:"method"`
	Url         string            `json:"url"`
	ContentType string            `json:"content_type"`
	Headers     map[string]string `json:"headers"`
}

func NewHttpClient(method, url, contentType string, headers map[string]string) *Client {
	return &Client{
		Method:      method,
		Url:         url,
		ContentType: contentType,
		Headers:     headers,
	}
}

func (client *Client) Request(data []byte) (*http.Response, error) {
	reader := bytes.NewReader(data)
	request, err := http.NewRequest(client.Method, client.Url, reader)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = request.Body.Close()
	}() // 程序在使用完回复后必须关闭回复的主体
	if len(client.ContentType) != 0 {
		request.Header.Set("Content-Type", client.ContentType)
	} else {
		request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	}
	if len(client.Headers) != 0 {
		for key, value := range client.Headers {
			request.Header.Set(key, value)
		}
	}
	c := http.Client{}
	resp, err := c.Do(request) // Do 方法发送请求，返回 HTTP 回复
	if err != nil {
		return nil, err
	}
	return resp, nil
}
