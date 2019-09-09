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


//EventMsg 文本消息
type EventMsg struct {
	CommonToken
	Content string `xml:"Content"`
}

//NewEventMsg 初始化文本消息
func NewEventMsg(content string) *EventMsg {
	eventmsg := new(EventMsg)
	eventmsg.Content = content
	return eventmsg
}


//Kfcustomer 客服消息
type Kfcustomer struct {
	*context.Context
}

//NewKfcustomer 实例化
func NewKfcustomer(context *context.Context) *Kfcustomer {
	kf := new(Kfcustomer)
	kf.Context = context
	return kf
}

//KfMsg 客服menu消息
type KfMsg struct {
	Touser string	`json:"touser"`
	Msgtype	string	`json:"msgtype"`
	Msgmenu struct{
		Headcontent string `json:"head_content"`
		List	[]SingleList	`json:"list"`
		Tailcontent	string	`json:"tail_content"`
	}	`json:"msgmenu"`
}
//SingleList 菜单消息
type SingleList struct {
	ID string	`json:"id"`
	Content	string	`json:"content"`
}
//resKfSend kf消息发送
type resKfSend struct {
	util.CommonError
}

//KfImgMsg kf图片消息
type KfImgMsg struct {
	Touser string	`json:"touser"`
	Msgtype	string	`json:"msgtype"`
	Image struct{
		Mediaid string `json:"media_id"`
	}	`json:"image"`
}



//KfSend 发送客服菜单模板消息
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

//KfImgSend 发送客服图片消息
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


