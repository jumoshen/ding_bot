package ding_bot

import (
	"time"
)

var (
	DefaultError   = "未知错误"
	RequestUrl     = "https://oapi.dingtalk.com/robot/send"
	RequestTimeout = 5 * time.Second
)

type (
	Param func(config *DingConfig)
)

func WithSecret(secret string) Param {
	return func(dt *DingConfig) {
		dt.secret = secret
	}
}
