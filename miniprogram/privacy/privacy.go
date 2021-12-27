package privacy

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

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

type SetPrivacySettingRequest struct {
	PrivacyVer   int           `json:"privacy_ver"`
	OwnerSetting OwnerSetting  `json:"owner_setting"`
	SettingList  []SettingItem `json:"setting_list"`
}

const (
	setPrivacySettingUrl = "https://api.weixin.qq.com/cgi-bin/component/setprivacysetting"
	PrivacyV1            = 1
	PrivacyV2            = 2
)

func (s *Privacy) SetPrivacySetting(privacyVer int, ownerSetting OwnerSetting, settingList []SettingItem) error {
	if privacyVer == PrivacyV1 && len(settingList) > 0 {
		return errors.New("当privacy_ver传2或者不传时，setting_list是必填；当privacy_ver传1时，该参数不可传")
	}
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return err
	}

	var (
		contentType string
		response    []byte
	)
	response, contentType, err = util.PostJSONWithRespContentType(fmt.Sprintf("%s?access_token=%s", setPrivacySettingUrl, accessToken), SetPrivacySettingRequest{
		PrivacyVer:   privacyVer,
		OwnerSetting: ownerSetting,
		SettingList:  settingList,
	})
	if err != nil {
		return err
	}
	if strings.HasPrefix(contentType, "application/json") {
		// 返回错误信息
		var result util.CommonError
		err = json.Unmarshal(response, &result)
		if err == nil && result.ErrCode != 0 {
			err = fmt.Errorf("fetchCode error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
			return err
		}
	} else {
		err = fmt.Errorf("fetchCode error : unknown response content type - %v", contentType)
	}
	return err
}
