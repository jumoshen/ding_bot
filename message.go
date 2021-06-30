package ding_bot

const (
	MessageText messageType = iota
	MessageLink
	MessageMarkdown
	MessageActionCard
	MessageFeedCard
)

type DingMessage interface {
	GetType() string
}

type messageType uint8

func (r messageType) String() string {
	switch r {
	case MessageText:
		return "text"
	case MessageLink:
		return "link"
	case MessageMarkdown:
		return "markdown"
	case MessageActionCard:
		return "actionCard"
	case MessageFeedCard:
		return "feedCard"
	default:
		return "Unknown"
	}
}
