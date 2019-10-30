package open

import (
	"fmt"
	"gitee.com/zhimiao/wechat-sdk/context"
	"net/http"
	"net/url"
)

const (
	// 授权跳转页
	componentloginpageURL = "https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%d"
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

// AuthURL 获取跳转的url地址
func (open *Open) AuthURL(redirectURI string, authType int) (string, error) {
	//url encode
	urlStr := url.QueryEscape(redirectURI)
	precode, err := open.GetPreCode()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(componentloginpageURL, open.AppID, precode, urlStr, authType), nil
}

// Auth 跳转到网页授权
func (open *Open) Auth(req *http.Request, writer http.ResponseWriter, redirectURI string, authType int) error {
	location, err := open.AuthURL(redirectURI, authType)
	if err != nil {
		return err
	}
	http.Redirect(writer, req, location, 302)
	return nil
}

