package open

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	// 授权跳转页
	COMPONENT_LOGIN_PAGE_URL = "https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%d"
)

// AuthURL 获取跳转的url地址
func (o *Open) AuthURL(redirectURI string, authType int) (string, error) {
	//url encode
	urlStr := url.QueryEscape(redirectURI)
	precode, err := o.GetPreCode()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(COMPONENT_LOGIN_PAGE_URL, o.AppID, precode, urlStr, authType), nil
}

// Auth 跳转到网页授权
func (o *Open) Auth(req *http.Request, writer http.ResponseWriter, redirectURI string, authType int) error {
	location, err := o.AuthURL(redirectURI, authType)
	if err != nil {
		return err
	}
	http.Redirect(writer, req, location, 302)
	return nil
}
