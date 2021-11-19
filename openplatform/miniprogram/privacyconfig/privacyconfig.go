package privacyconfig

import (
	"fmt"
	openContext "github.com/silenceper/wechat/v2/openplatform/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	uploadPrivacyExtfile = "https://api.weixin.qq.com/cgi-bin/component/uploadprivacyextfile?access_token=%s" //上传小程序用户隐私保护指引
	getPrivacySetting    = "https://api.weixin.qq.com/cgi-bin/component/getprivacysetting?access_token=%s"    //查询小程序用户隐私保护指引
	setPrivacySetting = "https://api.weixin.qq.com/cgi-bin/component/setprivacysetting?access_token=%s"    //配置小程序用户隐私保护指引
)

//Code struct
type PrivacyConfig struct {
	*openContext.Context
	appId string
}

//NewCode 实例
func NewPrivacyConfig(context *openContext.Context, appID string) *PrivacyConfig {
	privacyConfig := new(PrivacyConfig)
	privacyConfig.Context = context
	privacyConfig.appId = appID
	return privacyConfig
}

//上传小程序用户隐私保护指引
func (privacyConfig *PrivacyConfig) UploadPrivacyExtfile(fileUrl string) (fileMediaId string, err error) {
	var accessToken string
	accessToken, err = privacyConfig.GetAuthrAccessToken(privacyConfig.appId)
	if err != nil {
		return
	}
	body, err := util.PostFile("file", fileUrl, fmt.Sprintf(uploadPrivacyExtfile, accessToken))
	if err != nil {
		return
	}
	var result struct {
		util.CommonError
		ExtFileMediaId string `json:"ext_file_media_id"`
	}
	err = util.DecodeWithError(body, &result, "UploadPrivacyExtfile")
	fileMediaId = result.ExtFileMediaId
	return
}

//查询小程序用户隐私保护指引
func (privacyConfig *PrivacyConfig) GetPrivacySetting(privacyVer int8) (result *GetPrivacySettingResponse, err error) {
	var accessToken string
	accessToken, err = privacyConfig.GetAuthrAccessToken(privacyConfig.appId)
	if err != nil {
		return
	}
	var data struct {
		PrivacyVer int8 `json:"privacy_ver"`
	}
	data.PrivacyVer = privacyVer
	body, err := util.PostJSON(fmt.Sprintf(getPrivacySetting, accessToken), data)
	if err != nil {
		return
	}
	result = &GetPrivacySettingResponse{}
	err = util.DecodeWithError(body, result, "GetPrivacySetting")
	return
}

//配置小程序用户隐私保护指引
func (privacyConfig *PrivacyConfig) SetPrivacySetting(data *SetPrivacySettingParams) (err error) {
	var accessToken string
	accessToken, err = privacyConfig.GetAuthrAccessToken(privacyConfig.appId)
	if err != nil {
		return
	}
	body, err := util.PostJSON(fmt.Sprintf(setPrivacySetting, accessToken), data)
	if err != nil {
		return
	}
	var result struct {
		util.CommonError
	}
	err = util.DecodeWithError(body, &result, "SetPrivacySetting")
	return
}