package user

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/util"
)

const (
	userInfoURL = "https://api.weixin.qq.com/cgi-bin/user/info"
)

//User 用户管理
type User struct {
	*context.Context
}

//NewUser 实例化
func NewUser(context *context.Context) *User {
	user := new(User)
	user.Context = context
	return user
}

//Info 用户基本信息
type Info struct {
	util.CommonError

	Subscribe     int32    `json:"subscribe"`
	OpenID        string   `json:"openid"`
	Nickname      string   `json:"nickname"`
	Sex           int32    `json:"sex"`
	City          string   `json:"city"`
	Country       string   `json:"country"`
	Province      string   `json:"province"`
	Language      string   `json:"language"`
	Headimgurl    string   `json:"headimgurl"`
	SubscribeTime int32    `json:"subscribe_time"`
	UnionID       string   `json:"unionid"`
	Remark        string   `json:"remark"`
	GroupID       int32    `json:"groupid"`
	TagidList     []string `json:"tagid_list"`
}

//GetUserInfo 获取用户基本信息
func (user *User) GetUserInfo(openID string) (userInfo *Info, err error) {
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s&openid=%s&lang=zh_CN", userInfoURL, accessToken, openID)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	userInfo = new(Info)
	err = json.Unmarshal(response, userInfo)
	if err != nil {
		return
	}
	if userInfo.ErrCode != 0 {
		err = fmt.Errorf("GetUserInfo Error , errcode=%d , errmsg=%s", userInfo.ErrCode, userInfo.ErrMsg)
		return
	}
	return
}
