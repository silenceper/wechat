package basic

import (
	"fmt"

	openContext "github.com/silenceper/wechat/v2/openplatform/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	getAccountBasicInfoURL = "https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo"
	modifyDomain = "https://api.weixin.qq.com/wxa/modify_domain"
	modifyWebviewDomain = "https://api.weixin.qq.com/wxa/setwebviewdomain"
)

//Basic 基础信息设置
type Basic struct {
	*openContext.Context
	appID string
}
//ModifyDomain 修改服务器域名请求参数
type ModifyDomain struct {
	Action string `json:"action"`
	RequestDomain []string `json:"requestdomain"`
	WsRequestDomain []string `json:"wsrequestdomain"`
	UploadDomain []string `json:"uploaddomain"`
	DownloadDomain []string `json:"downloaddomain"`
}
//ModifyDomainRes 修改服务器域名请求结果
type ModifyDomainRes struct {
	util.CommonError
	RequestDomain []string `json:"requestdomain"`
	WsRequestDomain []string `json:"wsrequestdomain"`
	UploadDomain []string `json:"uploaddomain"`
	DownloadDomain []string `json:"downloaddomain"`
}
//ModifyWebviewDomain 修改业务域名请求参数
type ModifyWebviewDomain struct {
	Action string `json:"action"`
	WebviewDomain []string `json:"webviewdomain"`
}
//ModifyWebviewDomainRes 修改业务域名请求结果
type ModifyWebviewDomainRes struct {
	util.CommonError
}

//NewBasic new
func NewBasic(opContext *openContext.Context, appID string) *Basic {
	return &Basic{Context: opContext, appID: appID}
}

//AccountBasicInfo 基础信息
type AccountBasicInfo struct {
	util.CommonError
}

//GetAccountBasicInfo 获取小程序基础信息
//reference:https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/Mini_Program_Information_Settings.html
func (basic *Basic) GetAccountBasicInfo() (*AccountBasicInfo, error) {
	ak, err := basic.GetAuthrAccessToken(basic.AppID)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s?access_token=%s", getAccountBasicInfoURL, ak)
	data, err := util.HTTPGet(url)
	if err != nil {
		return nil, err
	}
	result := &AccountBasicInfo{}
	if err := util.DecodeWithError(data, result, "account/getaccountbasicinfo"); err != nil {
		return nil, err
	}
	return result, nil
}

//ModifyDomain 设置服务器域名
func (encryptor *Basic) ModifyDomain(data *ModifyDomain)(result *ModifyDomainRes,err error) {
	var accessToken string
	accessToken, err = encryptor.GetAuthrAccessToken(encryptor.appID)
	if err != nil {
		return
	}

	urlStr := fmt.Sprintf("%s?access_token=%s", modifyDomain, accessToken)
	body, err := util.PostJSON(urlStr, data)
	if err != nil {
		return
	}
	// 返回错误信息
	result = &ModifyDomainRes{}
	err = util.DecodeWithError(body, result, "modifyDomain")
	return
}

//ModifyWebviewDomain 设置业务域名
func (encryptor *Basic) ModifyWebviewDomain(data *ModifyWebviewDomain)(result *ModifyWebviewDomainRes,err error) {
	var accessToken string
	accessToken, err = encryptor.GetAuthrAccessToken(encryptor.appID)
	if err != nil {
		return
	}

	urlStr := fmt.Sprintf("%s?access_token=%s", modifyWebviewDomain, accessToken)
	body, err := util.PostJSON(urlStr, data)
	if err != nil {
		return
	}
	// 返回错误信息
	result = &ModifyWebviewDomainRes{}
	err = util.DecodeWithError(body, result, "modifyWebviewDomain")
	return
}
