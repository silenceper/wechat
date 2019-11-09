package open

import (
	"encoding/json"
	"fmt"
	"gitee.com/zhimiao/wechat-sdk/util"
)

const (
	// AccountBasicInfoURL 获取用户信息
	AccountBasicInfoURL = "https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo?"
)

// AccountBasicInfo 用户信息详情
type AccountBasicInfo struct {
	util.CommonError
	Appid          string `json:"appid"`
	AccountType    int    `json:"account_type"`
	PrincipalType  int    `json:"principal_type"`
	PrincipalName  string `json:"principal_name"`
	RealnameStatus int    `json:"realname_status"`
	WxVerifyInfo   struct {
		QualificationVerify   bool `json:"qualification_verify"`
		NamingVerify          bool
		AnnualReview          bool
		AnnualReviewBeginTime int64
		AnnualReviewEndTime   int64
	} `json:"wx_verify_info"`
	SignatureInfo struct {
		Signature       string `json:"signature"`
		ModifyUsedCount int    `json:"modify_used_count"`
		ModifyQuota     int    `json:"modify_quota"`
	} `json:"signature_info"`
	HeadImageInfo struct {
		HeadImageURL    string `json:"head_image_url"`
		ModifyUsedCount int    `json:"modify_used_count"`
		ModifyQuota     int    `json:"modify_quota"`
	} `json:"head_image_info"`
}

// GetAccountBasicInfo 调用本 API 可以获取小程序的基本信息
func (m *MiniPrograms) GetAccountBasicInfo() (ret AccountBasicInfo, err error) {
	var body []byte
	body, err = m.get(AccountBasicInfoURL, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
	}
	return
}
