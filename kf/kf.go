package kf

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/message"
	"github.com/silenceper/wechat/util"
)

const (
	KfOperURL = "https://api.weixin.qq.com/customservice/kfaccount"
	kfSendURL = "https://api.weixin.qq.com/cgi-bin/message/custom"
)

// 客服管理
type Kf struct {
	*context.Context
}

func NewCustomerServer(context *context.Context) *Kf {
	Kf := new(Kf)
	Kf.Context = context
	return Kf
}

// 获取客服列表
func (kf *Kf) KfList(msgRequest KfOperResponse, action string) (msgResponse *KfOperResponse, err error) {
	var accessToken string
	accessToken, err = kf.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s/uploadheadimg?access_token=%s", KfOperURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, msgRequest)
	if err != nil {
		return
	}
	msgResponse = new(KfOperResponse)
	err = json.Unmarshal(response, msgResponse)
	if err != nil {
		return
	}

	if msgResponse.ErrCode != 0 {
		err = fmt.Errorf("KfList Error , errcode=%d , errmsg=%s", msgResponse.ErrCode, msgResponse.ErrMsg)
		return
	}
	return
}

// 增加客服
func (kf *Kf) AddKf(msgRequest KfOperResponse) (msgResponse *KfOperResponse, err error) {
	var accessToken string
	accessToken, err = kf.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s/add?access_token=%s", KfOperURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, msgRequest)
	if err != nil {
		return
	}
	msgResponse = new(KfOperResponse)
	err = json.Unmarshal(response, msgResponse)
	if err != nil {
		return
	}

	if msgResponse.ErrCode != 0 {
		err = fmt.Errorf("AddKf Error , errcode=%d , errmsg=%s", msgResponse.ErrCode, msgResponse.ErrMsg)
		return
	}
	return
}

// 修改客服
func (kf *Kf) UpdateKf(msgRequest KfOperResponse) (msgResponse *KfOperResponse, err error) {
	var accessToken string
	accessToken, err = kf.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s/update?access_token=%s", KfOperURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, msgRequest)
	if err != nil {
		return
	}
	msgResponse = new(KfOperResponse)
	err = json.Unmarshal(response, msgResponse)
	if err != nil {
		return
	}

	if msgResponse.ErrCode != 0 {
		err = fmt.Errorf("UpdateKf Error , errcode=%d , errmsg=%s", msgResponse.ErrCode, msgResponse.ErrMsg)
		return
	}
	return
}

// 删除客服
func (kf *Kf) DeleteKf(msgRequest KfOperResponse) (msgResponse *KfOperResponse, err error) {
	var accessToken string
	accessToken, err = kf.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s/del?access_token=%s", KfOperURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, msgRequest)
	if err != nil {
		return
	}
	msgResponse = new(KfOperResponse)
	err = json.Unmarshal(response, msgResponse)
	if err != nil {
		return
	}

	if msgResponse.ErrCode != 0 {
		err = fmt.Errorf("DeleteKf Error , errcode=%d , errmsg=%s", msgResponse.ErrCode, msgResponse.ErrMsg)
		return
	}
	return
}

// 发送文本信息给客户
func (kf *Kf) SendTextMsg(toUser string, content string) (msgResponse *KfSendMsgResponse, err error) {
	var accessToken string
	accessToken, err = kf.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s/send?access_token=%s", kfSendURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, KfSendMsgRequest{
		ToUser:  toUser,
		MsgType: "text",
		Text:    message.Text{Content: content},
	})
	if err != nil {
		return
	}
	msgResponse = new(KfSendMsgResponse)
	err = json.Unmarshal(response, msgResponse)
	if err != nil {
		return
	}

	if msgResponse.ErrCode != 0 {
		err = fmt.Errorf("SendTextMsg Error , errcode=%d , errmsg=%s", msgResponse.ErrCode, msgResponse.ErrMsg)
		return
	}
	return
}

// 发送自定义多媒体消息给客户
func (kf *Kf) Send(msgRequest KfSendMsgRequest) (msgResponse *KfSendMsgResponse, err error) {
	var accessToken string
	accessToken, err = kf.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s/send?access_token=%s", kfSendURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, msgRequest)
	if err != nil {
		return
	}
	msgResponse = new(KfSendMsgResponse)
	err = json.Unmarshal(response, msgResponse)
	if err != nil {
		return
	}

	if msgResponse.ErrCode != 0 {
		err = fmt.Errorf("SendTextMsg Error , errcode=%d , errmsg=%s", msgResponse.ErrCode, msgResponse.ErrMsg)
		return
	}
	return
}

type KfOperRequest struct {
	KfAccount string `json:"kf_account"` // 完整客服账号，格式为：账号前缀@公众号微信号
	NickName  string `json:"nickname"`   // 客服昵称，最长6个汉字或12个英文字符
	PassWord  string `json:"password"`   // 客服账号登录密码，格式为密码明文的32位加密MD5值。该密码仅用于在公众平台官网的多客服功能中使用，若不使用多客服功能，则不必设置密码
}

type KfOperResponse struct {
	util.CommonError
}

type KfListResponse struct {
	util.CommonError
	KfList []struct {
		KfAccount    string `json:"kf_account"` // 完整客服账号，格式为：账号前缀@公众号微信号
		KfNick       string `json:"kf_nick"`    // 客服昵称
		KfId         string `json:"kf_id"`      // 客服工号
		KfHeadImgUrl string `json:"kf_headimgurl"`
	} `json:"kf_list"`
}

type KfSendMsgRequest struct {
	ToUser  string        `json:"touser"`
	MsgType string        `json:"msgtype"`
	Text    message.Text  `json:"text"`
	Image   message.Image `json:"image"`
	Voice   message.Voice `json:"voice"`
	Video   message.Video `json:"video"`
	Music   message.Music `json:"music"`
	News    message.News  `json:"news"`

	// 客服信息（若是要以指定客服账号发送的时候）
	CustomService struct {
		KfAccount string `json:"kf_account"`
	} `json:"customservice"`
}

type KfSendMsgResponse struct {
	util.CommonError
}
