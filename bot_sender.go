package ding_bot

import (
	"github.com/jumoshen/ding_bot/sender"
)

func (dc *DingConfig) RobotSendText(text string, options ...sender.SendOption) error {
	msg := sender.Text{Content: text}
	return dc.Request(sender.NewSender(msg, options...))
}

// RobotSendLink link类型的消息
func (dc *DingConfig) RobotSendLink(title, text, messageURL, picURL string, options ...sender.SendOption) error {
	msg := sender.Link{
		Title:      title,
		Text:       text,
		MessageURL: messageURL,
		PicURL:     picURL,
	}
	return dc.Request(sender.NewSender(msg, options...))
}

// RobotSendMarkdown markdown类型的消息
func (dc *DingConfig) RobotSendMarkdown(title, text string, options ...sender.SendOption) error {
	msg := sender.Markdown{
		Title: title,
		Text:  text,
	}
	return dc.Request(sender.NewSender(msg, options...))
}
