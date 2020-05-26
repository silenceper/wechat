package broadcast

import (
	"github.com/silenceper/wechat/v2/officialaccount/context"
	"github.com/silenceper/wechat/v2/util"
)

//MsgType 发送消息类型
type MsgType string

const (
	//图文消息
	MsgTypeNews MsgType = "mpnews"
	//文本
	MsgTypeText MsgType = "text"
	//语音/音频
	MsgTypeVoice MsgType = "voice"
	//图片
	MsgTypeImage MsgType = "mpvideo"
	//视频
	MsgTypeVideo MsgType = "mpvideo"
	//卡券
	MsgTypeWxCard MsgType = "wxcard"
)

//Broadcast 群发消息
type Broadcast struct {
	*context.Context
}

//NewBroadcast new
func NewBroadcast(ctx *context.Context) *Broadcast {
	return &Broadcast{ctx}
}

//User 发送的用户
type User struct {
	TagID  int64
	OpenID []string
}

//Result 群发返回结果
type Result struct {
	util.CommonError
	MsgID     int64 `json:"msg_id"`
	MsgDataID int64 `json:"msg_data_id"`
}

//SendText 群发文本
//user 为nil，表示全员发送
//&User{TagID:2} 根据tag发送
//&User{OpenID:[]string("xxx","xxx")} 根据openid发送
func (broadcast *Broadcast) SendText(user *User, content string) (*Result, error) {
	//TODO
	return nil, nil
}

//TODO 群发其他消息
