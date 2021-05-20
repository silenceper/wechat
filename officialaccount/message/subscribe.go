package message

import (
	"fmt"
	"strings"

	"github.com/silenceper/wechat/v2/officialaccount/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	//发送订阅消息
	//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
	subscribeSendURL = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send"
	//组合模板并添加到个人模板库
	subscribeAddURL = "https://api.weixin.qq.com/wxaapi/newtmpl/addtemplate"
	//删除帐号下的某个模板
	subscribeDelURL = "https://api.weixin.qq.com/wxaapi/newtmpl/deltemplate"
	//获取当前帐号所设置的类目信息
	subscribeCategoryURL = "https://api.weixin.qq.com/wxaapi/newtmpl/getcategory"
	//获取模板标题下的关键词库
	subscribeKeywordURL = "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatekeywords"
	//获取模板标题列表
	subscribePubListURL = "https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatetitles"
	// 获取当前帐号下的个人模板列表
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getTemplateList.html
	getTemplateURL = "https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate"
)

// Subscribe 订阅消息
type Subscribe struct {
	*context.Context
}

// NewSubscribe 实例化
func NewSubscribe(ctx *context.Context) *Subscribe {
	return &Subscribe{Context: ctx}
}

// Message 订阅消息请求参数
type Message struct {
	ToUser           string               `json:"touser"`            //必选，接收者（用户）的 openid
	TemplateID       string               `json:"template_id"`       //必选，所需下发的订阅模板id
	Page             string               `json:"page"`              //可选，点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	Data             map[string]*DataItem `json:"data"`              //必选, 模板内容
	MiniprogramState string               `json:"miniprogram_state"` //可选，跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang             string               `json:"lang"`              //入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}

//DataItem 模版内某个 .DATA 的值
type DataItem struct {
	Value string `json:"value"`
}

//TemplateItem template item
type SubscribeTemplateItem struct {
	PriTmplID string `json:"priTmplId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Example   string `json:"example"`
	Type      int64  `json:"type"`
}

//TemplateList template list
type TemplateList struct {
	util.CommonError
	Data []SubscribeTemplateItem `json:"data"`
}

type SubscribePubListParams struct {
	Ids   []string `json:"ids"`
	Start int64    `json:"start"`
	Limit int64    `json:"limit"`
}
type subscribePubListRes struct {
	util.CommonError
	Count int64                   `json:"count"`
	Data  []*subscribePubListData `json:"data"`
}

type subscribePubListData struct {
	Tid        int64  `json:"tid"`
	Title      string `json:"title"`
	Type       int64  `json:"type"`
	CategoryId string `json:"categoryId"`
}

type SubscribeCategoryRes struct {
	util.CommonError
	Data []*SubscribeCategoryData `json:"data"`
}
type SubscribeCategoryData struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// Send 发送订阅消息
func (s *Subscribe) Send(msg *Message) (err error) {
	var accessToken string
	accessToken, err = s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", subscribeSendURL, accessToken)
	response, err := util.PostJSON(uri, msg)
	if err != nil {
		return
	}
	return util.DecodeWithCommonError(response, "Send")
}

//ListTemplates 获取当前帐号下的个人模板列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getTemplateList.html
func (s *Subscribe) ListTemplates() (*TemplateList, error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", getTemplateURL, accessToken)
	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}
	templateList := TemplateList{}
	err = util.DecodeWithError(response, &templateList, "ListTemplates")
	if err != nil {
		return nil, err
	}
	return &templateList, nil
}

//GetTemplatePubList 获取公共模板列表
func (s *Subscribe) GetSubscribePubList(params *SubscribePubListParams) (templateList []*subscribePubListData, err error) {
	var accessToken string
	accessToken, err = s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s&ids=%s&start=%d&limit=%d", subscribePubListURL, accessToken, strings.Join(params.Ids, ","), params.Start, params.Limit)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var res subscribePubListRes
	err = util.DecodeWithError(response, &res, "ListPubTemplate")
	if err != nil {
		return
	}
	templateList = res.Data
	return
}

//GetCategory 获取公众号类目
func (s *Subscribe) GetCategory() (templateCategoryList []*SubscribeCategoryData, err error) {
	var accessToken string
	accessToken, err = s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", subscribeCategoryURL, accessToken)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var res SubscribeCategoryRes
	err = util.DecodeWithError(response, &res, "Category")
	if err != nil {
		return
	}
	templateCategoryList = res.Data
	return
}

type SubscribeKeywordRes struct {
	util.CommonError
	Data []*SubscribeKeywordData `json:"data"`
}
type SubscribeKeywordData struct {
	Kid     int64  `json:"kid"`
	Name    string `json:"name"`
	Example string `json:"example"`
	Rule    string `json:"rule"`
}

//GetTemplateKeyword 获取模板中的关键词
func (s *Subscribe) GetTemplateKeyword(tid string) (templateKeyword []*SubscribeKeywordData, err error) {
	var accessToken string
	accessToken, err = s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s&tid=%s", subscribeKeywordURL, accessToken, tid)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var res SubscribeKeywordRes
	err = util.DecodeWithError(response, &res, "keyword")
	if err != nil {
		return
	}
	templateKeyword = res.Data
	return
}

type SubscribeAddParams struct {
	Tid       string  `json:"tid"`
	KidList   []int64 `json:"kidList"`
	SceneDesc string  `json:"sceneDesc"`
}
type SubscribeAddRes struct {
	util.CommonError
	PriTmplId string `json:"priTmplId"`
}

//Add 获得模板ID
func (s *Subscribe) Add(params *SubscribeAddParams) (templateId string, err error) {
	var accessToken string
	var templateAdd SubscribeAddRes
	accessToken, err = s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", subscribeAddURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, params)

	if err != nil {
		return
	}
	err = util.DecodeWithError(response, &templateAdd, "ListTemplate")
	if err != nil {
		return
	}
	return templateAdd.PriTmplId, nil
}

type SubscribeDelRes struct {
	util.CommonError
}

//Del 删除模版ID
func (s *Subscribe) Del(priTmplId string) (err error) {
	var accessToken string
	var templateDel SubscribeDelRes
	accessToken, err = s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", subscribeDelURL, accessToken)
	var response []byte
	response, err = util.PostJSON(uri, struct {
		PriTmplId string `json:"priTmplId"`
	}{
		PriTmplId: priTmplId,
	})

	if err != nil {
		return
	}
	err = util.DecodeWithError(response, &templateDel, "Del")
	if err != nil {
		return
	}
	return
}
