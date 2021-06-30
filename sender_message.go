package ding_bot

// Text text类型
type Text struct {
	// Content 消息内容
	Content string `json:"content"`
}

func (Text) GetType() string {
	return MessageText.String()
}

// Link link类型
type Link struct {
	// Title 消息标题
	Title string `json:"title"`

	// Text 消息内容。如果太长只会部分展示
	Text string `json:"text,omitempty"`

	// MessageURL 点击消息跳转的URL
	MessageURL string `json:"messageUrl"`

	// PicURL 图片URL
	PicURL string `json:"picUrl,omitempty"`
}

func (Link) GetType() string {
	return MessageLink.String()
}

// Markdown markdown类型
type Markdown struct {
	// Title 首屏会话透出的展示内容
	Title string `json:"title"`

	// Text  markdown格式的消息
	Text string `json:"text"`
}

func (Markdown) GetType() string {
	return MessageMarkdown.String()
}

// Btn 按钮的信息
type Btn struct {
	// Title 按钮方案
	Title string `json:"title"`

	// ActionURL 点击按钮触发的URL
	ActionURL string `json:"actionURL"`
}

// ActionCard actionCard类型
type ActionCard struct {
	// Title 首屏会话透出的展示内容
	Title string `json:"title"`

	// Text markdown格式的消息
	Text string `json:"text"`

	// SingleTitle 单个按钮的方案。(设置此项和singleURL后btns无效)
	SingleTitle string `json:"singleTitle"`

	// SingleURL 点击singleTitle按钮触发的URL
	SingleURL string `json:"singleURL"`

	// HideAvatar      0-正常发消息者头像,1-隐藏发消息者头像
	HideAvatar string `json:"hideAvatar,omitempty"`

	// Btns 按钮的信息：title-按钮方案，actionURL-点击按钮触发的URL
	Btns []Btn `json:"btns,omitempty"`

	// BtnOrientation   0-按钮竖直排列，1-按钮横向排列
	BtnOrientation string `json:"btnOrientation,omitempty"`
}

func (ActionCard) GetType() string {
	return MessageActionCard.String()
}

type FeedCardLink struct {
	//Title 单条信息文本
	Title string `json:"title"`

	// MessageURL 点击单条信息到跳转链接
	MessageURL string `json:"messageURL"`

	// PicURL 单条信息后面图片的URL
	PicURL string `json:"picURL"`
}

// FeedCard 资讯类信息
type FeedCard struct {
	Links []FeedCardLink `json:"links"`
}

func (FeedCard) GetType() string {
	return MessageFeedCard.String()
}
