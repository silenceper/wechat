package message

import "encoding/xml"

// MsgType 基本消息类型
type MsgType string

// EventType 事件类型
type EventType string

// InfoType 第三方平台授权事件类型
type InfoType string

const (
	// MsgTypeText 文本消息
	MsgTypeText MsgType = "text"
	// MsgTypeImage 图片消息
	MsgTypeImage = "image"
	// MsgTypeLink 图文链接
	MsgTypeLink = "link"
	// MsgTypeMiniProgramPage 小程序卡片
	MsgTypeMiniProgramPage = "miniprogrampage"
	// MsgTypeEvent 事件
	MsgTypeEvent MsgType = "event"
	// DataTypeXML XML格式数据
	DataTypeXML = "xml"
	// DataTypeJSON JSON格式数据
	DataTypeJSON = "json"
)

const (
	//EventSubscribePopup 用户操作订阅通知弹窗事件推送，用户在图文等场景内订阅通知的操作
	EventSubscribePopup EventType = "subscribe_msg_popup_event"
	//EventSubscribeChange 用户管理订阅通知，用户在服务通知管理页面做通知管理时的操作
	EventSubscribeChange = "subscribe_msg_change_event"
	//EventSubscribeSent 发送订阅通知，调用 bizsend 接口发送通知
	EventSubscribeSent = "subscribe_msg_sent_event"
)

const (
	InfoTypeAcceptSubscribeMessage InfoType = "accept"
	InfoTypeRejectSubscribeMessage          = "reject"
)

// CommonToken 消息中通用的结构
type CommonToken struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      MsgType  `xml:"MsgType"`
}

// MiniProgramMixMessage 小程序回调的消息结构
type MiniProgramMixMessage struct {
	CommonToken

	MsgID int64 `xml:"MsgId"`

	// 文本消息
	Content string `xml:"Content"`

	// 图片消息
	PicURL  string `xml:"PicUrl"`
	MediaID string `xml:"MediaId"`

	// 小程序卡片消息
	Title        string `xml:"Title"`
	AppID        string `xml:"AppId"`
	PagePath     string `xml:"PagePath"`
	ThumbURL     string `xml:"ThumbUrl"`
	ThumbMediaID string `xml:"ThumbMediaId"`

	// 进入会话事件
	Event       string `xml:"Event"`
	SessionFrom string `xml:"SessionFrom"`

	// 用户操作订阅通知弹窗消息回调
	SubscribeMsgPopupEvent struct {
		List []SubscribeMessageList `xml:"List"`
	} `xml:"SubscribeMsgPopupEvent"`

	// 用户管理订阅通知回调
	SubscribeMsgChangeEvent struct {
		List []SubscribeMessageList `xml:"List"`
	} `xml:"SubscribeMsgChangeEvent"`

	// 用户发送订阅通知回调
	SubscribeMsgSentEvent struct {
		List []SubscribeMessageList `xml:"List"`
	} `xml:"SubscribeMsgSentEvent"`
}

// SubscribeMessageList 订阅消息事件列表
type SubscribeMessageList struct {
	TemplateId            string `xml:"TemplateId"`
	SubscribeStatusString string `xml:"SubscribeStatusString"`
	PopupScene            string `xml:"PopupScene"`
}

// EncryptedXMLMsg 安全模式下的消息体
type EncryptedXMLMsg struct {
	XMLName      struct{} `xml:"xml" json:"-"`
	ToUserName   string   `xml:"ToUserName" json:"toUserName"`
	EncryptedMsg string   `xml:"Encrypt" json:"Encrypt"`
}

// CDATA 使用这种类型，在序列化 xml 时文本会被解析器忽略
type CDATA string
