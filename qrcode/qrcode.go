package qrcode

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/util"
)

const (
	baseApi            = "https://api.weixin.qq.com/cgi-bin"
	createURL          = "/qrcode/create"
	QR_SCENE           = "QR_SCENE"
	QR_STR_SCENE       = "QR_STR_SCENE"
	QR_LIMIT_SCENE     = "QR_LIMIT_SCENE"
	QR_LIMIT_STR_SCENE = "QR_LIMIT_STR_SCENE"
)

type QrCode struct {
	*context.Context
}

type actionInfo struct {
	Scene map[string]string `json:"scene,omitempty"`
}

type reqQrCode struct {
	ActionName    string     `json:"action_name,omitempty"`
	ExpireSeconds int64      `json:"expire_seconds,omitempty"`
	ActionInfo    actionInfo `json:"action_info,omitempty"`
}

type ResCreate struct {
	util.CommonError
	Ticket string `json:"ticket,omitempty"`
	URL    string `json:"url,omitempty"`
}

//NewQrCode 实例
func NewQrCode(context *context.Context) *QrCode {
	qrcode := new(QrCode)
	qrcode.Context = context
	return qrcode
}

func (qrcode *QrCode) GetQrCode(actionName string, scene map[string]string, expireSec int64) (resCreate ResCreate, err error) {
	var accessToken string
	accessToken, err = qrcode.GetAccessToken()
	if err != nil {
		return
	}

	actionInfo := actionInfo{
		Scene: scene,
	}
	reqParams := reqQrCode{
		ActionName:    actionName,
		ExpireSeconds: expireSec,
		ActionInfo:    actionInfo,
	}

	fmt.Println("reqParams", reqParams)

	uri := fmt.Sprintf("%s%s?access_token=%s", baseApi, createURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, reqParams)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &resCreate)
	if err != nil {
		return
	}

	if resCreate.ErrCode != 0 {
		err = fmt.Errorf("AddGroup Error , errcode=%d , errmsg=%s", resCreate.ErrCode, resCreate.ErrMsg)
		return
	}
	return
}
