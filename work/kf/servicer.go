package kf

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	//添加接待人员
	receptionistAddAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/servicer/add?access_token=%s"
	//删除接待人员
	receptionistDelAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/servicer/del?access_token=%s"
	//获取接待人员列表
	receptionistListAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/servicer/list?access_token=%s&open_kfid=%s"
)

// ReceptionistOptions 添加接待人员请求参数
type ReceptionistOptions struct {
	OpenKFID   string   `json:"open_kfid"`   // 客服帐号ID
	UserIDList []string `json:"userid_list"` // 接待人员userid列表
}

// ReceptionistSchema 添加接待人员响应内容
type ReceptionistSchema struct {
	BaseModel
	ResultList []struct {
		UserID string `json:"userid"`
		BaseModel
	} `json:"result_list"`
}

// ReceptionistAdd 添加接待人员
func (r *Client) ReceptionistAdd(options ReceptionistOptions) (info ReceptionistSchema, err error) {
	data, err := util.PostJSON(fmt.Sprintf(receptionistAddAddr, r.accessToken), options)
	if err != nil {
		return info, err
	}
	_ = json.Unmarshal(data, &info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// ReceptionistDel 删除接待人员
func (r *Client) ReceptionistDel(options ReceptionistOptions) (info ReceptionistSchema, err error) {
	data, err := util.PostJSON(fmt.Sprintf(receptionistDelAddr, r.accessToken), options)
	if err != nil {
		return info, err
	}
	_ = json.Unmarshal(data, &info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// ReceptionistListSchema 获取接待人员列表响应内容
type ReceptionistListSchema struct {
	BaseModel
	ReceptionistList []struct {
		UserID string `json:"userid"` // 接待人员的userid。第三方应用获取到的为密文userid，即open_userid
		Status int    `json:"status"` // 接待人员的接待状态。0:接待中,1:停止接待。第三方应用需具有“管理帐号、分配会话和收发消息”权限才可获取
	} `json:"servicer_list"`
}

// ReceptionistList 获取接待人员列表
func (r *Client) ReceptionistList(kfID string) (info ReceptionistListSchema, err error) {
	data, err := util.HTTPGet(fmt.Sprintf(receptionistListAddr, r.accessToken, kfID))
	if err != nil {
		return info, err
	}
	_ = json.Unmarshal(data, &info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}
