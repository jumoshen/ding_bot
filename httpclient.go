package ding_bot

import (
	"bytes"
	"net/http"
	"time"
)

type (
	Requester interface {
		GetMethod() string
		GetHeader() map[string]string
		GetBody() ([]byte, error)
		GetSuccessCode() int64
	}

	HttpClient struct {
		client *http.Client
		Url    string
	}
)

func NewHttpClient(url string, timeout time.Duration) *HttpClient {
	return &HttpClient{
		client: &http.Client{Timeout: timeout},
		Url:    url,
	}
}

func (hc *HttpClient) Request(method string, header map[string]string, body []byte) (response *http.Response, err error) {
	req, err := http.NewRequest(method, hc.Url, bytes.NewReader(body))
	if err != nil {
		return
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	return hc.client.Do(req)
}
