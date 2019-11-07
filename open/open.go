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

type MiniPrograms struct {
	Open
	AppId        string
	RefreshToken string
}

// NewOpen 创建开放平台句柄
func NewOpen(ctx *context.Context) *Open {
	open := &Open{Context: ctx}
	return open
}

// NewMiniPrograms 创建开放平台代小程序句柄
func (o *Open) NewMiniPrograms(appid string, refrshToken string) *MiniPrograms {
	if appid == "" || refrshToken == "" {
		return nil
	}
	miniPrograms := &MiniPrograms{
		Open:         *o,
		AppId:        appid,
		RefreshToken: refrshToken,
	}
	return miniPrograms
}

func (o *Open) buildRequest(urlStr string, param map[string]string) (requestUrl string, err error) {
	accessToken, err := o.GetComponentAccessToken()
	if err != nil {
		return
	}
	u, err := url.Parse(urlStr)
	qs := u.Query()
	qs.Add("access_token", accessToken)
	if param != nil {
		for k, v := range param {
			qs.Set(k, v)
		}
	}
	u.RawQuery = qs.Encode()
	requestUrl = u.String()
	return
}

// fetchData 拉取统计数据
func (o *Open) post(urlStr string, body interface{}) (response []byte, err error) {
	sendUrl, err := o.buildRequest(urlStr, nil)
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

func (m *MiniPrograms) buildRequest(urlStr string, param map[string]string) (requestUrl string, err error) {
	accessToken, err := m.GetAuthrAccessToken(m.AppID)
	if err != nil {
		ret, err := m.RefreshAuthrToken(m.AppID, m.RefreshToken)
		if err != nil {
			return
		}
		accessToken = ret.AccessToken
	}
	u, err := url.Parse(urlStr)
	qs := u.Query()
	qs.Add("access_token", accessToken)
	if param != nil {
		for k, v := range param {
			qs.Set(k, v)
		}
	}
	u.RawQuery = qs.Encode()
	requestUrl = u.String()
	return
}

// fetchData 拉取统计数据
func (m *MiniPrograms) post(urlStr string, body interface{}) (response []byte, err error) {
	sendUrl, err := m.buildRequest(urlStr, nil)
	if err != nil {
		return
	}
	response, err = util.PostJSON(sendUrl, body)
	return
}

// fetchData 拉取统计数据
func (m *MiniPrograms) get(urlStr string, param map[string]string) (response []byte, err error) {
	sendUrl, err := m.buildRequest(urlStr, param)
	if err != nil {
		return
	}
	response, err = util.HTTPGet(sendUrl)
	return
}
