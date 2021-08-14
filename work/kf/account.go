package kf

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	//添加客服账号
	accountAddAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/add?access_token=%s"
	// 删除客服账号
	accountDelAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/del?access_token=%s"
	// 修改客服账号
	accountUpdateAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/update?access_token=%s"
	// 获取客服账号列表
	accountListAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/list?access_token=%s"
	//获取客服账号链接
	addContactWayAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/add_contact_way?access_token=%s"
)

// AccountAddOptions 添加客服账号请求参数
type AccountAddOptions struct {
	Name    string `json:"name"`     // 客服帐号名称, 不多于16个字符
	MediaID string `json:"media_id"` // 客服头像临时素材。可以调用上传临时素材接口获取, 不多于128个字节
}

// AccountAddSchema 添加客服账号响应内容
type AccountAddSchema struct {
	BaseModel
	OpenKFID string `json:"open_kfid"` // 新创建的客服张号ID
}

// AccountAdd 添加客服账号
func (r *Client) AccountAdd(options AccountAddOptions) (info AccountAddSchema, err error) {
	data, err := util.PostJSON(fmt.Sprintf(accountAddAddr, r.accessToken), options)
	if err != nil {
		return info, err
	}
	_ = json.Unmarshal(data, &info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// AccountDelOptions 删除客服账号请求参数
type AccountDelOptions struct {
	OpenKFID string `json:"open_kfid"` // 客服帐号ID, 不多于64字节
}

// AccountDel 删除客服账号
func (r *Client) AccountDel(options AccountDelOptions) (info BaseModel, err error) {
	data, err := util.PostJSON(fmt.Sprintf(accountDelAddr, r.accessToken), options)
	if err != nil {
		return info, err
	}
	_ = json.Unmarshal(data, &info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// AccountUpdateOptions 修改客服账号请求参数
type AccountUpdateOptions struct {
	OpenKFID string `json:"open_kfid"` // 客服帐号ID, 不多于64字节
	Name     string `json:"name"`      // 客服帐号名称, 不多于16个字符
	MediaID  string `json:"media_id"`  // 客服头像临时素材。可以调用上传临时素材接口获取, 不多于128个字节
}

// AccountUpdate 修复客服账号
func (r *Client) AccountUpdate(options AccountUpdateOptions) (info BaseModel, err error) {
	data, err := util.PostJSON(fmt.Sprintf(accountUpdateAddr, r.accessToken), options)
	if err != nil {
		return info, err
	}
	_ = json.Unmarshal(data, &info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// AccountInfoSchema 客服详情
type AccountInfoSchema struct {
	OpenKFID string `json:"open_kfid"` // 客服帐号ID
	Name     string `json:"name"`      // 客服帐号名称
	Avatar   string `json:"avatar"`    // 客服头像URL
}

// AccountListSchema 获取客服账号列表响应内容
type AccountListSchema struct {
	BaseModel
	AccountList []AccountInfoSchema `json:"account_list"` // 客服账号列表
}

// AccountList 获取客服账号列表
func (r *Client) AccountList() (info AccountListSchema, err error) {
	data, err := util.HTTPGet(fmt.Sprintf(accountListAddr, r.accessToken))
	if err != nil {
		return info, err
	}
	_ = json.Unmarshal(data, &info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// AddContactWayOptions 获取客服账号链接
type AddContactWayOptions struct {
	OpenKFID string `json:"open_kfid"` // 客服帐号ID, 不多于64字节
	Scene    string `json:"scene"`     // 场景值，字符串类型，由开发者自定义, 不多于32字节, 字符串取值范围(正则表达式)：[0-9a-zA-Z_-]*
}

// AddContactWaySchema 获取客服账号链接响应内容
type AddContactWaySchema struct {
	BaseModel
	URL string `json:"url"` // 客服链接，开发者可将该链接嵌入到H5页面中，用户点击链接即可向对应的微信客服帐号发起咨询。开发者也可根据该url自行生成需要的二维码图片
}

// AddContactWay 获取客服账号链接
func (r *Client) AddContactWay(options AddContactWayOptions) (info AddContactWaySchema, err error) {
	data, err := util.PostJSON(fmt.Sprintf(addContactWayAddr, r.accessToken), options)
	if err != nil {
		return info, err
	}
	_ = json.Unmarshal(data, &info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}
