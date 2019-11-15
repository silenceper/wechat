package device

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/util"
)

type ReqBind struct {
	Ticket   string `json:"ticket"`
	DeviceId string `json:"device_id"`
	OpenId   string `json:"open_id"`
}
type resBind struct {
	BaseResp util.CommonError `json:"base_resp"`
}

// Bind 设备绑定
func (d *Device) Bind(req ReqBind) (err error) {
	var accessToken string
	if accessToken, err = d.GetAccessToken(); err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", uriBind, accessToken)
	var response []byte
	if response, err = util.PostJSON(uri, req); err != nil {
		return
	}
	var result resBind
	if err = json.Unmarshal(response, result); err != nil {
		return
	}
	if result.BaseResp.ErrCode != 0 {
		err = fmt.Errorf("DeviceBind Error , errcode=%d , errmsg=%s", result.BaseResp.ErrCode, result.BaseResp.ErrMsg)
		return
	}
	return
}

// Bind 设备解绑
func (d *Device) Unbind(req ReqBind) (err error) {
	var accessToken string
	if accessToken, err = d.GetAccessToken(); err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", uriUnbind, accessToken)
	var response []byte
	if response, err = util.PostJSON(uri, req); err != nil {
		return
	}
	var result resBind
	if err = json.Unmarshal(response, result); err != nil {
		return
	}
	if result.BaseResp.ErrCode != 0 {
		err = fmt.Errorf("DeviceBind Error , errcode=%d , errmsg=%s", result.BaseResp.ErrCode, result.BaseResp.ErrMsg)
		return
	}
	return
}
