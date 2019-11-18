package open

import (
	"encoding/json"
	"fmt"
	"gitee.com/zhimiao/wechat-sdk/util"
)

const (
	// AccountBasicInfoURL 获取用户信息
	getAccountBasicInfoURL   = "https://api.weixin.qq.com/cgi-bin/account/getaccountbasicinfo"
	setWebViewDomainURL      = "https://api.weixin.qq.com/wxa/setwebviewdomain"
	modifyDomainURL          = "https://api.weixin.qq.com/wxa/modify_domain"
	changewxasearchstatusURL = "https://api.weixin.qq.com/wxa/changewxasearchstatus"
	getwxasearchstatusURL    = "https://api.weixin.qq.com/wxa/getwxasearchstatus"
)

type Action string

const (
	ActionAdd    Action = "add"
	ActionDelete        = "delete"
	ActionSet           = "set"
	ActionGet           = "get"
)

// GetWxaSearchStatus 通过本接口可以查询小程序当前的隐私设置，即是否可被搜索
func (m *MiniPrograms) GetWxaSearchStatus() (ret bool, err error) {
	var body []byte
	body, err = m.get(getwxasearchstatusURL, nil)
	if err != nil {
		return
	}
	data := util.CommonError{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}
	if data.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", data.ErrCode, data.ErrMsg)
	}

	return
}

// ModifyDomain 设置服务器域名
func (m *MiniPrograms) CanSearch(open bool) (err error) {
	var body []byte
	rmap := map[string]int{
		"status": 0,
	}
	if open {
		rmap["status"] = 1
	}
	body, err = m.post(changewxasearchstatusURL, rmap)
	if err != nil {
		return
	}
	ret := util.CommonError{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
	}
	return
}

// ModifyDomainParam 设置服务器域名
type ModifyDomainParam struct {
	Action          Action   `json:"action"`
	RequestDomain   []string `json:"requestdomain"`
	WSRequestDomain []string `json:"wsrequestdomain"`
	UploadDomain    []string `json:"uploaddomain"`
	DownloadDomain  []string `json:"downloaddomain"`
}

// ModifyDomain 设置服务器域名
func (m *MiniPrograms) ModifyDomain(param SetWebViewDomainURLParam) (err error) {
	var body []byte
	body, err = m.post(modifyDomainURL, param)
	if err != nil {
		return
	}
	ret := util.CommonError{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
	}
	return
}

// SetWebViewDomainURLParam 设置业务域名参数
type SetWebViewDomainURLParam struct {
	Action        Action   `json:"action"`
	WebViewDomain []string `json:"webviewdomain"`
}

// SetWebViewDomain 设置业务域名
func (m *MiniPrograms) SetWebViewDomain(param SetWebViewDomainURLParam) (err error) {
	var body []byte
	body, err = m.post(setWebViewDomainURL, param)
	if err != nil {
		return
	}
	ret := util.CommonError{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}
	if ret.ErrCode != 0 {
		err = fmt.Errorf("[%d]: %s", ret.ErrCode, ret.ErrMsg)
	}
	return
}

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

// GetAccountBasicInfo 调用本 API 可以获取小程序的基本信息 没啥卵用，不知道为啥
func (m *MiniPrograms) GetAccountBasicInfo() (ret AccountBasicInfo, err error) {
	var body []byte
	body, err = m.get(getAccountBasicInfoURL, nil)
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
