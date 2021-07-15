package sender

import (
	"github.com/jumoshen/ding_bot"
)

func (dc *ding_bot.DingConfig) RobotSendText(text string, options ...SendOption) error {
	msg := Text{Content: text}
	return dc.Request(NewSender(msg, options...))
}

// RobotSendLink link类型的消息
func (dc *ding_bot.DingConfig) RobotSendLink(title, text, messageURL, picURL string, options ...SendOption) error {
	msg := Link{
		Title:      title,
		Text:       text,
		MessageURL: messageURL,
		PicURL:     picURL,
	}
	return dc.Request(NewSender(msg, options...))
}

// RobotSendMarkdown markdown类型的消息
func (dc *ding_bot.DingConfig) RobotSendMarkdown(title, text string, options ...SendOption) error {
	msg := Markdown{
		Title: title,
		Text:  text,
	}
	return dc.Request(NewSender(msg, options...))
}
