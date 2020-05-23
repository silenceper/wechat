package basic

import (
	"fmt"

	"github.com/silenceper/wechat/v2/officialaccount/context"
	"github.com/silenceper/wechat/v2/util"
)

var (
	//获取微信服务器IP地址
	//文档：https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html
	getCallbackIPURL  = "https://api.weixin.qq.com/cgi-bin/getcallbackip"
	getAPIDomainIPURL = "https://api.weixin.qq.com/cgi-bin/get_api_domain_ip"
)

//Basic struct
type Basic struct {
	*context.Context
}

//NewBasic 实例
func NewBasic(context *context.Context) *Basic {
	basic := new(Basic)
	basic.Context = context
	return basic
}

//IPListRes 获取微信服务器IP地址 返回结果
type IPListRes struct {
	util.CommonError
	IPList []string `json:"ip_list"`
}

//GetCallbackIP 获取微信callback IP地址
func (basic *Basic) GetCallbackIP() ([]string, error) {
	ak, err := basic.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s?access_token=%s", getCallbackIPURL, ak)
	data, err := util.HTTPGet(url)
	if err != nil {
		return nil, err
	}
	ipListRes := &IPListRes{}
	err = util.DecodeWithError(data, ipListRes, "GetCallbackIP")
	return ipListRes.IPList, err
}

//GetAPIDomainIP 获取微信API接口 IP地址
func (basic *Basic) GetAPIDomainIP() ([]string, error) {
	ak, err := basic.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s?access_token=%s", getAPIDomainIPURL, ak)
	data, err := util.HTTPGet(url)
	if err != nil {
		return nil, err
	}
	ipListRes := &IPListRes{}
	err = util.DecodeWithError(data, ipListRes, "GetAPIDomainIP")
	return ipListRes.IPList, err
}
