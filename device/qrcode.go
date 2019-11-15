package device

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/util"
)

type resCreateQRCode struct {
	util.CommonError
	DeviceNum int `json:"device_num"`
	CodeList  []struct {
		DeviceId string `json:"device_id"`
		Ticket   string `json:"ticket"`
	} `json:"code_list"`
}

// CreateQRCode 获取设备二维码
func (d *Device) CreateQRCode(devices []string) (res resCreateQRCode, err error) {
	var accessToken string
	if accessToken, err = d.GetAccessToken(); err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", uriQRCode, accessToken)
	req := map[string]interface{}{
		"device_num":     len(devices),
		"device_id_list": devices,
	}
	var response []byte
	if response, err = util.PostJSON(uri, req); err != nil {
		return
	}
	if err = json.Unmarshal(response, &res); err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("DeviceCreateQRCode Error , errcode=%d , errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}

type resVerifyQRCode struct {
	util.CommonError
	DeviceType string `json:"device_type"`
	DeviceId   string `json:"device_id"`
	Mac        string `json:"mac"`
}

// VerifyQRCode 验证设备二维码
func (d *Device) VerifyQRCode(ticket string) (res resVerifyQRCode, err error) {
	var accessToken string
	if accessToken, err = d.GetAccessToken(); err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", uriVerifyQRCode, accessToken)
	req := map[string]interface{}{
		"ticket": ticket,
	}
	fmt.Println(req)
	var response []byte
	if response, err = util.PostJSON(uri, req); err != nil {
		return
	}
	if err = json.Unmarshal(response, &res); err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("DeviceCreateQRCode Error , errcode=%d , errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}
