package message

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/util"
)

const (
	kfMsgSendURL = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
)


//Text 文本消息
type EventMsg struct {
	CommonToken
	Content string `xml:"Content"`
}

//NewText 初始化文本消息
func NewEventMsg(content string) *EventMsg {
	eventmsg := new(EventMsg)
	eventmsg.Content = content
	return eventmsg
}


//客服消息
type Kfcustomer struct {
	*context.Context
}

//NewKfcustomer 实例化
func NewKfcustomer(context *context.Context) *Kfcustomer {
	kf := new(Kfcustomer)
	kf.Context = context
	return kf
}

//{
//	"touser": "OPENID"
//	"msgtype": "msgmenu",
//	"msgmenu": {
//		"head_content": "您对本次服务是否满意呢? "
//		"list": [
//			{
//				"id": "101",
//				"content": "满意"
//			},
//			{
//			"id": "102",
//			"content": "不满意"
//			}
//		],
//		"tail_content": "欢迎再次光临"
//	}
//}
//msgmenu 客服menu消息
type KfMsg struct {
	Touser string	`json:"touser"`
	Msgtype	string	`json:"msgtype"`
	Msgmenu struct{
		Headcontent string `json:"head_content"`
		List	[]SingleList	`json:"list"`
		Tailcontent	string	`json:"tail_content"`
	}	`json:"msgmenu"`
}
type SingleList struct {
	Id string	`json:"id"`
	Content	string	`json:"content"`
}

type resKfSend struct {
	util.CommonError
}


//发送图片消息
//{
//	"touser":"OPENID",
//	"msgtype":"image",
//	"image":
//	{
//		"media_id":"MEDIA_ID"
//	}
//}
type KfImgMsg struct {
	Touser string	`json:"touser"`
	Msgtype	string	`json:"msgtype"`
	Image struct{
		Mediaid string `json:"media_id"`
	}	`json:"image"`
}



//NewNews 初始化客服menu消息
//func NewKfMsg(singlelist []SingleList,headcontent string,tailcontent string,touser string,msgtype string) *KfMsg {
//	kfmsg := new(KfMsg)
//	kfmsg.Msgmenu.List = singlelist
//	kfmsg.Msgmenu.Headcontent = headcontent
//	kfmsg.Msgmenu.Tailcontent = tailcontent
//	kfmsg.Touser = touser
//	kfmsg.Msgtype = msgtype
//	return kfmsg
//}


//Send 发送客服菜单模板消息
func (kf *Kfcustomer) KfSend(msg *KfMsg) (err error) {
	var accessToken string
	accessToken, err = kf.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", kfMsgSendURL, accessToken)
	response, err := util.PostJSON(uri, msg)
	var result resKfSend
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("kf msg send error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

//Send 发送客服图片消息
func (kf *Kfcustomer) KfImgSend(msg *KfImgMsg) (err error) {
	var accessToken string
	accessToken, err = kf.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", kfMsgSendURL, accessToken)
	response, err := util.PostJSON(uri, msg)
	var result resKfSend
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("kf msg send error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}


