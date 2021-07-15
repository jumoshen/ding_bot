package ding_bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jumoshen/ding_bot/client"
	"github.com/jumoshen/ding_bot/param"
	"github.com/jumoshen/ding_bot/utils"
)

type (
	DingConfig struct {
		accessToken string
		mu          sync.Mutex
		url         string
		secret      string
		timeout     time.Duration
		client      *client.HttpClient
		response    *http.Response
		err         error
	}
)

func New(accessToken string, options ...param.Param) *DingConfig {
	dc := &DingConfig{
		accessToken: accessToken,
		url:         fmt.Sprintf("%s?access_token=%s", param.RequestUrl, accessToken),
		timeout:     param.RequestTimeout,
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
	dc.client = client.NewHttpClient(dc.url, dc.timeout)
}

func (dc *DingConfig) genQueryParams() string {
	params := url.Values{}
	if dc.secret != "" {
		timestamp := time.Now().UnixNano() / 1e6
		sign := utils.GenSign(dc.secret, timestamp)
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

func (dc *DingConfig) Request(req client.Requester) error {
	dc.mu.Lock()
	defer dc.mu.Unlock()
	if err := dc.checkURL(); err != nil {
		return param.NewError(param.ErrorCheckUrl, err)
	}
	body, err := dc.doRequest(req)
	if err != nil {
		return param.NewError(param.ErrorDoRequest, err, body)
	}
	if err := dc.checkResponse(req); err != nil {
		return param.NewError(param.ErrorResponse, err, body)
	}
	return nil
}

func (dc *DingConfig) doRequest(req client.Requester) (string, error) {
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

func (dc *DingConfig) checkResponse(req client.Requester) error {
	var err error
	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(dc.response.Body)

	data, err := ioutil.ReadAll(dc.response.Body)
	if err != nil {
		return err
	}
	dc.response.Body = ioutil.NopCloser(bytes.NewReader(data))

	if dc.response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid client status %d, body: %s", dc.response.StatusCode, data)
	}

	respMsg := client.ResponseMsg{}
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
