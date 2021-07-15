package param

import (
	"time"

	"github.com/jumoshen/ding_bot"
)

var (
	DefaultError   = "未知错误"
	RequestUrl     = "https://oapi.dingtalk.com/robot/send"
	RequestTimeout = 5 * time.Second
)

type (
	Param func(config *ding_bot.DingConfig)
)

func WithSecret(secret string) Param {
	return func(dt *ding_bot.DingConfig) {
		dt.secret = secret
	}
}
