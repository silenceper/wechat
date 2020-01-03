package context

import (
	"encoding/json"
	"fmt"
	"gitee.com/zhimiao/wechat-sdk/cache"
	"time"

	"gitee.com/zhimiao/wechat-sdk/util"
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
	util.CommonError
	AccessToken string `json:"component_access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// SetComponentVerifyTicket 保存每10min一次的微信令牌
func (ctx *Context) SetComponentVerifyTicket(ticket string) {
	err := ctx.Cache.Set(fmt.Sprintf(cache.ComponentVerifyTicket, ctx.AppID), ticket, 0)
	if err != nil {
		fmt.Printf("保存票据报错: %v\n", err)
	}
}

// GetComponentVerifyTicket 获取票据
func (ctx *Context) GetComponentVerifyTicket() (string, error) {
	val := ctx.Cache.Get(fmt.Sprintf(cache.ComponentVerifyTicket, ctx.AppID))
	if val == nil {
		return "", fmt.Errorf("cann't get component verify ticket")
	}
	if v, ok := val.(string); ok {
		return v, nil
	}
	return "", fmt.Errorf("cann't get component verify ticket")
}

// GetComponentAccessToken 获取 ComponentAccessToken
func (ctx *Context) GetComponentAccessToken() (string, error) {
	accessTokenCacheKey := fmt.Sprintf(cache.ComponentAccessToken, ctx.AppID)
	var result string
	t := ctx.Cache.Get(accessTokenCacheKey)
	if v, ok := t.(string); ok {
		result = v
	}
	if result == "" {
		t, err := ctx.GetComponentVerifyTicket()
		if err != nil {
			return "", err
		}
		at, err := ctx.SetComponentAccessToken(t)
		if err != nil {
			return "", err
		}
		result = at.AccessToken
	}
	if result == "" {
		return "", fmt.Errorf("ComponentAccessToken 获取失败")
	}
	return result, nil
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
	if at.ErrCode != 0 {
		return nil, fmt.Errorf("Componet access token err  [%d]: %s", at.ErrCode, at.ErrMsg)
	}
	accessTokenCacheKey := fmt.Sprintf(cache.ComponentAccessToken, ctx.AppID)
	expires := at.ExpiresIn - 1500
	err = ctx.Cache.Set(accessTokenCacheKey, at.AccessToken, time.Duration(expires)*time.Second)
	if err != nil {
		fmt.Println("Componet access token err ", err.Error())
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
	fmt.Println(string(body))
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
	ConfirmInfo       struct {
		NeedConfirm    int8 `json:"need_confirm"`
		AlreadyConfirm int8 `json:"already_confirm"`
		CanConfirm     int8 `json:"can_confirm"`
	} `json:"confirm_info"`
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
	ctx.Cache.Set(authrTokenKey, ret.AccessToken, time.Minute*80)

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
	NickName        string          `json:"nick_name"`
	HeadImg         string          `json:"head_img"`
	ServiceTypeInfo ID              `json:"service_type_info"`
	VerifyTypeInfo  ID              `json:"verify_type_info"`
	UserName        string          `json:"user_name"`
	Alias           string          `json:"alias"`
	QrcodeURL       string          `json:"qrcode_url"`
	BusinessInfo    map[string]int8 `json:"business_info"`
	PrincipalName   string          `json:"principal_name"`
	Signature       string          `json:"signature"` // 小程序名称
	// open_store	是否开通微信门店功能
	// open_scan	是否开通微信扫商品功能
	// open_pay	是否开通微信支付功能
	// open_card	是否开通微信卡券功能
	// open_shake	是否开通微信摇一摇功能
	MiniProgramInfo struct {
		Network struct {
			RequestDomain   []string
			WsRequestDomain []string
			UploadDomain    []string
			DownloadDomain  []string
			BizDomain       []string
			UDPDomain       []string
		} `json:"network"`
		Categories []struct {
			First  string `json:"first"`
			Second string `json:"second"`
		} `json:"categories"`
		VisitStatus int8 `json:"visit_status"`
	}
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
