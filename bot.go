package ding_bot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

type (
	Requester interface {
		GetMethod() string
		GetHeader() map[string]string
		GetBody() ([]byte, error)
		GetSuccessCode() int64
	}

	DingConfig struct {
		accessToken string
		mu          sync.Mutex
		url         string
		secret      string
		timeout     time.Duration
		client      *HttpClient
		response    *http.Response
		err         error
	}
)

func New(url string, options ...Param) *DingConfig {
	dc := &DingConfig{
		url:     url,
		timeout: 5 * time.Second,
	}
	for _, option := range options {
		option(dc)
	}
	dc.initClient()
	return dc
}

func (dc *DingConfig) initClient() {
	step := "?"
	if strings.Contains(dc.url, "?") {
		step = "&"
	}
	params := dc.genQueryParams()
	dc.url = strings.Join([]string{dc.url, params}, step)
	dc.client = NewHttpClient(dc.url, dc.timeout)
}

func (dc *DingConfig) genQueryParams() string {
	params := url.Values{}
	if dc.secret != "" {
		timestamp := time.Now().UnixNano() / 1e6
		sign := genSign(dc.secret, timestamp)
		params.Add("timestamp", strconv.FormatInt(timestamp, 10))
		params.Add("sign", sign)
	}
	return params.Encode()
}

func (dc *DingConfig) checkURL() error {
	_, err := url.Parse(dc.url)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DingConfig) Request(req Requester) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	if err := dc.checkURL(); err != nil {
		return dc.newError(ErrorCheckUrl, err)
	}
	body, err := dc.doRequest(req)
	if err != nil {
		return dc.newError(ErrorDoRequest, err, body)
	}
	if err := dc.checkResponse(req); err != nil {
		return dc.newError(ErrorResponse, err, body)
	}
	return nil
}

func (dc *DingConfig) doRequest(req Requester) (string, error) {
	method := req.GetMethod()
	header := req.GetHeader()
	body, err := req.GetBody()

	if err != nil {
		return "", err
	}
	log.Printf("url: %s, timeout: %s, method: %s, header: %v, body: %s",
		dc.url, dc.timeout, method, header, body)

	dc.response, err = dc.client.Request(method, header, body)
	if err != nil {
		return string(body), err
	}
	return string(body), nil
}

func (dc *DingConfig) checkResponse(req Requester) error {
	defer dc.response.Body.Close()
	data, err := ioutil.ReadAll(dc.response.Body)
	if err != nil {
		return err
	}
	dc.response.Body = ioutil.NopCloser(bytes.NewReader(data))

	if dc.response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid http status %d, body: %s", dc.response.StatusCode, data)
	}

	respMsg := ResponseMsg{}
	if err := json.Unmarshal(data, &respMsg); err != nil {
		return fmt.Errorf("body: %s, %w", data, err)
	}
	respMsg.ApplicationHost = dc.response.Header.Get("Application-Host")
	respMsg.ServiceHost = dc.response.Header.Get("Location-Host")
	if respMsg.ErrCode != req.GetSuccessCode() {
		return fmt.Errorf("%s", respMsg)
	}
	return nil
}

func genSign(secret string, timestamp int64) string {
	b := &[]byte{}
	*b = append(*b, strconv.FormatInt(timestamp, 10)...)
	*b = append(*b, '\n')
	*b = append(*b, secret...)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(*b)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
