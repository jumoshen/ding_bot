package sender

import (
	"encoding/json"

	"github.com/jumoshen/ding_bot/param"
)

type (
	Sender struct {
		at      *At
		message param.DingMessage
	}

	At struct {
		AtMobiles []string `json:"atMobiles,omitempty"`
		IsAtAll   bool     `json:"isAtAll,omitempty"`
	}
)

func (r *Sender) GetMethod() string {
	return "POST"
}

func (r *Sender) GetHeader() map[string]string {
	header := map[string]string{
		"Content-type":  "application/json;charset=UTF-8",
		"Cache-Control": "no-cache",
		"Connection":    "Keep-Alive",
		"User-Agent":    "ding talk robot send",
	}
	return header
}

func (r *Sender) GetBody() ([]byte, error) {
	msg := make(map[string]interface{}, 3)
	msg["msgtype"] = r.message.GetType()
	if r.at != nil {
		msg["at"] = r.at
	}
	name := r.message.GetType()
	msg[name] = r.message
	return json.Marshal(msg)
}

func (r *Sender) GetSuccessCode() int64 {
	return 0
}

func NewSender(message param.DingMessage, options ...SendOption) *Sender {
	r := &Sender{message: message}
	for _, option := range options {
		option(r)
	}
	return r
}
