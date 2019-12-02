package message

import "errors"

//ErrInvalidReply 无效的回复
var ErrInvalidReply = errors.New("无效的回复消息")

//ErrUnsupportReply 不支持的回复类型
var ErrUnsupportReply = errors.New("不支持的回复消息")

// ReplyScene 返回场景
type ReplyScene string

const (
	// ReplySceneKefu 客服场景
	ReplySceneKefu ReplyScene = "kefu"
	// ReplySceneOpen 开放平台
	ReplySceneOpen = "open"
	// 支付
	ReplyScenePay = "pay"
)

//Reply 消息回复
type Reply struct {
	ReplyScene   ReplyScene
	ResponseType ResponseType
	MsgData      interface{}
}
