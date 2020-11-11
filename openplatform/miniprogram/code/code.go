package code

import (
	"fmt"
	openContext "github.com/silenceper/wechat/v2/openplatform/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	uploadCode           = "https://api.weixin.qq.com/wxa/commit?access_token=%s"                 //上传代码
	getQrcode            = "https://api.weixin.qq.com/wxa/get_qrcode?access_token=%s&path=%s"     //获取体验二维码
	submitAudit          = "https://api.weixin.qq.com/wxa/submit_audit?access_token=%s"           //提交审核
	getAuditstatus       = "https://api.weixin.qq.com/wxa/get_auditstatus?access_token=%s"        //获取审核状态
	getLatestAuditstatus = "https://api.weixin.qq.com/wxa/get_latest_auditstatus?access_token=%s" //获取最后一次审核状态
	release              = "https://api.weixin.qq.com/wxa/release?access_token=%s"                //发布代码
	auditRecall          = "https://api.weixin.qq.com/wxa/undocodeaudit?access_token=%s"          //审核撤回
)

//Code struct
type Code struct {
	*openContext.Context
	appId string
}

//NewCode 实例
func NewCode(context *openContext.Context, appID string) *Code {
	Code := new(Code)
	Code.Context = context
	Code.appId = appID
	return Code
}

// UploadCode 上传代码
func (code *Code) UploadCode(data *UploadCodeParams) (err error) {
	var accessToken string

	accessToken, err = code.GetAuthrAccessToken(code.appId)
	if err != nil {
		return
	}

	urlStr := fmt.Sprintf(uploadCode, accessToken)
	body, err := util.PostJSON(urlStr, data)
	if err != nil {
		return
	}
	// 返回错误信息
	var result struct {
		util.CommonError
	}
	err = util.DecodeWithError(body, &result, "UploadCode")

	return
}

// SubmitAudit 代码提交审核
func (code *Code) SubmitAudit(data *SubmitAuditParams) (auditid int64, err error) {
	var accessToken string
	accessToken, err = code.GetAuthrAccessToken(code.appId)
	if err != nil {
		return
	}

	urlStr := fmt.Sprintf(submitAudit, accessToken)
	body, err := util.PostJSON(urlStr, data)
	if err != nil {
		return
	}
	// 返回错误信息
	var result struct {
		util.CommonError
		Auditid int64 `json:"auditid"`
	}
	err = util.DecodeWithError(body, &result, "SubmitAudit")
	auditid = result.Auditid
	return
}

// GetAuditstatus 获取指定审核状态
func (code *Code) GetAuditstatus(auditid string) (status int64, reason string, screenshot string, err error) {
	var accessToken string
	accessToken, err = code.GetAuthrAccessToken(code.appId)
	if err != nil {
		return
	}

	urlStr := fmt.Sprintf(getAuditstatus, accessToken)
	var data struct {
		Auditid string `json:"auditid"`
	}
	data.Auditid = auditid
	body, err := util.PostJSON(urlStr, data)
	if err != nil {
		return
	}
	// 返回错误信息
	var result = &GetAuditstatusResponse{}
	err = util.DecodeWithError(body, result, "GetAuditstatus")
	status = result.Status
	reason = result.Reason
	screenshot = result.Screenshot
	return
}

// GetLastAuditstatus 查询最新一次提交的审核状态
func (code *Code) GetLastAuditstatus() (auditid int64, status int64, reason string, screenshot string, err error) {
	var accessToken string
	accessToken, err = code.GetAuthrAccessToken(code.appId)
	if err != nil {
		return
	}

	urlStr := fmt.Sprintf(getLatestAuditstatus, accessToken)

	body, err := util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	// 返回错误信息
	var result = &GetLastAuditstatusResponse{}
	err = util.DecodeWithError(body, result, "GetLastAuditstatus")
	status = result.Status
	reason = result.Reason
	screenshot = result.Screenshot
	auditid = result.Auditid
	return
}

// Release 代码发布
func (code *Code) Release() (err error) {
	var accessToken string
	accessToken, err = code.GetAuthrAccessToken(code.appId)
	if err != nil {
		return
	}

	urlStr := fmt.Sprintf(release, accessToken)

	body, err := util.PostJSON(urlStr, struct {}{})
	if err != nil {
		return
	}
	// 返回错误信息
	var result struct {
		util.CommonError
	}
	err = util.DecodeWithError(body, &result, "Release")

	return
}

func (code *Code) GetQrCode() (body []byte, err error) {
	var accessToken string
	accessToken, err = code.GetAuthrAccessToken(code.appId)
	if err != nil {
		return
	}
	urlStr := fmt.Sprintf(getQrcode, accessToken, "")
	body, err = util.HTTPGet(urlStr)
	return
}

// GetLastAuditstatus 查询最新一次提交的审核状态
func (code *Code) AuditRecall() (err error) {
	var accessToken string
	accessToken, err = code.GetAuthrAccessToken(code.appId)
	if err != nil {
		return
	}

	urlStr := fmt.Sprintf(auditRecall, accessToken)

	body, err := util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	var result struct {
		util.CommonError
	}
	err = util.DecodeWithError(body, &result, "AuditRecall")
	return
}
