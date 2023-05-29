package openapi

import (
	"fmt"
	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	clearQuotaUrl            = "https://api.weixin.qq.com/cgi-bin/clear_quota"       // 重置API调用次数
	getApiQuotaUrl           = "https://api.weixin.qq.com/cgi-bin/openapi/quota/get" // 查询API调用额度
	getRidInfoUrl            = "https://api.weixin.qq.com/cgi-bin/openapi/rid/get"   // 查询rid信息
	clearQuotaByAppSecretUrl = "https://api.weixin.qq.com/cgi-bin/clear_quota/v2"    // 使用AppSecret重置 API 调用次数
)

// OpenApi openApi管理
type OpenApi struct {
	*context.Context
}

// NewOpenApi 实例化
func NewOpenApi(ctx *context.Context) *OpenApi {
	return &OpenApi{Context: ctx}
}

// ClearQuota 重置API调用次数
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/openApi-mgnt/clearQuota.html
func (o *OpenApi) ClearQuota() error {
	payload := map[string]string{
		"appid": o.AppID,
	}
	res, err := o.doPostRequest(clearQuotaUrl, payload)
	if err != nil {
		return err
	}

	return util.DecodeWithCommonError(res, "ClearQuota")
}

// ApiQuota API调用额度
type ApiQuota struct {
	util.CommonError
	Quota struct {
		DailyLimit int64 `json:"daily_limit"` // 当天该账号可调用该接口的次数
		Used       int64 `json:"used"`        // 当天已经调用的次数
		Remain     int64 `json:"remain"`      // 当天剩余调用次数
	} `json:"quota"` // 详情
}

// GetApiQuota 查询API调用额度
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/openApi-mgnt/getApiQuota.html
func (o *OpenApi) GetApiQuota(cgiPath string) (quota ApiQuota, err error) {
	payload := map[string]string{
		"cgi_path": cgiPath,
	}
	res, err := o.doPostRequest(getApiQuotaUrl, payload)
	if err != nil {
		return
	}

	err = util.DecodeWithError(res, &quota, "GetApiQuota")
	return
}

// RidInfo rid信息
type RidInfo struct {
	util.CommonError
	Request struct {
		InvokeTime   int64  `json:"invoke_time"`   // 发起请求的时间戳
		CostInMs     int64  `json:"cost_in_ms"`    // 请求毫秒级耗时
		RequestUrl   string `json:"request_url"`   // 请求的URL参数
		RequestBody  string `json:"request_body"`  // post请求的请求参数
		ResponseBody string `json:"response_body"` // 接口请求返回参数
		ClientIp     string `json:"client_ip"`     // 接口请求的客户端ip
	} `json:"request"` // 该rid对应的请求详情
}

// GetRidInfo 查询rid信息
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/openApi-mgnt/getRidInfo.html
func (o *OpenApi) GetRidInfo(rid string) (r RidInfo, err error) {
	payload := map[string]string{
		"rid": rid,
	}
	res, err := o.doPostRequest(getRidInfoUrl, payload)
	if err != nil {
		return
	}

	err = util.DecodeWithError(res, &r, "GetRidInfo")
	return
}

// ClearQuotaByAppSecret 使用AppSecret重置 API 调用次数
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/openApi-mgnt/clearQuotaByAppSecret.html
func (o *OpenApi) ClearQuotaByAppSecret() error {
	uri := fmt.Sprintf("%s?appid=%s&appsecret=%s", clearQuotaByAppSecretUrl, o.AppID, o.AppSecret)
	res, err := util.HTTPPost(uri, "")
	if err != nil {
		return err
	}

	return util.DecodeWithCommonError(res, "ClearQuotaByAppSecret")
}

func (o *OpenApi) doPostRequest(uri string, payload interface{}) ([]byte, error) {
	ak, err := o.GetAccessToken()
	if err != nil {
		return nil, err
	}

	uri = fmt.Sprintf("%s?access_token=%s", uri, ak)
	return util.PostJSON(uri, payload)
}
