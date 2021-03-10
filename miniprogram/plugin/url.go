package plugin

import (
	"fmt"
	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	pluginUrl = "https://api.weixin.qq.com/wxa/plugin?access_token=%s"
)

//plugin struct
type Plugin struct {
	*context.Context
}

//PostBody 请求数据
type PostBody struct {
	Action      string `json:"action"`
	PluginAppid string `json:"plugin_appid"`
	UserVersion string `json:"user_version"`
}

//PluginRes 返回格式
type PluginRes struct {
	util.CommonError
	PluginList []*PluginInfo `json:"plugin_list"`
}

//PluginInfo 插件列表
type PluginInfo struct {
	Appid      string       `json:"appid"`
	Status     PluginStatus `json:"status"`
	Nickname   string       `json:"nickname"`
	Headimgurl string       `json:"headimgurl"`
}

//PluginStatus 插件状态
type PluginStatus int64

const PluginStatusApplying PluginStatus = 1  //插件申请状态:申请中
const PluginStatusApplyPass PluginStatus = 2 //插件申请状态:申请通过
const PluginApplyRefuse PluginStatus = 3     //插件申请状态:申请拒绝
const PluginApplyTimeout PluginStatus = 4    //插件申请状态:申请超时

//NewPlugin 实例
func NewPlugin(context *context.Context) *Plugin {
	plugin := new(Plugin)
	plugin.Context = context
	return plugin
}

//Apply 插件申请
func (plugin *Plugin) Apply(pluginAppid string) (res PluginRes, err error) {
	var (
		accessToken string
		urlStr      string
		body        = &PostBody{}
		response    []byte
	)

	accessToken, err = plugin.GetAccessToken()
	if err != nil {
		return
	}
	urlStr = fmt.Sprintf(pluginUrl, accessToken)
	body.Action = "apply"
	body.PluginAppid = pluginAppid
	response, err = util.PostJSON(urlStr, body)
	err = util.DecodeWithError(response, &res, "PluginApply")
	return
}

//Apply 插件列表

func (plugin *Plugin) List() (res PluginRes, err error) {
	var (
		accessToken string
		urlStr      string
		body        = &PostBody{}
		response    []byte
	)

	accessToken, err = plugin.GetAccessToken()
	if err != nil {
		return
	}
	urlStr = fmt.Sprintf(pluginUrl, accessToken)
	body.Action = "list"
	response, err = util.PostJSON(urlStr, body)
	err = util.DecodeWithError(response, &res, "PluginList")
	return
}

//Apply 插件解除绑定
func (plugin *Plugin) Unbind(pluginAppid string) (res PluginRes, err error) {
	var (
		accessToken string
		urlStr      string
		body        = &PostBody{}
		response    []byte
	)

	accessToken, err = plugin.GetAccessToken()
	if err != nil {
		return
	}
	urlStr = fmt.Sprintf(pluginUrl, accessToken)
	body.Action = "unbind"
	body.PluginAppid = pluginAppid

	response, err = util.PostJSON(urlStr, body)
	err = util.DecodeWithError(response, &res, "PluginUnbind")
	return
}

//Apply 插件更新
func (plugin *Plugin) Update(pluginAppid string, userVersion string) (res PluginRes, err error) {
	var (
		accessToken string
		urlStr      string
		body        = &PostBody{}
		response    []byte
	)

	accessToken, err = plugin.GetAccessToken()
	if err != nil {
		return
	}
	urlStr = fmt.Sprintf(pluginUrl, accessToken)
	body.Action = "update"

	body.PluginAppid = pluginAppid
	body.UserVersion = userVersion

	response, err = util.PostJSON(urlStr, body)
	err = util.DecodeWithError(response, &res, "PluginUpdate")
	return
}
