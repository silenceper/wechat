package open

import (
	"gitee.com/zhimiao/wechat-sdk/context"
	"gitee.com/zhimiao/wechat-sdk/util"
	"net/url"
)

const (
	// SUCCESS 成功
	SUCCESS string = "success"
)

// Open struct extends context
type Open struct {
	*context.Context
}

// MiniPrograms 代小程序
type MiniPrograms struct {
	Open
	AuthAppID        string
	AuthRefreshToken string
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
		Open:             *o,
		AuthAppID:        appid,
		AuthRefreshToken: refrshToken,
	}
	return miniPrograms
}

func (o *Open) buildRequest(urlStr string, param map[string]string) (requestURL string, err error) {
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
	requestURL = u.String()
	return
}

// fetchData 拉取统计数据
func (o *Open) post(urlStr string, body interface{}) (response []byte, err error) {
	sendURL, err := o.buildRequest(urlStr, nil)
	if err != nil {
		return
	}
	response, err = util.PostJSON(sendURL, body)
	return
}

// fetchData 拉取统计数据
func (o *Open) get(urlStr string, param map[string]string) (response []byte, err error) {
	sendURL, err := o.buildRequest(urlStr, param)
	if err != nil {
		return
	}
	response, err = util.HTTPGet(sendURL)
	return
}

func (m *MiniPrograms) buildRequest(urlStr string, param map[string]string) (requestURL string, err error) {
	accessToken, err := m.GetAuthrAccessToken(m.AuthAppID)
	if err != nil {
		var ret *context.AuthrAccessToken
		ret, err = m.RefreshAuthrToken(m.AuthAppID, m.AuthRefreshToken)
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
	requestURL = u.String()
	return
}

// fetchData 拉取统计数据
func (m *MiniPrograms) post(urlStr string, body interface{}) (response []byte, err error) {
	sendURL, err := m.buildRequest(urlStr, nil)
	if err != nil {
		return
	}
	response, err = util.PostJSON(sendURL, body)
	return
}

// fetchData 拉取统计数据
func (m *MiniPrograms) get(urlStr string, param map[string]string) (response []byte, err error) {
	sendURL, err := m.buildRequest(urlStr, param)
	if err != nil {
		return
	}
	response, err = util.HTTPGet(sendURL)
	return
}
