package privacy

import (
	"errors"
	"fmt"

	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

// Privacy 小程序授权隐私设置
type Privacy struct {
	*context.Context
}

// NewCustomerMessageManager 实例化消息管理者
func NewPrivacy(context *context.Context) *Privacy {
	return &Privacy{
		context,
	}
}

type OwnerSetting struct {
	ContactEmail         string `json:"contact_email"`
	ContactPhone         string `json:"contact_phone"`
	ContactQQ            string `json:"contact_qq"`
	ContactWeixin        string `json:"contact_weixin"`
	ExtFileMediaID       string `json:"ext_file_media_id"`
	NoticeMethod         string `json:"notice_method"`
	StoreExpireTimestamp string `json:"store_expire_timestamp"`
}

type SettingItem struct {
	PrivacyKey  string `json:"privacy_key"`
	PrivacyText string `json:"privacy_text"`
}

type SettingResponseItem struct {
	PrivacyKey   string `json:"privacy_key"`
	PrivacyText  string `json:"privacy_text"`
	PrivacyLabel string `json:"privacy_label"`
}

type SetPrivacySettingRequest struct {
	PrivacyVer   int           `json:"privacy_ver"`
	OwnerSetting OwnerSetting  `json:"owner_setting"`
	SettingList  []SettingItem `json:"setting_list"`
}

const (
	setPrivacySettingUrl    = "https://api.weixin.qq.com/cgi-bin/component/setprivacysetting"
	getPrivacySettingUrl    = "https://api.weixin.qq.com/cgi-bin/component/getprivacysetting"
	uploadPrivacyExtFileUrl = "https://api.weixin.qq.com/cgi-bin/component/uploadprivacyextfile"

	PrivacyV1 = 1
	PrivacyV2 = 2
)

type GetPrivacySettingResponse struct {
	util.CommonError
	CodeExist    int                   `json:"code_exist"`
	PrivacyList  []string              `json:"privacy_list"`
	SettingList  []SettingResponseItem `json:"setting_list"`
	UpdateTime   int64                 `json:"update_time"`
	OwnerSetting OwnerSetting          `json:"owner_setting"`
	PrivacyDesc  PrivacyDescList       `json:"privacy_desc"`
}

type PrivacyDescList struct {
	PrivacyDescList []PrivacyDesc `json:"privacy_desc_list"`
}

type PrivacyDesc struct {
	PrivacyDesc string `json:"privacy_desc"`
	PrivacyKey  string `json:"privacy_key"`
}

func (s *Privacy) GetPrivacySetting(privacyVer int) (GetPrivacySettingResponse, error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return GetPrivacySettingResponse{}, err
	}

	response, err := util.PostJSON(fmt.Sprintf("%s?access_token=%s", getPrivacySettingUrl, accessToken), map[string]int{
		"privacy_ver": privacyVer,
	})
	if err != nil {
		return GetPrivacySettingResponse{}, err
	}
	// 返回错误信息
	var result GetPrivacySettingResponse
	if err = util.DecodeWithError(response, &result, "getprivacysetting"); err == nil {
		return GetPrivacySettingResponse{}, err
	}

	return result, nil
}

func (s *Privacy) SetPrivacySetting(privacyVer int, ownerSetting OwnerSetting, settingList []SettingItem) error {
	if privacyVer == PrivacyV1 && len(settingList) > 0 {
		return errors.New("当privacy_ver传2或者不传时，setting_list是必填；当privacy_ver传1时，该参数不可传")
	}
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return err
	}

	response, err := util.PostJSON(fmt.Sprintf("%s?access_token=%s", setPrivacySettingUrl, accessToken), SetPrivacySettingRequest{
		PrivacyVer:   privacyVer,
		OwnerSetting: ownerSetting,
		SettingList:  settingList,
	})
	if err != nil {
		return err
	}

	// 返回错误信息
	if err = util.DecodeWithCommonError(response, "setprivacysetting"); err == nil {
		return err
	}

	return err
}

type UploadPrivacyExtFileResponse struct {
	util.CommonError
	ExtFileMediaID string `json:"ext_file_media_id"`
}

func (s *Privacy) UploadPrivacyExtFile(fileData []byte) (UploadPrivacyExtFileResponse, error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return UploadPrivacyExtFileResponse{}, err
	}

	response, err := util.PostJSON(fmt.Sprintf("%s?access_token=%s", uploadPrivacyExtFileUrl, accessToken), map[string][]byte{
		"file": fileData,
	})
	if err != nil {
		return UploadPrivacyExtFileResponse{}, err
	}

	// 返回错误信息
	var result UploadPrivacyExtFileResponse
	if err = util.DecodeWithError(response, &result, "setprivacysetting"); err == nil {
		return UploadPrivacyExtFileResponse{}, err
	}

	return result, err
}
