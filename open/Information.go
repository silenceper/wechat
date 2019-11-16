package open

import (
	"encoding/json"
	"fmt"
	"gitee.com/zhimiao/wechat-sdk/util"
)

const (
	// AccountBasicInfoURL 获取用户信息
	AccountBasicInfoURL = "https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo?"
	// GetCodePageURL 获取已上传的代码的页面列表
	GetCodePageURL = "https://api.weixin.qq.com/wxa/get_page"
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

type CodePageList struct {
	util.CommonError
	PageList []string `json:"page_list"`
}

// GetAccountBasicInfo 调用本 API 可以获取小程序的基本信息 没啥卵用，不知道为啥
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

// GetCodePage 获取已上传的代码的页面列表
func (m *MiniPrograms) GetCodePage() (ret CodePageList, err error) {
	var body []byte
	body, err = m.get(GetCodePageURL, nil)
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
