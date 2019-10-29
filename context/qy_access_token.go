package context

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"gitee.com/zhimiao/wechat-sdk/util"
)

const (
	//qyAccessTokenURL 获取access_token的接口
	qyAccessTokenURL = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
)

//ResQyAccessToken struct
type ResQyAccessToken struct {
	util.CommonError

	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

//SetQyAccessTokenLock 设置读写锁（一个appID一个读写锁）
func (ctx *Context) SetQyAccessTokenLock(l *sync.RWMutex) {
	ctx.accessTokenLock = l
}

//GetQyAccessToken 获取access_token
func (ctx *Context) GetQyAccessToken() (accessToken string, err error) {
	ctx.accessTokenLock.Lock()
	defer ctx.accessTokenLock.Unlock()

	accessTokenCacheKey := fmt.Sprintf("qy_access_token_%s", ctx.AppID)
	val := ctx.Cache.Get(accessTokenCacheKey)
	if val != nil {
		accessToken = val.(string)
		return
	}

	//从微信服务器获取
	var resQyAccessToken ResQyAccessToken
	resQyAccessToken, err = ctx.GetQyAccessTokenFromServer()
	if err != nil {
		return
	}

	accessToken = resQyAccessToken.AccessToken
	return
}

//GetQyAccessTokenFromServer 强制从微信服务器获取token
func (ctx *Context) GetQyAccessTokenFromServer() (resQyAccessToken ResQyAccessToken, err error) {
	log.Printf("GetQyAccessTokenFromServer")
	url := fmt.Sprintf(qyAccessTokenURL, ctx.AppID, ctx.AppSecret)
	var body []byte
	body, err = util.HTTPGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resQyAccessToken)
	if err != nil {
		return
	}
	if resQyAccessToken.ErrCode != 0 {
		err = fmt.Errorf("get qy_access_token error : errcode=%v , errormsg=%v", resQyAccessToken.ErrCode, resQyAccessToken.ErrMsg)
		return
	}

	qyAccessTokenCacheKey := fmt.Sprintf("qy_access_token_%s", ctx.AppID)
	expires := resQyAccessToken.ExpiresIn - 1500
	err = ctx.Cache.Set(qyAccessTokenCacheKey, resQyAccessToken.AccessToken, time.Duration(expires)*time.Second)
	return
}
