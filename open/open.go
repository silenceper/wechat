package open

import (
	"fmt"
	"gitee.com/zhimiao/wechat-sdk/context"
	"net/http"
	"net/url"
)

const (
	SUCCESS string = "success"
)

// Open struct extends context
type Open struct {
	*context.Context
}

// NewOpen 创建开放平台句柄
func NewOpen(ctx *context.Context) *Open {
	open := &Open{Context: ctx}
	return open
}

// fetchData 拉取统计数据
func (o *Open) post(urlStr string, body interface{}) (response []byte, err error) {
	var accessToken string
	accessToken, err = wxa.GetAccessToken()
	if err != nil {
		return
	}
	urlStr = fmt.Sprintf(urlStr, accessToken)
	response, err = util.PostJSON(urlStr, body)
	return
}

// fetchData 拉取统计数据
func (o *Open) get(urlStr string, param map[string]string) (response []byte, err error) {
	var accessToken string
	accessToken, err = wxa.GetAccessToken()
	if err != nil {
		return
	}
	urlStr = fmt.Sprintf(urlStr, accessToken, param...)
	response, err = util.HTTPGet(urlStr)
	return
}
