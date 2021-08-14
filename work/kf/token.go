package kf

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	//获取调用凭证AccessToken
	getTokenAddr = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
)

// AccessTokenSchema 获取调用凭证响应数据
type AccessTokenSchema struct {
	BaseModel
	AccessToken string `json:"access_token"` // 获取到的凭证，最长为512字节
	ExpiresIn   int    `json:"expires_in"`   // 凭证的有效时间（秒）
}

// GetAccessToken 获取调用凭证access_token
func (r *Client) GetAccessToken() (info AccessTokenSchema, err error) {
	data, err := util.HTTPGet(fmt.Sprintf(getTokenAddr, r.corpID, r.secret))
	if err != nil {
		return info, err
	}
	fmt.Println(string(data))
	_ = json.Unmarshal(data, &info)
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}

// RefreshAccessToken 刷新调用凭证access_token
func (r *Client) RefreshAccessToken() error {
	//初始化AccessToken
	tokenInfo, err := r.GetAccessToken()
	if err != nil {
		return err
	}
	if err = r.setAccessToken(tokenInfo.AccessToken); err != nil {
		return err
	}
	r.accessToken = tokenInfo.AccessToken
	return nil
}

func (r *Client) initAccessToken() error {
	//如果关闭自动缓存则直接刷新AccessToken
	if r.isCloseCache {
		if err := r.RefreshAccessToken(); err != nil {
			return err
		}
		return nil
	}

	//判断是否已初始化完成，如果己初始化则直接返回当前实例
	token := r.getAccessToken()

	if token == "" {
		if err := r.RefreshAccessToken(); err != nil {
			return err
		}
	} else {
		r.accessToken = token
	}
	return nil
}

func (r *Client) getAccessToken() string {
	token, ok := r.cache.Get("wechat:kf:" + r.corpID).(string)
	if !ok {
		return ""
	}
	return token
}

func (r *Client) setAccessToken(token string) error {
	return r.cache.Set("wechat:kf:"+r.corpID, token, r.expireTime)
}
