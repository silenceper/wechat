package device

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/util"
)

const (
	deviceAuthorize = "https://api.weixin.qq.com/device/authorize_device"
	deviceQRCode    = "https://api.weixin.qq.com/device/create_qrcode"
)

//Device struct
type Device struct {
	*context.Context
}

//NewDevice 实例
func NewDevice(context *context.Context) *Device {
	device := new(Device)
	device.Context = context
	return device
}

// DeviceAuthorize 设备授权
func (d *Device) DeviceAuthorize(devices []ReqDevice, opType int, productId string) (res []resBaseInfo, err error) {
	var accessToken string
	accessToken, err = d.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("%s?access_token=%s", deviceAuthorize, accessToken)
	req := reqDeviceAuthorize{
		DeviceNum:  fmt.Sprintf("%d", len(devices)),
		DeviceList: devices,
		OpType:     fmt.Sprintf("%d", opType),
		ProductId:  productId,
	}
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}
	var result ResDeviceAuthorize
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DeviceAuthorize Error , errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
		return
	}
	res = result.Resp
	return
}

// CreateQRCode 获取设备二维码
func (d *Device) CreateQRCode(devices []string) (res []resQRCode, err error) {
	var accessToken string
	accessToken, err = d.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", deviceQRCode, accessToken)
	req := map[string]interface{}{
		"device_num":     len(devices),
		"device_id_list": devices,
	}
	fmt.Println(req)
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return
	}
	var result ResCreateQRCode
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateQRCode Error , errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
		return
	}
	res = result.CodeList
	return
}
