package context

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/silenceper/wechat/v2/util"
)

const (
	componentAccessTokenURL = "https://api.weixin.qq.com/cgi-bin/component/api_component_token"
	getPreCodeURL           = "https://api.weixin.qq.com/cgi-bin/component/api_create_preauthcode?component_access_token=%s"
	queryAuthURL            = "https://api.weixin.qq.com/cgi-bin/component/api_query_auth?component_access_token=%s"
	refreshTokenURL         = "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token=%s"
	getComponentInfoURL     = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_info?component_access_token=%s"
	getComponentConfigURL   = "https://api.weixin.qq.com/cgi-bin/component/api_get_authorizer_option?component_access_token=%s"
)

// ComponentAccessToken 第三方平台
type ComponentAccessToken struct {
	AccessToken string `json:"component_access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// GetComponentAccessToken 获取 ComponentAccessToken
func (ctx *Context) GetComponentAccessToken() (string, error) {
	accessTokenCacheKey := fmt.Sprintf("component_access_token_%s", ctx.AppID)
	val := ctx.Cache.Get(accessTokenCacheKey)
	if val == nil {
		return "", fmt.Errorf("cann't get component access token")
	}
	return val.(string), nil
}

// SetComponentAccessToken 通过component_verify_ticket 获取 ComponentAccessToken
func (ctx *Context) SetComponentAccessToken(verifyTicket string) (*ComponentAccessToken, error) {
	body := map[string]string{
		"component_appid":         ctx.AppID,
		"component_appsecret":     ctx.AppSecret,
		"component_verify_ticket": verifyTicket,
	}
	respBody, err := util.PostJSON(componentAccessTokenURL, body)
	if err != nil {
		return nil, err
	}

	at := &ComponentAccessToken{}
	if err := json.Unmarshal(respBody, at); err != nil {
		return nil, err
	}

	accessTokenCacheKey := fmt.Sprintf("component_access_token_%s", ctx.AppID)
	expires := at.ExpiresIn - 1500
	if err := ctx.Cache.Set(accessTokenCacheKey, at.AccessToken, time.Duration(expires)*time.Second); err != nil {
		return nil, nil
	}
	return at, nil
}

// GetPreCode 获取预授权码
func (ctx *Context) GetPreCode() (string, error) {
	cat, err := ctx.GetComponentAccessToken()
	if err != nil {
		return "", err
	}
	req := map[string]string{
		"component_appid": ctx.AppID,
	}
	uri := fmt.Sprintf(getPreCodeURL, cat)
	body, err := util.PostJSON(uri, req)
	if err != nil {
		return "", err
	}

	var ret struct {
		PreCode string `json:"pre_auth_code"`
	}
	if err := json.Unmarshal(body, &ret); err != nil {
		return "", err
	}

	return ret.PreCode, nil
}

// ID 微信返回接口中各种类型字段
type ID struct {
	ID int `json:"id"`
}

// AuthBaseInfo 授权的基本信息
type AuthBaseInfo struct {
	AuthrAccessToken
	FuncInfo []AuthFuncInfo `json:"func_info"`
}

// AuthFuncInfo 授权的接口内容
type AuthFuncInfo struct {
	FuncscopeCategory ID `json:"funcscope_category"`
}

// AuthrAccessToken 授权方AccessToken
type AuthrAccessToken struct {
	Appid        string `json:"authorizer_appid"`
	AccessToken  string `json:"authorizer_access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"authorizer_refresh_token"`
}

// QueryAuthCode 使用授权码换取公众号或小程序的接口调用凭据和授权信息
func (ctx *Context) QueryAuthCode(authCode string) (*AuthBaseInfo, error) {
	cat, err := ctx.GetComponentAccessToken()
	if err != nil {
		return nil, err
	}

	req := map[string]string{
		"component_appid":    ctx.AppID,
		"authorization_code": authCode,
	}
	uri := fmt.Sprintf(queryAuthURL, cat)
	body, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	var ret struct {
		Info *AuthBaseInfo `json:"authorization_info"`
	}

	if err := json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}

	return ret.Info, nil
}

// RefreshAuthrToken 获取（刷新）授权公众号或小程序的接口调用凭据（令牌）
func (ctx *Context) RefreshAuthrToken(appid, refreshToken string) (*AuthrAccessToken, error) {
	cat, err := ctx.GetComponentAccessToken()
	if err != nil {
		return nil, err
	}

	req := map[string]string{
		"component_appid":          ctx.AppID,
		"authorizer_appid":         appid,
		"authorizer_refresh_token": refreshToken,
	}
	uri := fmt.Sprintf(refreshTokenURL, cat)
	body, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}

	ret := &AuthrAccessToken{}
	if err := json.Unmarshal(body, ret); err != nil {
		return nil, err
	}

	authrTokenKey := "authorizer_access_token_" + appid
	if err := ctx.Cache.Set(authrTokenKey, ret.AccessToken, time.Minute*80); err != nil {
		return nil, err
	}
	return ret, nil
}

// GetAuthrAccessToken 获取授权方AccessToken
func (ctx *Context) GetAuthrAccessToken(appid string) (string, error) {
	authrTokenKey := "authorizer_access_token_" + appid
	val := ctx.Cache.Get(authrTokenKey)
	if val == nil {
		return "", fmt.Errorf("cannot get authorizer %s access token", appid)
	}
	return val.(string), nil
}

// AuthorizerInfo 授权方详细信息
type AuthorizerInfo struct {
	NickName        string `json:"nick_name"`
	HeadImg         string `json:"head_img"`
	ServiceTypeInfo ID     `json:"service_type_info"`
	VerifyTypeInfo  ID     `json:"verify_type_info"`
	UserName        string `json:"user_name"`
	PrincipalName   string `json:"principal_name"`
	BusinessInfo    struct {
		OpenStore string `json:"open_store"`
		OpenScan  string `json:"open_scan"`
		OpenPay   string `json:"open_pay"`
		OpenCard  string `json:"open_card"`
		OpenShake string `json:"open_shake"`
	}
	Alias     string `json:"alias"`
	QrcodeURL string `json:"qrcode_url"`
}

// GetAuthrInfo 获取授权方的帐号基本信息
func (ctx *Context) GetAuthrInfo(appid string) (*AuthorizerInfo, *AuthBaseInfo, error) {
	cat, err := ctx.GetComponentAccessToken()
	if err != nil {
		return nil, nil, err
	}

	req := map[string]string{
		"component_appid":  ctx.AppID,
		"authorizer_appid": appid,
	}

	uri := fmt.Sprintf(getComponentInfoURL, cat)
	body, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, nil, err
	}

	var ret struct {
		AuthorizerInfo    *AuthorizerInfo `json:"authorizer_info"`
		AuthorizationInfo *AuthBaseInfo   `json:"authorization_info"`
	}
	if err := json.Unmarshal(body, &ret); err != nil {
		return nil, nil, err
	}

	return ret.AuthorizerInfo, ret.AuthorizationInfo, nil
}
