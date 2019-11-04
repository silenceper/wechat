package open

import (
	"gitee.com/zhimiao/wechat-sdk/context"
	"gitee.com/zhimiao/wechat-sdk/util"
	"net/url"
)

const (
	SUCCESS string = "success"
)

type common struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// Open struct extends context
type Open struct {
	*context.Context
}

// NewOpen 创建开放平台句柄
func NewOpen(ctx *context.Context) *Open {
	open := &Open{Context: ctx}
	return open
}

func (o *Open) buildRequest(urlStr string, param map[string]string) (requestUrl string, err error) {
	accessToken, err := o.GetComponentAccessToken()
	if err != nil {
		return
	}
	u, err := url.Parse(urlStr)
	qs := u.Query()
	qs.Add("access_token", accessToken)
	for k, v := range param {
		qs.Set(k, v)
	}
	u.RawQuery = qs.Encode()
	requestUrl = u.String()
	return
}

// fetchData 拉取统计数据
func (o *Open) post(urlStr string, body interface{}) (response []byte, err error) {
	sendUrl, err := o.buildRequest(urlStr, map[string]string{})
	if err != nil {
		return
	}
	response, err = util.PostJSON(sendUrl, body)
	return
}

// fetchData 拉取统计数据
func (o *Open) get(urlStr string, param map[string]string) (response []byte, err error) {
	sendUrl, err := o.buildRequest(urlStr, param)
	if err != nil {
		return
	}
	response, err = util.HTTPGet(sendUrl)
	return
}
