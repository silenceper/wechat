package message

// MsgType 基本消息类型
type MsgType string

// EventType 事件类型
type EventType string

// InfoType 第三方平台授权事件类型
type InfoType string

const (
	//MsgTypeText 文本消息
	MsgTypeText MsgType = "text"
	//MsgTypeImage 图片消息
	MsgTypeImage = "image"
	//MsgTypeLink 图文链接
	MsgTypeLink = "link"
	//MsgTypeMiniProgramPage 小程序卡片
	MsgTypeMiniProgramPage = "miniprogrampage"
)
